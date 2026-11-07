package handler

import (
	"shoollove/model"
	"shoollove/service"
)

// 根据用户id查询用户信息
func UserInfoById(id int) (model.Account, error) {
	return service.UserInfoById(id)
}
