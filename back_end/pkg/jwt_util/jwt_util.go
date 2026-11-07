package jwtutil

import (
	"errors"
	"github.com/gofiber/contrib/websocket"
	"shoollove/config"
	"shoollove/model"
	"shoollove/pkg/db"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwt(key any, method jwt.SigningMethod, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(method, claims)
	return token.SignedString(key)
}

func ParseJwt(key any, jwtStr string, options ...jwt.ParserOption) (jwt.Claims, error) {
	token, err := jwt.Parse(jwtStr, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	}, options...)
	if err != nil {
		return nil, err
	}
	// 校验 Claims 对象是否有效，基于 exp（过期时间），nbf（不早于），iat（签发时间）等进行判断（如果有这些声明的话）。
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token.Claims, nil
}

func ParseJwtWithClaims(key any, jwtStr string, options ...jwt.ParserOption) (jwt.MapClaims, error) {
	mc := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(jwtStr, mc, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	}, options...)
	if err != nil {
		return nil, err
	}
	// 校验 Claims 对象是否有效，基于 exp（过期时间），nbf（不早于），iat（签发时间）等进行判断（如果有这些声明的话）。
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token.Claims.(jwt.MapClaims), nil
}

func UserInfo(c *fiber.Ctx) (model.Account, error) {
	token := GetToken(c)
	claims, err := ParseJwtWithClaims([]byte(config.Config.Jwt.Key), token)
	if err != nil {
		return model.Account{}, err
	}
	// token中只有email
	email := claims["phone"].(string)
	var account model.Account
	err = db.Gorm.Table("account").Where("phone = ?", email).First(&account).Error
	if err != nil {
		return model.Account{}, err
	}
	return account, nil
}

func GetToken(c *fiber.Ctx) string {
	token := c.Get("Authorization")
	if !strings.Contains(token, "Bearer ") {
		return ""
	}
	return splitBearerToken(token)

}

func splitBearerToken(atuhStr string) string {
	atuhStr = strings.Replace(atuhStr, "Bearer", "", 1)
	return strings.TrimSpace(atuhStr)
}

func GetWSToken(c *websocket.Conn) string {
	token := c.Headers("Authorization")
	if !strings.Contains(token, "Bearer ") {
		return ""
	}
	return splitBearerToken(token)
}

func WSUserInfo(c *websocket.Conn) (model.Account, error) {
	token := GetWSToken(c)
	claims, err := ParseJwtWithClaims([]byte(config.Config.Jwt.Key), token)
	if err != nil {
		return model.Account{}, err
	}
	// token中只有email
	email := claims["phone"].(string)
	var account model.Account
	err = db.Gorm.Table("account").Where("phone = ?", email).First(&account).Error
	if err != nil {
		return model.Account{}, err
	}
	return account, nil
}
