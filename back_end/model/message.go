package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameMessage = "message"

// Message mapped from table <message>
type Message struct {
	ID          int32          `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`               // 主键
	SenderID    int32          `gorm:"column:sender_id;not null;comment:发送人主键" json:"sender_id"`                   // 发送人主键
	ReceiverID  int32          `gorm:"column:receiver_id;comment:接收人主键" json:"receiver_id"`                        // 接收人主键
	GroupID     int32          `gorm:"column:group_id;comment:类型" json:"group_id"`                                 // 类型
	MessageText string         `gorm:"column:message_text;not null;comment:消息内容" json:"message_text"`              // 消息内容
	Status      string         `gorm:"column:status;comment:发送时是否在线" json:"status"`                                // 发送时是对方在线还是离线
	Type        string         `gorm:"column:type;comment:类型" json:"type"`                                         // 私信还是群组
	Unread      string         `gorm:"column:unread;comment:是否已读" json:"unread"`                                   // 是否已读
	CreatedAt   time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt   time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`                           // 更新时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                           // 删除时间
}

// TableName Message's table name
func (*Message) TableName() string {
	return TableNameMessage
}

type MessageAdd struct {
	ID          int32  `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`  // 主键
	SenderID    int32  `gorm:"column:sender_id;not null;comment:发送人主键" json:"sender_id"`      // 发送人主键
	ReceiverID  int32  `gorm:"column:receiver_id;comment:接收人主键" json:"receiver_id"`           // 接收人主键
	GroupID     int32  `gorm:"column:group_id;comment:类型" json:"group_id"`                    // 类型
	Type        string `gorm:"column:type;comment:类型" json:"type"`                            // 私信还是群组
	MessageText string `gorm:"column:message_text;not null;comment:消息内容" json:"message_text"` // 消息内容
}

type MessageList struct {
	ID          int32  `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"` // 主键
	SenderID    int32  `gorm:"column:sender_id;not null;comment:发送人主键" json:"sender_id"`     // 发送人主键
	ReceiverID  int32  `gorm:"column:receiver_id;comment:接收人主键" json:"receiver_id"`          // 接收人主键
	GroupID     int32  `gorm:"column:group_id;comment:类型" json:"group_id"`                   // 类型
	Status      string `gorm:"column:status;comment:发送时是否在线" json:"status"`
	MessageText string `gorm:"column:message_text;not null;comment:消息内容" json:"message_text"`              // 消息内容
	CreatedAt   string `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt   string `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`                           // 更新时间
}

// 消息列表返回最新消息
//
//	[{
//		id: 2,
//		name: '李宇',
//		avatar: '/images/avatar2.png',
//		lastMessage: '好的，那就这么说定了！',
//		time: '11:37',
//		unreadCount: 0
//	  }],
type MessageNew struct {
	Id          int32  `json:"id"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	LastMessage string `json:"lastMessage"`
	Time        string `json:"time"`
	UnreadCount int32  `json:"unreadCount"`
}

type MessageNewReq struct {
	Type            string `json:"type"` // // "private" | "group"
	MessageId       int32  `json:"messageId"`
	SenderId        int32  `json:"senderId"`
	ReceiverId      int32  `json:"receiverId"`
	GroupId         int32  `json:"groupId"`
	LastMessageTime string `json:"lastMessageTime"`
}

type UnReadMessage struct {
	OpponentId      int32  `json:"opponentId"`
	LastMessageTime string `json:"lastMessageTime"`
	PageReq
}

type DeleteMessage struct {
	Ids []int32 `json:"ids"`
}

type MessageHistory struct {
	Id       int32  `json:"id,omitempty"`
	Content  string `json:"content"`
	Time     string `json:"time"`
	IsSelf   bool   `json:"isSelf"`
	ShowTime bool   `json:"showTime"`
}
