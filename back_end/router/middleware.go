package router

import (
	"errors"
	"shoollove/config"
	"shoollove/pkg/db"
	jwtutil "shoollove/pkg/jwt_util"
	"shoollove/pkg/respond"

	"github.com/gofiber/fiber/v2"
)

func JwtAuth(c *fiber.Ctx) error {
	token := jwtutil.GetToken(c)
	claim, err := jwtutil.ParseJwtWithClaims([]byte(config.Config.Jwt.Key), token)
	if err != nil {
		return respond.Fail(c, "没有登陆权限", errors.New("invalid token"))
	}
	// 查询用户是否存在
	phone := claim["phone"]
	if phone == "" {
		return respond.Fail(c, "没有登陆权限", errors.New("invalid token"))
	}
	var count int64
	err = db.Gorm.Table("account").Where("phone = ?", phone).Count(&count).Error
	if err != nil || count < 1 {
		return respond.Fail(c, "没有登陆权限", errors.New("invalid token"))
	}
	return c.Next()
}


func SocketAuth(c *fiber.Ctx) error {
	return JwtAuth(c)
}
