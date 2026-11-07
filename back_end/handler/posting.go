package handler

import (
	"errors"
	"shoollove/enum"
	"shoollove/model"
	jwtutil "shoollove/pkg/jwt_util"
	"shoollove/pkg/respond"
	"shoollove/service"
	"slices"

	"github.com/gofiber/fiber/v2"
)

// PostingAdd 添加
func PostingAdd(c *fiber.Ctx) error {

	var req model.PostingAdd

	if err := c.BodyParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}

	resp, err := service.PostingAdd(req)
	if err != nil {
		return respond.Fail(c, "添加失败", err)
	}
	return respond.Ok(c, resp)
}

// PostingEdit 修改
func PostingEdit(c *fiber.Ctx) error {

	var req model.Posting

	if err := c.BodyParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}

	resp, err := service.PostingEdit(req)
	if err != nil {
		return respond.Fail(c, "", err)
	}
	return respond.Ok(c, resp)
}

// PostingDelete 删除
func PostingDelete(c *fiber.Ctx) error {

	var req model.PostingDelete

	if err := c.BodyParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}
	if req.Id == 0 {
		return respond.Fail(c, "删除帖子状态错误", errors.New("缺少主键id"))
	}
	resp, err := service.PostingDelete(req)
	if err != nil {
		return respond.Fail(c, "删除失败", err)
	}
	return respond.Ok(c, resp)
}

// PostingGet 查询
// 查询帖子列表和点赞最多的三条评论
func PostingGet(c *fiber.Ctx) error {
	var req model.PostGet
	if err := c.QueryParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}
	info, err := jwtutil.UserInfo(c)
	if err != nil {
	}
	if !slices.Contains(enum.POSTING_TYPE, req.Type) {
		return respond.Fail(c, "缺少必要查询参数！", nil)
	}
	resp, err := service.PostingGet(req, info)
	if err != nil {
		return respond.Fail(c, "查询失败！", err)
	}
	return respond.Ok(c, resp)
}

// PostingForm 精准查询
func PostingForm(c *fiber.Ctx) error {

	var req model.PostingForm

	if err := c.QueryParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}
	info, err := jwtutil.UserInfo(c)
	if err != nil {
	}
	if !slices.Contains(enum.POSTING_TYPE, req.Type) {
		return respond.Fail(c, "缺少必要查询参数！", nil)
	}
	resp, err := service.PostingForm(req, info)
	if err != nil {
		return respond.Fail(c, "", err)
	}
	return respond.Ok(c, resp)
}

// 查询评论区
func GetCommentsById(c *fiber.Ctx) error {
	var req model.PostingForm
	if err := c.QueryParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}
	info, err := jwtutil.UserInfo(c)
	if err != nil {
	}
	resp, err := service.GetCommentById(req, info)
	if err != nil {
		return respond.Fail(c, "评论查询失败", err)
	}
	return respond.Ok(c, resp)
}

// 点赞
func PostingPraise(c *fiber.Ctx) error {
	var req model.PostingLike
	if err := c.BodyParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}
	info, err := jwtutil.UserInfo(c)
	if err != nil {

	}
	resp, err := service.PostingPraise(req, info)
	if err != nil {
		return respond.Fail(c, "点赞失败", err)
	}
	return respond.Ok(c, resp)
}
