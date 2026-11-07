package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameMessageGroup = "message_group"

// MessageGroup mapped from table <message_group>
type MessageGroup struct {
	ID          int32          `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`               // 主键
	CreateUser  int32          `gorm:"column:create_user;not null;comment:创建人主键" json:"create_user"`               // 创建人主键
	CreatePhone string         `gorm:"column:create_phone;not null;comment:创建人账户" json:"create_phone"`             // 创建人账户
	GroupName   string         `gorm:"column:group_name;not null;comment:群组名称" json:"group_name"`                  // 群组名称
	CreatedAt   time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt   time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`                           // 更新时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                           // 删除时间
}

// TableName MessageGroup's table name
func (*MessageGroup) TableName() string {
	return TableNameMessageGroup
}
