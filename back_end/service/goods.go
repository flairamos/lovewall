package service

import (
	"errors"
	"log"
	"shoollove/enum"
	"shoollove/model"
	"shoollove/pkg/db"
	"shoollove/pkg/tool"
	"slices"
	"time"
)

// GoodsAdd 添加
func GoodsAdd(req model.GoodsAdd, userInfo model.Account) (resp any, err error) {
	if req.Title == "" || req.Category == "" || req.Images == "" || req.Price == "" {
		return nil, errors.New("缺少必要参数")
	}
	param := model.Goods{
		Title:         req.Title,
		Description:   req.Description,
		Images:        req.Images,
		Category:      req.Category,
		Price:         req.Price,
		OriginalPrice: req.OriginalPrice,
		Status:        enum.GOODS_SOLD,
		UserID:        int(userInfo.ID),
	}
	// 调用支付

	err = db.Gorm.Model(&model.Goods{}).Create(&param).Error
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GoodsEdit 编辑
func GoodsEdit(req model.Goods) (resp any, err error) {
	//session := db.Gorm
	//
	//param := model.Goods{}
	//
	//// var goods goods
	//// 请求参数处理后添加到数据库
	//// 赋值操作
	//err = session.Updates(&param).Error
	//if err != nil {
	//    log.ZapSqlLog().Error(context.Background(), "UserUpdate", err.Error())
	//    return resp, serr.Serr{Code: serr.Fail, Msg: err.Error()}
	//}
	return resp, nil
}

// GoodsDelete 删除
func GoodsDelete(req model.Goods) (resp any, err error) {
	//session := db.Gorm
	//
	//param := model.Goods{}
	//
	//// var goods goods
	//// 请求参数处理后添加到数据库
	//// 赋值操作
	//err = session.Delete(&param).Error
	//if err != nil {
	//	log.ZapSqlLog().Error(context.Background(), "UserDelete", err.Error())
	//	return resp, serr.Serr{Code: serr.Fail, Msg: err.Error()}
	//}
	return resp, nil
}

// GoodsGet 查询
func GoodsGet(req model.GoodsGet) (resp any, err error) {
	session := db.Gorm.Model(&model.Goods{})

	var (
		param  = make([]model.Goods, 0)
		result = make([]model.GoodsList, 0)
	)

	if req.Title != "" {
		session = session.Where("title like ?", "%"+req.Title+"%")
	}
	if !slices.Contains(enum.GOODS_TYPE_LIST, req.Category) || req.Category == enum.GOODS_ALL {
	} else {
		session = session.Where("category = ?", req.Category)
	}
	session.Order("created_at desc")
	pageData, err := tool.IPage(req.Page, req.Limit, session, param)
	if err != nil {
		log.Println("GoodsGet error :", err)
		return resp, err
	}
	data := pageData.Data.([]model.Goods)
	for _, goods := range data {
		user, uerr := UserInfoById(goods.UserID)
		if uerr != nil {
			log.Println("GoodsGet User error :", err)
			continue
		}
		result = append(result, model.GoodsList{
			Title:         goods.Title,
			Description:   goods.Description,
			Images:        goods.Images,
			Category:      enum.GOODS_TYPE_MAP[goods.Category],
			Price:         goods.Price,
			OriginalPrice: goods.OriginalPrice,
			PublishTime:   tool.FormatTimeInterval(goods.CreatedAt, time.Now()),
			Status:        goods.Status,
			UserID:        goods.UserID,
			Seller: model.AccountAvatar{
				ID:     user.ID,
				Phone:  user.Phone,
				Name:   user.Name,
				Avatar: user.Avatar,
			},
		})
	}
	pageData.Data = result
	return pageData, nil
}

// GoodsForm 查询
func GoodsForm(req model.Goods) (resp any, err error) {
	//session := db.Gorm
	//
	//param := model.Goods{}
	//
	//// var goods goods
	//// 请求参数处理后添加到数据库
	//// 赋值操作
	//err = session.Take(&param, req.Id).Error
	//if err != nil {
	//	log.ZapSqlLog().Error(context.Background(), "UserDelete", err.Error())
	//	return resp, serr.Serr{Code: serr.Fail, Msg: err.Error()}
	//}
	return resp, nil
}
