package service

import (
	"errors"
	"log"
	"shoollove/enum"
	"shoollove/model"
	"shoollove/pkg/db"
	"shoollove/pkg/tool"
	"slices"
)

// PostingAdd 添加
func PostingAdd(req model.PostingAdd) (resp any, err error) {
	param := model.Posting{
		Phone:      req.Phone,
		UserId:     req.UserId,
		Type:       req.Type,
		Content:    req.Content,
		Photos:     tool.StringSliceToString(req.Photos),
		RefComment: req.RefComment,
	}

	err = db.Gorm.Table("posting").Create(&param).Error
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// PostingEdit 编辑
func PostingEdit(req model.Posting) (resp any, err error) {
	session := db.Gorm

	param := model.Posting{}

	// var posting posting
	// 请求参数处理后添加到数据库
	// 赋值操作
	err = session.Updates(&param).Error
	if err != nil {

		return resp, err
	}
	return resp, nil
}

// PostingDelete 删除
func PostingDelete(req model.PostingDelete) (resp any, err error) {
	err = db.Gorm.Unscoped().Table("posting").Where("id = ?", req.Id).Delete(&model.Posting{}).Error
	if err != nil {
		return resp, err
	}
	return resp, nil

}

// PostingGet 查询
// 查询帖子列表
func PostingGet(req model.PostGet, userInfo model.Account) (resp any, err error) {
	var (
		param  = make([]model.Posting, 0)
		result = make([]model.PostingTree, 0)
	)
	session := db.Gorm.Model(&model.Posting{})
	if req.Name != "" {
		session.Where("content like ?", "%"+req.Name+"%")
	}
	if req.Phone != "" {
		session.Where("phone = ?", req.Phone)
	}
	// 帖子类型
	if req.Type != enum.ALL {
		session.Where("type = ?", req.Type)
	}
	if req.LastTime != "" {
		session.Where("created_at < ?", req.LastTime)
	}
	offset, limit := tool.PageParam(req.Page, req.Limit)
	var total int64
	err = session.
		Where("ref_comment = ?", 0).
		Count(&total).
		Offset(offset).Limit(limit).
		Order("created_at desc").
		Find(&param).Error
	if err != nil {
		return resp, err
	}
	for _, posting := range param {
		// 查询点赞人信息
		liker, errlike := FindLiker(posting.ID)
		if errlike != nil {
			log.Println("----PostingGet err:查询点赞人信息失败----", errlike)
			continue
		}
		// 查询上传人信息
		poster, erruser := UserInfo(posting.Phone)
		if erruser != nil {
			log.Println("----PostingGet err:查询上传人信息失败----", err)
			continue
		}

		// 查询点赞最多的三条评论
		comments := ThreeRankComments(posting.ID)
		tmpNode := model.PostingTree{
			ID:                 posting.ID,
			Phone:              posting.Phone,
			Name:               poster.Name,
			Avatar:             poster.Avatar,
			Authentication:     poster.Authentication,
			AuthenticationName: poster.AuthenticationName,
			Type:               posting.Type,
			Content:            posting.Content,
			Photos:             tool.StringToStringSlice(posting.Photos),
			RefComment:         posting.RefComment,
			IsLike:             IsLike(int(userInfo.ID), posting.ID),                     // 是否收藏按钮
			CanDelete:          func() bool { return userInfo.Phone == posting.Phone }(), // 是否显示删除按钮
			CreatedAt:          tool.TimeStr(posting.CreatedAt),
			UpdatedAt:          tool.TimeStr(posting.UpdatedAt),
			Liker:              liker,
			Children:           comments,
			AllComment:         FindAllCommentNum(posting.ID),
		}
		result = append(result, tmpNode)
	}
	return tool.SimplePage(result, req.Page, req.Limit, total), nil
}

// PostingForm 查询
// 返回帖子的评论
func PostingForm(req model.PostingForm, userInfo model.Account) (resp any, err error) {
	var (
		param  = make([]model.Posting, 0)
		result = make([]model.PostingTree, 0)
	)
	session := db.Gorm.Model(&model.Posting{})
	if req.Id != 0 {
		session.Where("id = ?", req.Id)
	}
	session.Where("type = ?", req.Type)
	err = session.
		Where("ref_comment = ?", 0).
		Order("created_at desc").
		Find(&param).Error
	if err != nil {
		return resp, err
	}
	for _, posting := range param {
		// 查询点赞人信息
		liker, errlike := FindLiker(posting.ID)
		if errlike != nil {
			log.Println("----查询点赞人信息失败----", errlike)
			continue
		}
		// 查询上传人信息
		poster, erruser := UserInfo(posting.Phone)
		if erruser != nil {
			log.Println("----查询上传人信息失败----", err)
			continue
		}
		tmpNode := model.PostingTree{
			ID:                 posting.ID,
			Phone:              posting.Phone,
			Name:               poster.Name,
			Avatar:             poster.Avatar,
			Authentication:     poster.Authentication,
			AuthenticationName: poster.AuthenticationName,
			Type:               posting.Type,
			Content:            posting.Content,
			Photos:             tool.StringToStringSlice(posting.Photos),
			RefComment:         posting.RefComment,
			IsLike:             IsLike(int(poster.ID), posting.ID),                       // 是否收藏按钮
			CanDelete:          func() bool { return userInfo.Phone == posting.Phone }(), // 是否显示删除按钮
			CreatedAt:          tool.TimeStr(posting.CreatedAt),
			UpdatedAt:          tool.TimeStr(posting.UpdatedAt),
			// DeletedAt: posting.DeletedAt,
			Liker:    liker,
			Children: nil,
		}
		result = append(result, buildPostingTree(tmpNode))
	}
	return result, nil
}

func PostingPraise(req model.PostingLike, account model.Account) (resp any, err error) {
	// 查询点赞帖子的id
	var likerIds string
	err = db.Gorm.Table("posting").Select("liker_id").
		Where("id = ?", req.Id).Scan(&likerIds).Error
	if err != nil {
		log.Println("点赞查询点赞人ID失败")
		return nil, err
	}
	slice, err := tool.StringToIntSlice(likerIds)
	if err != nil {
		log.Println("转换点赞查询点赞人ID失败")
		return nil, err
	}
	// 判断是否点赞
	var flag bool = false
	var index int
	for i, uid := range slice {
		if uid == int(account.ID) {
			flag = true
			index = i
		}
	}
	if flag {
		slice = append(slice[:index], slice[index+1:]...)
	} else {
		slice = append(slice, int(account.ID))
	}
	newslice := tool.IntSliceToString(slice)
	err = db.Gorm.Table("posting").Where("id = ?", req.Id).Update("liker_id", newslice).Error
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// 查询评论区
func GetCommentById(req model.PostingForm, account model.Account) (resp any, err error) {
	if !slices.Contains(enum.POSTING_TYPE, req.Type) {
		return nil, errors.New("类型错误！")
	}
	var total int64
	session := db.Gorm.Table("posting").
		Where("ref_comment = ?", req.Id).
		Where("type = ?", req.Type).
		Count(&total)
	offset, limit := tool.PageParam(req.Page, req.Limit)
	var tmp []model.Posting
	err = session.Offset(offset).Limit(limit).Find(&tmp).Error
	if err != nil {
		return resp, err
	}
	var result = make([]model.PostingTree, 0)
	for i := 0; i < len(tmp); i++ {
		info, uerr := UserInfo(tmp[i].Phone)
		if uerr != nil {
			log.Println("查询评论区用户信息失败！")
			continue
		}
		result = append(result, buildPostingTree(model.PostingTree{
			ID:                 tmp[i].ID,
			Phone:              tmp[i].Phone,
			Name:               info.Name,
			Avatar:             info.Avatar,
			Authentication:     info.Authentication,
			AuthenticationName: info.AuthenticationName,
			Type:               int16(req.Type),
			Content:            tmp[i].Content,
			Photos:             tool.StringToStringSlice(tmp[i].Photos),
			RefComment:         tmp[i].RefComment,
			IsLike:             false,
			CanDelete:          false,
			CreatedAt:          tool.TimeStr(tmp[i].CreatedAt),
			//UpdatedAt:          "",
			// DeletedAt:          "",
			Liker:    nil,
			Children: nil,
		}))
	}
	return tool.SimplePage(result, req.Page, req.Limit, total), nil
}

func buildPostingTree(node model.PostingTree) model.PostingTree {
	children := getPostingChildren(node.ID)
	node.Children = children
	if len(children) == 0 {
		return node
	}
	for i := range children {
		children[i] = buildPostingTree(children[i])
	}
	return node
}

func getPostingChildren(nodeId int32) []model.PostingTree {
	var (
		param  []model.Posting
		result []model.PostingTree
	)
	err := db.Gorm.Table("posting").
		Where("ref_comment = ?", nodeId).
		Find(&param).Error
	if err != nil {
		return []model.PostingTree{}
	}
	for _, v := range param {
		// 查询点赞人信息
		likers, errlike := FindLiker(v.ID)
		if errlike != nil {
			log.Println("----查询点赞人信息失败----", errlike)
			continue
		}
		// 查询上传人信息
		// 查询上传人信息
		poster, erruser := UserInfo(v.Phone)
		if erruser != nil {
			log.Println("----查询上传人信息失败----", err)
			continue
		}
		tmp := model.PostingTree{
			ID:                 v.ID,
			Phone:              v.Phone,
			Name:               poster.Name,
			Avatar:             poster.Avatar,
			Authentication:     poster.Authentication,
			AuthenticationName: poster.AuthenticationName,
			Type:               v.Type,
			Content:            v.Content,
			Photos:             tool.StringToStringSlice(v.Photos),
			RefComment:         v.RefComment,
			IsLike:             IsLike(int(poster.ID), v.ID),
			CanDelete:          false,
			CreatedAt:          tool.TimeStr(v.CreatedAt),
			UpdatedAt:          tool.TimeStr(v.UpdatedAt),
			// DeletedAt:          v.DeletedAt,
			Liker:    likers,
			Children: nil,
		}
		result = append(result, tmp)
	}
	return result
}
