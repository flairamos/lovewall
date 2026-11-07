package model

// Posting 模型
import (
	"time"

	"gorm.io/gorm"
)

const TableNamePosting = "posting"

type Posting struct {
	ID         int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Phone      string         `gorm:"column:phone;not null;comment:手机号" json:"phone"` // 手机号
	UserId     int32          `gorm:"column:user_id;not null;comment:用户ID" json:"user_id"`
	Type       int16          `gorm:"column:type;not null;comment:帖子类型" json:"type"`                    // 帖子类型
	Content    string         `gorm:"column:content;not null;comment:帖子内容" json:"content"`              // 帖子内容
	Photos     string         `gorm:"column:photos;not null;comment:帖子图片" json:"photos"`                // 帖子图片
	RefComment int32          `gorm:"column:ref_comment;not null;ref_comment:父评论id" json:"ref_comment"` // 父亲评论id
	LikerId    string         `gorm:"column:liker_id;not null;comment:liker_id" json:"liker_id"`        // 帖子点赞用户的id集合
	CreatedAt  time.Time      `gorm:"column:created_at;comment:创建时间" json:"created_at"`                 // 创建时间
	UpdatedAt  time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`                 // 更新时间
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                 // 删除时间
}

// TableName Posting's table name
func (*Posting) TableName() string {
	return TableNamePosting
}

type PostingAdd struct {
	Phone      string   `gorm:"column:phone;not null;comment:手机号" json:"phone"` // 手机号
	UserId     int32    `gorm:"column:user_id;not null;comment:用户ID" json:"user_id"`
	Type       int16    `gorm:"column:type;not null;comment:帖子类型" json:"type"`                    // 帖子类型
	Content    string   `gorm:"column:content;not null;comment:帖子内容" json:"content"`              // 帖子内容
	Photos     []string `gorm:"column:photos;not null;comment:帖子图片" json:"photos"`                // 帖子图片
	RefComment int32    `gorm:"column:ref_comment;not null;ref_comment:父评论id" json:"ref_comment"` // 父亲评论id
}

// 帖子请求参数
type PostGet struct {
	Name     string `json:"name" query:"name"`
	Phone    string `json:"phone" query:"phone"`
	Type     int    `json:"type" query:"type"`           // 0-话题 1-表白墙 2-晒图
	LastTime string `json:"last_time" query:"last_time"` // 最新时间
	PageReq
}

type PostingForm struct {
	Id   int `json:"id"`
	Type int `json:"type"` // 0-话题 1-表白墙 2-晒图
	PageReq
}

type PostingLike struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
}

type PostingDelete struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
}

type PostingTree struct {
	ID                 int32           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Phone              string          `gorm:"column:phone;not null;comment:手机号" json:"phone"`                              // 手机号
	Name               string          `gorm:"column:name;not null;comment:昵称" json:"name"`                                 // 昵称
	Avatar             string          `gorm:"column:avatar;not null;avatar:昵称" json:"avatar"`                              // 用户头像
	Authentication     bool            `gorm:"column:authentication;not null;comment:是否认证" json:"authentication"`           // 是否认证
	AuthenticationName int32           `gorm:"column:authentication_name;not null;comment:认证信息" json:"authentication_name"` // 认证信息
	Type               int16           `gorm:"column:type;not null;comment:帖子类型" json:"type"`                               // 帖子类型
	Content            string          `gorm:"column:content;not null;comment:帖子内容" json:"content"`                         // 帖子内容
	Photos             []string        `gorm:"column:photos;not null;comment:帖子图片" json:"photos"`                           // 帖子图片
	RefComment         int32           `gorm:"column:ref_comment;not null;ref_comment:评论数" json:"ref_comment"`              // 父亲评论id
	IsLike             bool            `gorm:"-" json:"is_like"`
	CanDelete          bool            `gorm:"-" json:"can_delete"`
	CreatedAt          string          `gorm:"column:created_at;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt          string          `gorm:"column:updated_at;comment:更新时间" json:"updated_at"` // 更新时间
	DeletedAt          string          `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"` // 删除时间
	Liker              []AccountAvatar `gorm:"-" json:"liker"`
	Children           []PostingTree   `gorm:"-" json:"children"`
	AllComment         int32           `gorm:"-" json:"all_comment"` // 所有评论数量
}
