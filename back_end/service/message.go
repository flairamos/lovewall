package service

import (
	"github.com/gofiber/contrib/websocket"
	"shoollove/model"
	"shoollove/pkg/db"
	"shoollove/pkg/tool"
)

// 写入消息
func writeMessage(c *websocket.Conn) {
	var req model.MessageAdd
	err := c.ReadJSON(&req)
	if err != nil {
	}
}

// 查询用户的最新消息

func MessageNewGet(req model.MessageNewReq) []model.MessageNew {
	var message []model.Message
	err := db.Gorm.Table("message").
		Where("receiver_id = ?", req.ReceiverId).
		Order("created_at desc").Limit(1).Find(message).Error
	if err != nil {
		return []model.MessageNew{}
	}
	var results []model.MessageNew
	for _, m := range message {
		// 查询用户名称
		user, err1 := UserInfoById(int(m.ID))
		if err1 != nil {
			continue
		}
		// 查询是否还有未读消息
		var msgCount int64
		err = db.Gorm.Table("message").Where("receiver_id = ?", req.ReceiverId).
			Where("created_at between ? and ?", req.LastMessageTime, tool.TimeStr(m.CreatedAt)).
			Count(&msgCount).Error
		if err != nil {
			msgCount = 0
		}
		results = append(results, model.MessageNew{
			Id:          m.ID,
			Name:        user.Name,
			Avatar:      user.Avatar,
			LastMessage: m.MessageText,
			Time:        tool.TimeStr(m.CreatedAt),
			UnreadCount: int32(msgCount),
		})
	}
	return results
}
