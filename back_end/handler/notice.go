package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"shoollove/model"
	jwtutil "shoollove/pkg/jwt_util"
	"shoollove/pkg/respond"
	"shoollove/service"
)

// TopicAdd 添加
func TopicAdd(c *fiber.Ctx) error {

	var req model.NoticeAddReq

	if err := c.BodyParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}

	_, err := service.TopicAdd(req)
	if err != nil {
		return respond.Fail(c, "添加失败！", err)
	}
	return respond.Ok(c, nil, "添加成功")
}

// TopicEdit 修改
func TopicEdit(c *fiber.Ctx) error {

	var req model.Notice

	if err := c.BodyParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}

	resp, err := service.TopicEdit(req)
	if err != nil {
		return respond.Fail(c, "", err)
	}
	return respond.Ok(c, resp)
}

// TopicDelete 删除
func TopicDelete(c *fiber.Ctx) error {

	var req model.Notice

	if err := c.BodyParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}

	resp, err := service.TopicDelete(req)
	if err != nil {
		return respond.Fail(c, "", err)
	}
	return respond.Ok(c, resp)
}

// TopicGet 查询
func TopicGet(c *fiber.Ctx) error {

	var req model.PageReq

	if err := c.QueryParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}
	userInfo, err := jwtutil.UserInfo(c)
	if err != nil {

	}
	resp, err := service.TopicGet(req, userInfo)
	if err != nil {
		return respond.Fail(c, "查询失败", err)
	}
	return respond.Ok(c, resp)
}

// TopicForm 精准查询
func TopicForm(c *fiber.Ctx) error {

	var req = c.QueryInt("id", 0)

	if req == 0 {
		return respond.Fail(c, "参数错误", errors.New("empty primary key id"))
	}

	resp, err := service.TopicForm(model.Notice{ID: int32(req)})
	if err != nil {
		return respond.Fail(c, "", err)
	}
	return respond.Ok(c, resp)
}

func TopicPraise(c *fiber.Ctx) error {
	var id = c.QueryInt("id", 0)
	if id == 0 {
		return respond.Fail(c, "缺少公告主键", errors.New("empty primary key id"))
	}
	info, err := jwtutil.UserInfo(c)
	if err != nil {
		return respond.Fail(c, "请先登录", err)
	}
	resp, err := service.TopicPraise(id, info)
	if err != nil {
		return respond.Fail(c, "点赞失败", err)
	}
	return respond.Ok(c, resp)
}
