package service

import (
	"shoollove/model"
	"shoollove/pkg/db"
	"shoollove/pkg/tool"
)

// SelfieAdd 添加
func SelfieAdd(req model.SelfieAdd, userInfo model.Account) (resp any, err error) {
	param := model.Selfie{
		Phone:   userInfo.Phone,
		UserID:  userInfo.ID,
		LikeNum: 0,
		Photos:  tool.StringSliceToString(req.Photos),
		Title:   req.Title,
	}

	err = db.Gorm.Table("selfie").Create(&param).Error
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SelfieEdit 编辑
func SelfieEdit(req model.Selfie) (resp any, err error) {
	session := db.Gorm

	param := model.Selfie{}

	// var selfie selfie
	// 请求参数处理后添加到数据库
	// 赋值操作
	err = session.Updates(&param).Error
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// SelfieDelete 删除
func SelfieDelete(req model.Selfie) (resp any, err error) {
	session := db.Gorm

	param := model.Selfie{}

	// var selfie selfie
	// 请求参数处理后添加到数据库
	// 赋值操作
	err = session.Delete(&param).Error
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// SelfieGet 查询
func SelfieGet(req model.SelfieGetByType, user model.Account) (resp any, err error) {
	session := db.Gorm.Table("selfie")
	// if req.Name != "" {
	// 	session.Where("title like ?", "%"+req.Name+"%")
	// }
	if req.Type == 1 {
		// session.Order("like_num desc")
	} else if req.Type == 3 {
		slice, err := tool.StringToIntSlice(user.Following)
		if err != nil {
			return resp, err
		}
		session.Where("user_id in ?", slice)
	} else if req.Type == 2 {
		session.Order("like_num desc")
	}
	session.Order("created_at desc")
	var param = make([]model.Selfie, 0)
	page, err := tool.IPage(req.Page, req.Limit, session, param)
	if err != nil {
		return resp, err
	}
	var (
		result    model.SelfieList
		rightList = make([]model.SelfieListItem, 0)
		leftList  = make([]model.SelfieListItem, 0)
	)
	selfies := page.Data.([]model.Selfie)

	for i := 0; i < len(selfies); i++ {
		comUser, err := UserInfoById(int(selfies[i].UserID))
		if err != nil {
			continue
		}
		if (i+1)%2 == 0 {
			rightList = append(rightList, model.SelfieListItem{
				Id:      selfies[i].ID,
				LikeNum: selfies[i].LikeNum,
				Photos:  tool.StringToStringSlice(selfies[i].Photos),
				Poster: model.AccountAvatar{
					Avatar: comUser.Avatar,
					ID:     selfies[i].UserID,
					Name:   comUser.Name,
					Phone:  selfies[i].Phone,
				},
				Title: selfies[i].Title,
			})
		} else {
			leftList = append(leftList, model.SelfieListItem{
				Id:      selfies[i].ID,
				LikeNum: selfies[i].LikeNum,
				Photos:  tool.StringToStringSlice(selfies[i].Photos),
				Poster: model.AccountAvatar{
					Avatar: comUser.Avatar,
					ID:     selfies[i].UserID,
					Name:   comUser.Name,
					Phone:  selfies[i].Phone,
				},
				Title: selfies[i].Title,
			})
		}
	}
	result.LeftList = leftList
	result.RightList = rightList
	return result, nil
}

// SelfieForm 查询
func SelfieForm(req model.SelfieGetByType) (resp any, err error) {

	return resp, nil
}

func SelfiePraise(req model.SelfiePraise) (resp any, err error) {
	err = db.Gorm.Exec("UPDATE selfie SET like_num = like_num + 1 WHERE id = ?", req.Id).Error
	if err != nil {
		return nil, err
	}
	return nil, nil
}
