package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameNotice = "notice"

// Notice mapped from table <notice>
type Notice struct {
	ID            int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Phone         string         `gorm:"column:phone;not null;comment:手机号" json:"phone"`                   // 手机号
	Title         string         `gorm:"column:title;not null;comment:标题" json:"title"`                    // 标题
	Content       string         `gorm:"column:content;not null;comment:内容" json:"content"`                // 内容
	Attachments   string         `gorm:"column:attachments;not null;comment:附件" json:"attachments"`        // 附件
	Likers        string         `gorm:"column:likers;not null;comment:点赞数" json:"likers"`                 // 点赞数
	ViewNumber    int64          `gorm:"column:view_number;not null;comment:浏览数" json:"view_number"`       // 浏览数
	CommentNumber int64          `gorm:"column:comment_number;not null;comment:评论数" json:"comment_number"` // 评论数
	Status        int16          `gorm:"column:status;not null;comment:状态" json:"status"`                  // 状态
	CreatedAt     time.Time      `gorm:"column:created_at;comment:创建时间" json:"created_at"`                 // 创建时间
	UpdatedAt     time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`                 // 更新时间
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                 // 删除时间
}

// TableName Notice's table name
func (*Notice) TableName() string {
	return TableNameNotice
}

type NoticeAddReq struct {
	Phone       string `gorm:"column:phone;not null;comment:手机号" json:"phone,omitempty"`            // 手机号
	Title       string `gorm:"column:title;not null;comment:标题" json:"title,omitempty"`             // 标题
	Content     string `gorm:"column:content;not null;comment:内容" json:"content,omitempty"`         // 内容
	Attachments string `gorm:"column:attachments;not null;comment:附件" json:"attachments,omitempty"` // 附件
}

type NoticeResp struct {
	ID            int32  `json:"id"`
	Phone         string `json:"phone"`          // 手机号
	Title         string `json:"title"`          // 标题
	Content       string `json:"content"`        // 内容
	Attachments   string `json:"attachments"`    // 附件
	LikeNumber    int64  `json:"like_number"`    // 点赞数
	ViewNumber    int64  `json:"view_number"`    // 浏览数
	CommentNumber int64  `json:"comment_number"` // 评论数
	Status        int16  `json:"status"`         // 状态
	IsLike        bool   `json:"is_like"`        // 是否点赞
	CreatedAt     string `json:"created_at"`     // 创建时间
	UpdatedAt     string `json:"updated_at"`     // 更新时间
	DeletedAt     string `json:"deleted_at"`     // 删除时间
}
