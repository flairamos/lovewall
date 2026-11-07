package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSelfie = "selfie"

// Selfie mapped from table <selfie>
type Selfie struct {
	ID        int32          `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"` // 主键
	Phone     string         `gorm:"column:phone;not null;comment:手机号" json:"phone"`               // 手机号
	UserID    int32          `gorm:"column:user_id;not null;comment:用户id" json:"user_id"`          // 用户id
	LikeNum   int32          `gorm:"column:like_num;not null;comment:点赞数" json:"like_num"`         // 点赞数
	Photos    string         `gorm:"column:photos;not null;comment:图片" json:"photos"`              // 图片
	Title     string         `gorm:"column:title;not null;comment:标题" json:"title"`                // 标题
	CreatedAt time.Time      `gorm:"column:created_at;comment:创建时间" json:"created_at"`             // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`             // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`             // 删除时间
}

// TableName Selfie's table name
func (*Selfie) TableName() string {
	return TableNameSelfie
}

type SelfieGetByType struct {
	Type int    `json:"type"`
	Name string `json:"name"`
	PageReq
}

type SelfiePraise struct {
	Id int `query:"id" json:"id"`
}

type SelfieAdd struct {
	Photos []string `json:"photos"` // 图片
	Title  string   `json:"title"`  // 标题
}

type SelfieListItem struct {
	Id      int32         `json:"id"`
	Poster  AccountAvatar `json:"poster"`
	Title   string        `json:"title"`
	LikeNum int32         `json:"like_num"`
	Photos  []string      `json:"photos"`
}

type SelfieList struct {
	LeftList  []SelfieListItem `json:"leftList"`
	RightList []SelfieListItem `json:"rightList"`
}
