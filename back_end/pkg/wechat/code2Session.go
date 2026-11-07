package wechat

import (
	"fmt"
	"shoollove/config"
)

const code2Session = "https://api.weixin.qq.com/sns/jscode2session"

func Code2SessionUrl(code string) string {
	return code2Session + fmt.Sprintf("?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", config.Config.Miniprogram.AppID, config.Config.Miniprogram.AppSecret, code)
}

type Code2SessionResp struct {
	SessionKey string `json:"session_key"`
	Openid     string `json:"openid"`
	Unionid    string `json:"unionid"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}
