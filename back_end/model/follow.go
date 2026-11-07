package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameFollow = "follow"

// Follow mapped from table <Follow>
type Follow struct {
	ID           int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CollectPhone string         `gorm:"column:collect_phone;not null;comment:手机号" json:"collect_phone"` // 手机号
	CollectId    int32          `gorm:"column:collect_id;not null;comment:" json:"collect_id"`          // 收集id
	UserId       int32          `gorm:"column:user_id;not null;comment:内容" json:"user_id"`              // 内容
	Phone        string         `gorm:"column:phone;not null;comment:手机号" json:"phone"`                 // 手机号
	CreatedAt    time.Time      `gorm:"column:created_at;comment:创建时间" json:"created_at"`               // 创建时间
	UpdatedAt    time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`               // 更新时间
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`               // 删除时间
}

// TableName Notice's table name
func (*Follow) TableName() string {
	return TableNameFollow
}

type FollowAdd struct {
	Id           int    `query:"id" json:"id"`
	UserId       int32  `query:"user_id" json:"user_id"`
	CollectId    int32  `query:"collect_id" json:"collect_id"`
	CollectPhone string `query:"collect_phone" json:"collect_phone"`
	Phone        string `query:"phone" json:"phone"`
}

type FollowInfo struct {
	Phone string `query:"phone" json:"phone"`
}
