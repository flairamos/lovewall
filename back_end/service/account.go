package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"shoollove/model"
	"shoollove/pkg/db"
	redisutil "shoollove/pkg/redis_util"
	"shoollove/pkg/wechat"
	"strings"
)

// AccountAdd 添加
func AccountAdd(req model.RegisterCode) (resp any, err error) {
	if req.Name == "" || req.Password == "" {
		return "账户密码不能为空", errors.New("name or pass can not be not")
	}
	// 获取验证码
	req.Name = strings.TrimSpace(req.Name)
	req.Password = strings.TrimSpace(req.Password)
	codeTmp, err := redisutil.Client.Cmd("GET", req.Name)
	if err != nil {
		return "校验验证码失败", err
	}
	if codeTmp == nil {
		return "验证码失效或者验证码错误", errors.New("验证码失效或者验证码错误")
	}
	var by []byte = codeTmp.([]uint8)
	code := string(by)
	if code != req.Code {
		return "验证码错误", errors.New("验证码错误")
	}
	// 查询是否存在
	var emailNum int64
	err = db.Gorm.Table("account").Where("phone = ?", req.Name).Count(&emailNum).Error
	if err != nil {
		return "系统错误请稍候再试", err
	}
	if emailNum > 0 {
		return "账户已存在，请勿重复注册", errors.New("账户已存在，请勿重复注册")
	}
	// 获取用户openid
	if req.JsCode == "" {
		return "缺少小程序登录码", errors.New("缺少小程序登录码")
	}
	httpclient := &http.Client{}
	response, err := httpclient.Get(wechat.Code2SessionUrl(strings.TrimSpace(req.JsCode)))
	if err != nil {
		return "获取openid失败", errors.New("获取openid失败")
	}
	if response.StatusCode != http.StatusOK {
		return "系统繁忙，此时请开发者稍候再试", errors.New("系统繁忙，此时请开发者稍候再试")
	}
	all, err := io.ReadAll(response.Body)
	if err != nil {
		return "读取请求体失败", errors.New("读取请求体失败")
	}
	var code2session wechat.Code2SessionResp
	err = json.Unmarshal(all, &code2session)
	if err != nil {
		return "解析请求体失败", errors.New("解析请求体失败")
	}
	if code2session.Openid == "" {
		return "系统错误，请联系管理员", errors.New("系统错误，请联系管理员")
	}
	param := model.Account{
		Phone:      req.Name,
		Password:   req.Password,
		Openid:     code2session.Openid,
		SessionKey: code2session.SessionKey,
	}

	err = db.Gorm.Create(&param).Error
	if err != nil {
		return "创建用户失败，请稍后再试", err
	}
	return resp, nil
}

func SendRegisterCode(req model.RegisterCode) (resp any, err error) {
	code := rand.Intn(10000)
	// redis中间件存入验证码
	_, err = redisutil.Client.Cmd("set", req.Name, code)
	if err != nil {
		return "系统错误，请稍后再试", err
	}
	// 发送邮件
	err = SendRegisterEamil(fmt.Sprintf("%d", code), strings.TrimSpace(req.Name))
	if err != nil {
		return "发送邮件失败", err
	}
	return resp, nil
}

func CheckRegisterCode(req model.RegisterCode) (resp any, err error) {
	reply, err := redisutil.Client.Cmd("get", strings.TrimSpace(req.Name))
	if err != nil {
		return nil, errors.New("获取验证码失败")
	}
	if reply == nil {
		return nil, errors.New("验证码已失效，请重新发送")
	}

	if s, ok := reply.(string); ok {
		if s == strings.TrimSpace(req.Code) {
			return nil, nil
		}
	}
	return nil, errors.New("验证码错误")
}

// AccountEdit 编辑
func AccountEdit(req model.Account) (resp any, err error) {
	session := db.Gorm

	param := model.Account{}

	// var account account
	// 请求参数处理后添加到数据库
	// 赋值操作
	err = session.Updates(&param).Error
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// AccountDelete 删除
func AccountDelete(req model.Account) (resp any, err error) {
	session := db.Gorm

	param := model.Account{}

	err = session.Delete(&param).Error
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// AccountGet 查询
func AccountGet(req model.Account) (resp any, err error) {
	// session := db.Gorm

	// param := model.Account{}

	// var account account
	// 请求参数处理后添加到数据库
	// 赋值操作
	// result, err := tool.IPage(req.CurrentPage, req.PageSize, session, param)
	// if err != nil {
	// 	return resp, err
	// }
	return resp, nil
}

// AccountForm 查询
func AccountForm(req model.AccountAvatar) (resp any, err error) {
	if req.Phone == "" {
		return resp, errors.New("账户不能为空")
	}
	param := model.Account{}

	err = db.Gorm.Table("account").Where("phone = ?", req.Phone).First(&param).Error
	if err != nil {
		return resp, err
	}
	return param, nil
}

func AccountById(req model.AccountAvatar) (resp any, err error) {
	if req.ID == 0 {
		return resp, errors.New("账户ID不能为空")
	}
	param := model.AccountAvatar{}

	err = db.Gorm.Table("account").
		Select("id", "phone", "name", "avatar").
		Where("id = ?", req.ID).First(&param).Error
	if err != nil {
		return resp, err
	}
	return param, nil
}
