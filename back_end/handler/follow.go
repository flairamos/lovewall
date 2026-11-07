package handler

import (
	"gorm.io/gorm"
	"shoollove/model"
	"shoollove/pkg/db"
	jwtutil "shoollove/pkg/jwt_util"
	"shoollove/pkg/respond"
	"shoollove/service"

	"github.com/gofiber/fiber/v2"
)

func FollowAdd(c *fiber.Ctx) error {

	var req model.FollowAdd

	if err := c.BodyParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}
	userInfo, err := jwtutil.UserInfo(c)
	if err != nil {
		return respond.Fail(c, "请先登录", err)
	}
	resp, err := service.FollowAdd(req, userInfo)
	if err != nil {
		return respond.Fail(c, resp.(string), err)
	}
	return respond.Ok(c, resp, "关注成功")
}

func FollowDelete(c *fiber.Ctx) error {

	var req model.FollowAdd

	if err := c.BodyParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}
	userInfo, err := jwtutil.UserInfo(c)
	if err != nil {
		return respond.Fail(c, "请先登录", err)
	}
	resp, err := service.FollowDelete(req, userInfo)
	if err != nil {
		return respond.Fail(c, resp.(string), err)
	}
	return respond.Ok(c, resp, "取关成功")
}

// 判断是否关注
func FollowGet(c *fiber.Ctx) error {
	var req model.FollowAdd
	if err := c.QueryParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}
	userInfo, err := jwtutil.UserInfo(c)
	if err != nil {
		return respond.Fail(c, "请先登录", err)
	}
	resp, err := service.FollowGet(req, userInfo)
	if err != nil {
		return respond.Fail(c, "查询失败", err)
	}
	return respond.Ok(c, resp.(bool))
}

// 获取用户的关注信息
func FollowInfo(c *fiber.Ctx) error {
	var req model.FollowInfo
	if err := c.QueryParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}
	resp, err := service.FollowInfo(req)
	if err != nil {
		return respond.Fail(c, "查询失败", err)
	}
	return respond.Ok(c, resp)
}

// 查询关注列表
func FindFollowList(c *fiber.Ctx) error {
	userInfo, err := jwtutil.UserInfo(c)
	if err != nil {
		return respond.Fail(c, "请先登录", err)
	}
	var result = make([]model.AccountAvatar, 0)
	errTx := db.Gorm.Transaction(func(tx *gorm.DB) error {
		var follow []model.Follow
		err = tx.Table("follow").Where("user_id = ?", userInfo.ID).Find(&follow).Error
		if err != nil {
			return err
		}
		for _, fo := range follow {
			acc := model.AccountAvatar{}
			err = tx.Table("account").
				Select("id, name, phone, avatar").
				Where("phone = ?", fo.CollectPhone).First(&acc).Error
			if err != nil {
				return err
			}
			result = append(result, acc)
		}
		return nil
	})
	if errTx != nil {
		return respond.Fail(c, "查询关注列表是失败", errTx)
	}
	return respond.Ok(c, result)
}
