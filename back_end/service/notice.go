package service

import (
	"shoollove/enum"
	"shoollove/model"
	"shoollove/pkg/db"
	"shoollove/pkg/tool"
	"slices"
	"time"
)

// TopicAdd 添加
func TopicAdd(req model.NoticeAddReq) (resp any, err error) {

	param := model.Notice{
		Phone:       req.Phone,
		Title:       req.Title,
		Content:     req.Content,
		Attachments: req.Attachments,
		Status:      1,
		CreatedAt:   time.Now(),
	}
	err = db.Gorm.Create(&param).Error

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// TopicEdit 编辑
func TopicEdit(req model.Notice) (resp any, err error) {
	session := db.Gorm

	param := model.Notice{}

	// var topic topic
	// 请求参数处理后添加到数据库
	// 赋值操作
	err = session.Updates(&param).Error
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// TopicDelete 删除
func TopicDelete(req model.Notice) (resp any, err error) {
	session := db.Gorm

	param := model.Notice{}

	// var topic topic
	// 请求参数处理后添加到数据库
	// 赋值操作
	err = session.Delete(&param).Error
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// TopicGet 查询
func TopicGet(req model.PageReq, userInfo model.Account) (resp any, err error) {

	var (
		param    model.Notice
		resParam model.NoticeResp
		isLike   = false
		likerNum = 0
	)
	err = db.Gorm.Model(&model.Notice{}).Order("created_at desc").Limit(1).Find(&param).Error
	if err != nil {
		return resParam, err
	}
	// 查询是否点赞
	likers, err := tool.StringToIntSlice(param.Likers)
	if err != nil {
		return resParam, err
	}
	if slices.Contains(likers, int(userInfo.ID)) {
		isLike = true
	}
	// 查询总品论数
	var count int64
	err = db.Gorm.Table("posting").Where("type = ?", enum.CMT_TOPIC).Count(&count).Error
	if err != nil {
	}
	// 查询总点赞数
	slice, err := tool.StringToIntSlice(param.Likers)
	if err != nil {

	}
	likerNum = len(slice)
	resParam = model.NoticeResp{
		ID:            param.ID,
		Phone:         param.Phone,
		Title:         param.Title,
		Content:       param.Content,
		Attachments:   param.Attachments,
		LikeNumber:    int64(likerNum),
		ViewNumber:    param.ViewNumber,
		CommentNumber: count,
		Status:        param.Status,
		IsLike:        isLike,
		CreatedAt:     tool.TimeStr(param.CreatedAt),
		//UpdatedAt:     "",
		//DeletedAt:     "",
	}
	return resParam, nil
}

// TopicForm 查询
func TopicForm(req model.Notice) (resp any, err error) {
	//session := db.Gorm
	//
	//param := model.Topic{}
	//
	//// var topic topic
	//// 请求参数处理后添加到数据库
	//// 赋值操作
	//err = session.Take(&param, req.Id).Error
	//if err != nil {
	//	log.ZapSqlLog().Error(context.Background(), "UserDelete", err.Error())
	//	return resp, serr.Serr{Code: serr.Fail, Msg: err.Error()}
	//}
	return resp, nil
}

func TopicPraise(id int, user model.Account) (resp any, err error) {
	var param model.Notice
	err = db.Gorm.Table("notice").Where("id = ?", id).First(&param).Error
	if err != nil {
		return resp, err
	}
	var likers []int
	if param.Likers == "" {
		err = db.Gorm.Table("notice").Where("id = ?", id).Update("likers", tool.IntSliceToString([]int{int(user.ID)})).Error
		if err != nil {
			return resp, err
		}
		return resp, nil
	} else {
		likers, err = tool.StringToIntSlice(param.Likers)
		if err != nil {
			return resp, err
		}
		var flag = false
		var index int
		for i, v := range likers {
			if v == int(user.ID) {
				flag = true
				index = i
			}
		}
		if flag {
			likers = append(likers[:index], likers[index+1:]...)
		} else {
			likers = append(likers, int(user.ID))
		}
		err = db.Gorm.Table("notice").Where("id = ?", id).Update("likers", tool.IntSliceToString(likers)).Error
		if err != nil {
			return resp, err
		}
		return resp, nil
	}
}
