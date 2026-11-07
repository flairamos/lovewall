package service

import (
	"log"
	"shoollove/model"
	"shoollove/pkg/db"
	"shoollove/pkg/tool"
	"slices"
)

// 判断该账户是否收藏
func IsLike(userId int, collectId int32) bool {
	var post model.Posting
	err := db.Gorm.Table("posting").Where("id = ?", collectId).First(&post).Error
	if err != nil {
		log.Printf("用户Id为%d查询查询点赞信息失败:%v", userId, err)
		return false
	}
	likers, err := tool.StringToIntSlice(post.LikerId)
	if err != nil {
		return false
	}
	if slices.Contains(likers, userId) {
		return true
	}
	return false
}

// 传入帖子ID查询点赞用户信息
func FindLiker(collectId int32) ([]model.AccountAvatar, error) {
	var (
		likerId string
		acc     = make([]model.AccountAvatar, 0)
	)
	// 查询点赞人信息
	err := db.Gorm.Table("posting").
		Select("liker_id").
		Where("id = ?", collectId).
		Scan(&likerId).Error
	if err != nil {
		log.Println("FindLiker err 查询帖子点赞信息失败:", err)
	}
	// 查询点赞人详细信息
	idSlice, err := tool.StringToIntSlice(likerId)
	if err != nil {
		log.Println("FindLiker err 解析帖子点赞信息失败:", err)
	}
	for _, liker := range idSlice {
		tmpAcc := model.AccountAvatar{}
		err = db.Gorm.Table("account").Select("id,name, phone ,avatar").
			Where("id = ?", liker).
			First(&tmpAcc).Error
		if err != nil {
			continue
		}
		acc = append(acc, model.AccountAvatar{
			ID:     tmpAcc.ID,
			Phone:  tmpAcc.Phone,
			Name:   tmpAcc.Name,
			Avatar: tmpAcc.Avatar,
		})
	}
	return acc, nil
}

func UserInfo(phone string) (model.Account, error) {
	var poster model.Account
	err := db.Gorm.Table("account").Select("id,name,avatar,authentication,authentication_name").
		Where("phone = ?", phone).
		Find(&poster).Error
	if err != nil {
		return model.Account{}, err
	}
	return poster, nil
}

func UserInfoById(id int) (model.Account, error) {
	var poster model.Account
	err := db.Gorm.Table("account").Select("id,name,avatar,authentication,authentication_name").
		Where("id = ?", id).
		Find(&poster).Error
	if err != nil {
		return model.Account{}, err
	}
	return poster, nil
}

// 查询点赞最多的三条评论
func ThreeRankComments(postingId int32) []model.PostingTree {
	var (
		comments       []model.Posting
		resultComments = make([]model.PostingTree, 0)
	)
	err := db.Gorm.Table("posting").
		Where("ref_comment = ?", postingId).
		Order("LENGTH(liker_id)").
		Limit(3).
		Find(&comments).Error
	if err != nil {
		log.Println("ThreeRankComments err :查询最热三个帖子出错！", err)
	}
	for _, comment := range comments {
		info, uerr := UserInfo(comment.Phone)
		if uerr != nil {
			log.Println("ThreeRankComments err:查询用户信息失败")
		}
		tmp := model.PostingTree{
			ID:                 comment.ID,
			Phone:              comment.Phone,
			Name:               info.Name,
			Avatar:             info.Avatar,
			Authentication:     info.Authentication,
			AuthenticationName: info.AuthenticationName,
			Type:               comment.Type,
			Content:            comment.Content,
			Photos:             tool.StringToStringSlice(comment.Photos),
			RefComment:         comment.RefComment,
			CreatedAt:          tool.TimeStr(comment.CreatedAt),
		}
		resultComments = append(resultComments, tmp)
	}
	return resultComments
}

func FindAllCommentNum(postingId int32) int32 {
	var count int64
	err := db.Gorm.Table("posting").
		Select("count(*) as count").
		Where("ref_comment = ?", postingId).Scan(&count).Error
	if err != nil {
		return 0
	}
	return int32(count)
}
