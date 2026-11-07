package handler

import (
	"github.com/gofiber/fiber/v2"
	"shoollove/model"
	jwtutil "shoollove/pkg/jwt_util"
	"shoollove/pkg/respond"
	"shoollove/service"
)

// GoodsAdd 添加
func GoodsAdd(c *fiber.Ctx) error {
	var req model.GoodsAdd
	if err := c.BodyParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}
	userInfo, err := jwtutil.UserInfo(c)
	if err != nil {
		return respond.Fail(c, "请先登录", err)
	}
	resp, err := service.GoodsAdd(req, userInfo)
	if err != nil {
		return respond.Fail(c, "添加失败", err)
	}
	return respond.Ok(c, resp)
}

// GoodsEdit 修改
func GoodsEdit(c *fiber.Ctx) error {

	var req model.Goods

	if err := c.BodyParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}

	resp, err := service.GoodsEdit(req)
	if err != nil {
		return respond.Fail(c, "编辑失败", err)
	}
	return respond.Ok(c, resp)
}

// GoodsDelete 删除
func GoodsDelete(c *fiber.Ctx) error {

	var req model.Goods

	if err := c.BodyParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}

	resp, err := service.GoodsDelete(req)
	if err != nil {
		return respond.Fail(c, "删除失败", err)
	}
	return respond.Ok(c, resp)
}

// GoodsGet 查询
func GoodsGet(c *fiber.Ctx) error {

	var req model.GoodsGet

	if err := c.QueryParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}

	resp, err := service.GoodsGet(req)
	if err != nil {
		return respond.Fail(c, "查询失败", err)
	}
	return respond.Ok(c, resp)
}

// GoodsForm 精准查询
func GoodsForm(c *fiber.Ctx) error {

	var req model.Goods

	if err := c.QueryParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}

	resp, err := service.GoodsForm(req)
	if err != nil {
		return respond.Fail(c, "查询失败", err)
	}
	return respond.Ok(c, resp)
}
