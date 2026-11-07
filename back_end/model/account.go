package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameAccount = "account"

// Account mapped from table <account>
type Account struct {
	ID                 int32          `gorm:"column:id;autoIncrement:true;comment:主键" json:"id"` // 主键
	Phone              string         `gorm:"column:phone;not null;comment:账户" json:"phone"`     // 账户
	Name               string         `gorm:"column:name;not null;comment:昵称" json:"name"`       // 昵称
	Password           string         `gorm:"column:password;not null;comment:密码" json:"-"`
	Sex                bool           `gorm:"column:sex;not null;comment:性别" json:"sex"`                       // 性别
	Work               int32          `gorm:"column:work;not null;comment:运动步数" json:"work"`                   // 运动步数
	Openid             string         `gorm:"column:openid;not null;comment:wechat_miniprogram" json:"openid"` // wechat_miniprogram
	SessionKey         string         `gorm:"column:session_key" json:"session_key"`
	Authentication     bool           `gorm:"column:authentication;not null;comment:认证状态" json:"authentication"`           // 认证状态
	AuthenticationName int32          `gorm:"column:authentication_name;not null;comment:认证信息" json:"authentication_name"` // 认证信息
	AuthenticationEtc  string         `gorm:"column:authentication_etc;not null;comment:认证材料" json:"authentication_etc"`   // 认证材料
	BackgroundImg      string         `gorm:"column:background_img;not null;comment:背景图片" json:"background_img"`           // 背景图片
	Following          string         `gorm:"column:following;not null;comment:关注" json:"following"`                       // 关注
	Follower           string         `gorm:"column:follower;not null;comment:粉丝" json:"follower"`                         // 粉丝
	Avatar             string         `gorm:"column:avatar;not null;avatar:头像" json:"avatar"`
	CreatedAt          time.Time      `gorm:"column:created_at;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt          time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updated_at"` // 更新时间
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"` // 删除时间
}

// TableName Account's table name
func (*Account) TableName() string {
	return TableNameAccount
}

type AccountAvatar struct {
	ID     int32  `gorm:"column:id;" json:"id"`                           // 主键
	Phone  string `gorm:"column:phone;not null;comment:手机号" json:"phone"` // 邮箱
	Name   string `gorm:"column:name;not null;comment:昵称" json:"name"`    // 昵称
	Avatar string `gorm:"column:avatar;not null;avatar:头像" json:"avatar"` // 头像
}

type RegisterCode struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	JsCode   string `json:"js_code"` // 小程序登录码
	Code     string `json:"code"`
}
