package enum

// 评论类型

const (
	CMT_LOVE = iota + 1
	CMT_TOPIC
	CMT
)

var CMT_TYPE_LIST = []int{CMT_LOVE, CMT_TOPIC, CMT}
