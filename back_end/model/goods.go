package model

import (
	"time"
)

const TableNameGoods = "goods"

// Goods 模型
type Goods struct {
	ID            int       `gorm:"id" json:"id"`
	Title         string    `gorm:"title" json:"title"`
	Description   string    `gorm:"description" json:"description"`
	Images        string    `gorm:"images" json:"images"`
	Category      string    `gorm:"category" json:"category"`
	Price         string    `gorm:"price" json:"price"`
	OriginalPrice string    `gorm:"original_price" json:"original_price"`
	PublishTime   string    `gorm:"publish_time" json:"publish_time"`
	Status        string    `gorm:"status" json:"status"`
	UserID        int       `gorm:"user_id" json:"user_id"`
	CreatedAt     time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"updated_at" json:"updated_at"`
	DeletedAt     time.Time `gorm:"deleted_at" json:"deleted_at"`
}

func (*Goods) TableName() string {
	return TableNameGoods
}

type GoodsList struct {
	ID            int           `gorm:"id" json:"id"`
	Title         string        `gorm:"title" json:"title"`
	Description   string        `gorm:"description" json:"description"`
	Images        string        `gorm:"images" json:"images"`
	Category      string        `gorm:"category" json:"category"`
	Price         string        `gorm:"price" json:"price"`
	OriginalPrice string        `gorm:"original_price" json:"original_price"`
	PublishTime   string        `gorm:"publish_time" json:"publish_time"`
	Status        string        `gorm:"status" json:"status"`
	UserID        int           `gorm:"user_id" json:"user_id"`
	Seller        AccountAvatar `gorm:"-" json:"seller"`
}

type GoodsAdd struct {
	ID            int    `gorm:"id" json:"id"`
	Title         string `gorm:"title" json:"title"`
	Description   string `gorm:"description" json:"description"`
	Images        string `gorm:"images" json:"images"`
	Category      string `gorm:"category" json:"category"`
	Price         string `gorm:"price" json:"price"`
	OriginalPrice string `gorm:"original_price" json:"original_price"`
}

type GoodsGet struct {
	Title    string `query:"title" json:"title"`
	Category string `query:"category" json:"category"`
	PageReq
}
