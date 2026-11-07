package service

import (
	"gorm.io/gorm"
	"shoollove/model"
	"shoollove/pkg/db"
)

func FollowAdd(req model.FollowAdd, userInfo model.Account) (resp any, err error) {
	var param = model.Follow{
		CollectPhone: req.CollectPhone,
		CollectId:    req.CollectId,
		UserId:       userInfo.ID,
		Phone:        userInfo.Phone,
	}
	err = db.Gorm.Table("follow").Create(&param).Error
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func FollowDelete(req model.FollowAdd, userInfo model.Account) (resp any, err error) {
	err = db.Gorm.Unscoped().Table("follow").
		Where("phone = ?", userInfo.Phone).
		Where("collect_phone = ?", req.CollectPhone).
		Delete(&model.Follow{}).Error
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func FollowGet(req model.FollowAdd, userInfo model.Account) (resp any, err error) {
	var count int64
	err = db.Gorm.Table("follow").Where("phone = ?", userInfo.Phone).
		Where("collect_phone = ?", req.CollectPhone).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count < 1 {
		return false, nil
	}
	return true, nil
}

func FollowInfo(req model.FollowInfo) (resp any, err error) {
	var (
		follows, fans, likes int64
	)
	errTx := db.Gorm.Transaction(func(tx *gorm.DB) error {
		err = tx.Table("follow").
			Select("count(*) as follows").
			Where("phone = ?", req.Phone).
			Scan(&follows).Error
		if err != nil {
			return err
		}
		err = tx.Table("follow").
			Select("count(*) as fans").
			Where("collect_phone = ?", req.Phone).
			Scan(&fans).Error
		if err != nil {
			return err
		}
		err = tx.Table("selfie").Select("sum(like_num) as likes").
			Where("phone = ?", req.Phone).Scan(&likes).Error
		if err != nil {
			return err
		}
		return nil
	})
	if errTx != nil {
		return nil, errTx
	}
	resp = map[string]interface{}{
		"follows": follows,
		"fans":    fans,
		"likes":   likes,
	}
	return resp, nil
}
