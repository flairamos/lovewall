package handler

import (
	"errors"
	"shoollove/config"
	"shoollove/model"
	"shoollove/pkg/db"
	jwtutil "shoollove/pkg/jwt_util"
	"shoollove/pkg/respond"
	"shoollove/pkg/tool"
	"shoollove/service"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// AccountAdd 添加
func AccountAdd(c *fiber.Ctx) error {

	var req model.RegisterCode

	if err := c.BodyParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}

	resp, err := service.AccountAdd(req)
	if err != nil {
		return respond.Fail(c, resp.(string), err)
	}
	return respond.Ok(c, resp, "注册成功")
}

// AccountEdit 修改
func AccountEdit(c *fiber.Ctx) error {
	var req model.Account

	if err := c.BodyParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}

	resp, err := service.AccountEdit(req)
	if err != nil {
		return respond.Fail(c, "编辑失败！", err)
	}
	return respond.Ok(c, resp)
}

// AccountDelete 删除
func AccountDelete(c *fiber.Ctx) error {

	var req model.Account

	if err := c.BodyParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}

	resp, err := service.AccountDelete(req)
	if err != nil {
		return respond.Fail(c, "删除失败", err)
	}
	return respond.Ok(c, resp)
}

// AccountGet 查询
func AccountGet(c *fiber.Ctx) error {

	var req model.Account

	if err := c.QueryParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}

	resp, err := service.AccountGet(req)
	if err != nil {
		return respond.Fail(c, "查询失败", err)
	}
	return respond.Ok(c, resp)
}

// AccountForm 精准查询
func AccountForm(c *fiber.Ctx) error {

	var req model.AccountAvatar

	if err := c.QueryParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}

	resp, err := service.AccountForm(req)
	if err != nil {
		return respond.Fail(c, "查询失败", err)
	}
	return respond.Ok(c, resp)
}

func AccountById(c *fiber.Ctx) error {
	var req model.AccountAvatar
	if err := c.QueryParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}
	resp, err := service.AccountById(req)
	if err != nil {
		return respond.Fail(c, "查询失败", err)
	}
	return respond.Ok(c, resp)
}

// 发验证码
func SendRegisterCode(c *fiber.Ctx) error {
	var req model.RegisterCode
	if err := c.BodyParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}
	resp, err := service.SendRegisterCode(req)
	if err != nil {
		return respond.Fail(c, resp.(string), err)
	}
	return respond.Ok(c, resp, "发送成功")
}

func CheckRegisterCode(c *fiber.Ctx) error {
	var req model.RegisterCode
	if err := c.BodyParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}
	_, err := service.CheckRegisterCode(req)
	if err != nil {
		return respond.Fail(c, "验证码验证失败", err)
	}
	return respond.Ok(c, nil)
}

func Login(c *fiber.Ctx) error {
	var req model.RegisterCode
	if err := c.BodyParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}
	if req.Name == "" || req.Password == "" {
		return respond.Fail(c, "账号或者密码不能为空", errors.New("name or pass can not be empty"))
	}
	if !tool.ValidateEmail(req.Name) {
		return respond.Fail(c, "邮箱格式不对", errors.New("invalid email"))
	}
	var account []model.Account
	err := db.Gorm.Table("account").
		Where("phone = ?", req.Name).
		Find(&account).Error
	if err != nil {
		return respond.Fail(c, "系统异常请稍后再试！", err)
	}
	if len(account) < 1 {
		return respond.Fail(c, "不存在此用户,请先注册", errors.New("no such user"))
	}
	if account[0].Password != req.Password {
		return respond.Fail(c, "账号或密码错误", errors.New("invalid password"))
	}
	message := make(jwt.MapClaims)
	message["phone"] = req.Name
	jwtStr, err := jwtutil.GenerateJwt([]byte(config.Config.Jwt.Key), jwt.SigningMethodHS256, message)
	if err != nil {
		return respond.Fail(c, "jwt生成失败", err)
	}
	var result = make(map[string]any)
	result["token"] = jwtStr
	result["user_info"] = account[0]
	return respond.Ok(c, result, "登录成功")
}
func CheckLogin(c *fiber.Ctx) error {
	token := jwtutil.GetToken(c)
	if token == "" {
		return respond.Fail(c, "没有登陆权限", errors.New("invalid token"))
	}
	_, err := jwtutil.ParseJwt([]byte(config.Config.Jwt.Key), token)
	if err != nil {
		return respond.Fail(c, "权限错误", err)
	}
	return respond.Ok(c, "操作成功")
}
