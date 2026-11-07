package tool

import (
	"math"

	"gorm.io/gorm"
)

// 分页组件
// 除gorm其他类型均为非指针类型
// currentPage - 当前页  perPageSize - 每页大小 param 查询的结构体
// param不传地址，因为已经做了取地址处理
func IPage(currentPage, perPageSize int, db *gorm.DB, param any) (*dataPage, error) {
	var offset, limit int
	if currentPage <= 1 {
		currentPage = 1
	}
	if perPageSize <= 0 {
		perPageSize = 10
	}
	offset = (currentPage - 1) * perPageSize
	limit = perPageSize

	var count int64
	err := db.Count(&count).Offset(offset).Limit(limit).
		Find(&param).Error
	if err != nil {
		return nil, err
	}

	pageNum := int(math.Ceil(float64(count) / float64(perPageSize)))
	result := dataPage{
		Data:        param,
		CurrentPage: currentPage,
		PerPageSize: perPageSize,
		Total:       count,
		PageNum:     pageNum,
		HasPrx: func() bool {
			if currentPage == 1 {
				return false
			}
			return true
		}(),
		HasNext: func() bool {
			if currentPage == pageNum {
				return false
			}
			return true
		}(),
	}
	return &result, nil
}

type dataPage struct {
	Data        any   `json:"data"`
	CurrentPage int   `json:"current_page"`
	PerPageSize int   `json:"per_page_size"`
	Total       int64 `json:"total"`
	PageNum     int   `json:"page_num"`
	HasPrx      bool  `json:"has_prx"`
	HasNext     bool  `json:"has_next"`
}

func PageParam(currentPage, perPageSize int) (int, int) {
	var offset, limit int
	if currentPage <= 1 {
		currentPage = 1
	}
	if perPageSize <= 0 {
		perPageSize = 10
	}
	offset = (currentPage - 1) * perPageSize
	limit = perPageSize
	return offset, limit
}

func SimplePage(data any, currentPage, perPageSize int, total int64) *dataPage {
	if currentPage <= 1 {
		currentPage = 1
	}
	if perPageSize <= 0 {
		perPageSize = 10
	}
	pageNum := int(math.Ceil(float64(total) / float64(perPageSize)))
	result := dataPage{
		Data:        data,
		CurrentPage: currentPage,
		PerPageSize: perPageSize,
		Total:       total,
		PageNum:     pageNum,
		HasPrx: func() bool {
			if currentPage == 1 {
				return false
			}
			return true
		}(),
		HasNext: func() bool {
			if currentPage == pageNum {
				return false
			}
			return true
		}(),
	}
	return &result
}
