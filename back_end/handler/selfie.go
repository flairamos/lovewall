package handler

import (
	"shoollove/model"
	jwtutil "shoollove/pkg/jwt_util"
	"shoollove/pkg/respond"
	"shoollove/service"

	"github.com/gofiber/fiber/v2"
)

// SelfieAdd 添加
func SelfieAdd(c *fiber.Ctx) error {

	var req model.SelfieAdd

	if err := c.BodyParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}

	userInfo, err := jwtutil.UserInfo(c)
	if err != nil {
		return respond.Fail(c, "请先登录", err)
	}
	resp, err := service.SelfieAdd(req, userInfo)
	if err != nil {
		return respond.Fail(c, "添加失败", err)
	}
	return respond.Ok(c, resp)
}

// SelfieEdit 修改
func SelfieEdit(c *fiber.Ctx) error {
	var req model.Selfie
	if err := c.BodyParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}

	resp, err := service.SelfieEdit(req)
	if err != nil {
		return respond.Fail(c, err.Error(), err)
	}
	return respond.Ok(c, resp)
}

// SelfieDelete 删除
func SelfieDelete(c *fiber.Ctx) error {

	var req model.Selfie

	if err := c.BodyParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}

	resp, err := service.SelfieDelete(req)
	if err != nil {
		return respond.Fail(c, err.Error(), err)
	}
	return respond.Ok(c, resp)
}

// SelfieGet 查询
func SelfieGet(c *fiber.Ctx) error {

	var req model.SelfieGetByType

	if err := c.QueryParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}
	user, err := jwtutil.UserInfo(c)
	if err != nil {
	}
	resp, err := service.SelfieGet(req, user)
	if err != nil {
		return respond.Fail(c, err.Error(), err)
	}
	return respond.Ok(c, resp)
}

// 帖子点赞
func SelfiePraise(c *fiber.Ctx) error {
	id := c.QueryInt("id", 0)
	if id == 0 {
		return respond.Fail(c, "点赞帖子不存在", nil)
	}
	_, err := service.SelfiePraise(model.SelfiePraise{Id: id})
	if err != nil {
		return respond.Fail(c, "点赞失败", err)
	}
	return respond.Ok(c, "")
}

// SelfieForm 精准查询
func SelfieForm(c *fiber.Ctx) error {

	var req model.SelfieGetByType

	if err := c.QueryParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}

	resp, err := service.SelfieForm(req)
	if err != nil {
		return respond.Fail(c, err.Error(), err)
	}
	return respond.Ok(c, resp)
}
