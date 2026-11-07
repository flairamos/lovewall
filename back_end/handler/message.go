package handler

import (
	"errors"
	"log"
	"shoollove/model"
	"shoollove/pkg/db"
	jwtutil "shoollove/pkg/jwt_util"
	"shoollove/pkg/respond"
	"shoollove/pkg/tool"
	"sync"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

// 定义一个用户连接映射，用于存储在线每个用户的连接
var clients = make(map[int32]*websocket.Conn)
var clientsMutex sync.Mutex

func MessageSocket(c *websocket.Conn) {
	user, err := jwtutil.WSUserInfo(c)
	if err != nil {
		log.Println("MessageSocket jwt err:", err)
		return
	}
	// 保留在线状态
	clientsMutex.Lock()
	clients[user.ID] = c
	clientsMutex.Unlock()

	// 断开时清理
	defer func() {
		clientsMutex.Lock()
		delete(clients, user.ID)
		clientsMutex.Unlock()
	}()

	// 处理消息
	for {
		var (
			req     model.MessageAdd
			message model.Message
		)
		if err = c.ReadJSON(&req); err != nil {
			break
		}
		message = model.Message{
			SenderID:    req.SenderID,
			ReceiverID:  req.ReceiverID,
			GroupID:     req.GroupID,
			MessageText: req.MessageText,
		}
		// 私信
		if req.Type == "private" {
			clientsMutex.Lock()
			targetConn, exist := clients[req.ReceiverID]
			clientsMutex.Unlock()
			if exist {
				targetConn.WriteJSON(req)
				message.Status = "online"
				message.Type = "private"
			} else {
				message.Status = "offline"
				message.Type = "group"
			}
			err = db.Gorm.Table("message").Create(&message).Error
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func MessagePushNew(c *websocket.Conn) {
	user, err := jwtutil.WSUserInfo(c)
	if err != nil {
		log.Println("MessagePushNew jwt err:", err)
		return
	}

	lastTime := c.Query("last_time")
	// 保留在线状态
	clientsMutex.Lock()
	clients[user.ID] = c
	clientsMutex.Unlock()

	// 断开时清理
	defer func() {
		clientsMutex.Lock()
		delete(clients, user.ID)
		clientsMutex.Unlock()
	}()
	// 处理消息，推送每个人给该用户发送的最新消息
	for {
		var (
			message    []model.Message
			messageNew []model.MessageNew
		)
		clientsMutex.Lock()
		targetConn := clients[user.ID]
		clientsMutex.Unlock()
		sql := `WITH RankedMessages AS (
					SELECT 
						id,
						sender_id,
						receiver_id,
						group_id,
						message_text,
						created_at,
						updated_at,
						deleted_at,
						status,
						ROW_NUMBER() OVER (PARTITION BY receiver_id, sender_id ORDER BY created_at DESC) as rn
					FROM 
						public.message
				)
				SELECT 
					id,
					sender_id,
					receiver_id,
					group_id,
					message_text,
					created_at,
					updated_at,
					deleted_at,
					status
				FROM 
					RankedMessages
				WHERE 
					rn = 1
				AND receiver_id = ?
				AND created_at BETWEEN ? AND ?
				`
		if lastTime == "" {
			targetConn.WriteJSON(fiber.Map{
				"code":  false,
				"msg":   "缺少时间参数",
				"error": err.Error(),
			})
		}
		err = db.Gorm.Raw(sql, user.ID, lastTime, tool.TimeStr(time.Now())).Scan(&message).Error
		if err != nil {
			log.Println("MessagePushNew push message err:", err)
			targetConn.WriteJSON(fiber.Map{
				"code":  false,
				"msg":   "消息推送失败",
				"error": err.Error(),
			})
			return
		}
		for _, m := range message {
			// 查询发送着信息
			info, uerr := UserInfoById(int(m.SenderID))
			if uerr != nil {
				log.Println("MessagePushNew user info err:", uerr)
				continue
			}
			// 查询未读信息
			var count int64
			err = db.Gorm.Table("message").
				Where("created_at between ? and ?", lastTime, tool.TimeStr(time.Now())).
				Where("receiver_id = ? and sender_id = ?", m.ReceiverID, m.SenderID).
				Count(&count).Error
			if err != nil {
				log.Println("MessagePushNew count err:", err)
				continue
			}
			lastTime = tool.TimeStr(time.Now())
			messageNew = append(messageNew, model.MessageNew{
				Id:          m.SenderID, // 该id是用户id用于进入私聊界面
				Name:        info.Name,
				Avatar:      info.Avatar,
				LastMessage: m.MessageText,
				Time:        tool.FormatTimeInterval(m.CreatedAt, time.Now()),
				UnreadCount: int32(count),
			})
		}
		if len(messageNew) > 0 {
			targetConn.WriteJSON(messageNew)
		}
	}
}

func SendMessage(c *fiber.Ctx) error {
	var (
		req     model.MessageAdd
		message model.Message
	)
	if err := c.BodyParser(&req); err != nil {
		return respond.Fail(c, "发送失败", err)
	}
	message = model.Message{
		SenderID:    req.SenderID,
		ReceiverID:  req.ReceiverID,
		GroupID:     req.GroupID,
		MessageText: req.MessageText,
	}
	message.Status = "online"
	message.Type = "private"
	err := db.Gorm.Table("message").Create(&message).Error
	if err != nil {
		log.Println(err)
		return respond.Fail(c, "发送失败", err)
	}
	return respond.Ok(c, "")
}

func ClearUnreadMessage(c *fiber.Ctx) error {
	var req model.DeleteMessage
	if err := c.BodyParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}
	id := c.QueryInt("user_id", 0)
	if id == 0 {
		return respond.Fail(c, "参数错误", errors.New("缺少必要参数"))
	}
	if len(req.Ids) == 0 {
		return respond.Fail(c, "", errors.New("没有修改参数"))
	}
	if err := db.Gorm.Model(&model.Message{}).Where("id in ?", req.Ids).Update("unread", "2").Error; err != nil {
		log.Println(err)
		return respond.Fail(c, "更新未读失败", err)
	}
	return respond.Ok(c, "")
}

func GetMessage(c *fiber.Ctx) error {
	userInfo, err := jwtutil.UserInfo(c)
	if err != nil {
		return respond.Ok(c, "", "获取登录状态失败")
	}
	var (
		message    []model.Message
		messageNew []model.MessageNew
	)
	sql := `WITH RankedMessages AS (
					SELECT 
						id,
						sender_id,
						receiver_id,
						group_id,
						message_text,
						created_at,
						updated_at,
						deleted_at,
						status,
						unread,
						ROW_NUMBER() OVER (PARTITION BY receiver_id, sender_id ORDER BY created_at DESC) as rn
					FROM 
						public.message
				)
				SELECT 
					id,
					sender_id,
					receiver_id,
					group_id,
					message_text,
					created_at,
					updated_at,
					deleted_at,
					status,
				    unread
				FROM 
					RankedMessages
				WHERE 
					rn = 1
				AND receiver_id = ?
				AND unread = ?
				`
	err = db.Gorm.Raw(sql, userInfo.ID, "1").Scan(&message).Error
	if err != nil {
		log.Println("get message err:", err)
		return err
	}
	for _, m := range message {
		// 查询发送着信息
		info, uerr := UserInfoById(int(m.SenderID))
		if uerr != nil {
			log.Println("MessagePushNew user info err:", uerr)
			continue
		}
		// 查询未读信息
		// 未读消息根据根据更新时间判断，
		var count int64
		err = db.Gorm.Table("message").
			Where("created_at < ? and unread = ?", tool.TimeStr(m.CreatedAt), "1").
			Where("receiver_id = ? and sender_id = ?", m.ReceiverID, m.SenderID).
			Count(&count).Error
		if err != nil {
			log.Println("MessagePushNew count err:", err)
			continue
		}
		messageNew = append(messageNew, model.MessageNew{
			Id:          m.SenderID, // 该id是用户id用于进入私聊界面
			Name:        info.Name,
			Avatar:      info.Avatar,
			LastMessage: m.MessageText,
			Time:        tool.FormatTimeInterval(m.CreatedAt, time.Now()),
			UnreadCount: int32(count),
		})
	}
	return respond.Ok(c, messageNew, "查询成功")
}

func GetUnReadMessage(c *fiber.Ctx) error {
	var req model.UnReadMessage
	if err := c.QueryParser(&req); err != nil {
		return respond.Fail(c, "参数错误", err)
	}
	userInfo, err := jwtutil.UserInfo(c)
	if err != nil {
		return respond.Fail(c, "查询用户信息失败", err)
	}
	var list []model.Message
	session := db.Gorm.Table("message").
		Where("receiver_id in ? and sender_id in ? and created_at > ?", []int32{userInfo.ID, req.OpponentId}, []int32{userInfo.ID, req.OpponentId}, req.LastMessageTime)
	page, err := tool.IPage(req.Page, req.Limit, session, list)
	if err != nil {
		log.Println("查询未读信息失败err:", err)
		return respond.Fail(c, "查询失败", err)
	}
	var result []model.MessageHistory
	var lastShowTime time.Time = tool.StrTime(req.LastMessageTime)
	for _, message := range page.Data.([]model.Message) {
		isSelf := message.SenderID == userInfo.ID
		showTime := message.CreatedAt.Sub(lastShowTime) > 5*time.Minute
		if showTime {
			lastShowTime = message.CreatedAt
		}

		result = append(result, model.MessageHistory{
			Id:       message.ID,
			Content:  message.MessageText,
			Time:     tool.TimeStr(message.CreatedAt),
			IsSelf:   isSelf,
			ShowTime: showTime,
		})
	}
	return respond.Ok(c, result)
}
