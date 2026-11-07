package enum

const (
	ALL  = iota
	LOVE // 表白
	FIND // 寻人寻物
	VOTE // 找搭子
	PICTURE
)

var POSTING_TYPE = []int{ALL, LOVE, FIND, VOTE, PICTURE}

// 货物类型
type GOODS_TYPE = string

const (
	GOODS_ALL     = "all"
	GOODS_BOOKS   = "books"
	GOODS_DIGITAL = "digital"
	GOODS_CLOTHES = "clothes"
	GOODS_DAILY   = "daily"
	GOODS_OTHERS  = "others"
)

var GOODS_TYPE_LIST = []GOODS_TYPE{GOODS_ALL, GOODS_BOOKS, GOODS_DIGITAL, GOODS_CLOTHES, GOODS_DAILY, GOODS_OTHERS}

var GOODS_TYPE_MAP = map[GOODS_TYPE]string{
	GOODS_ALL:     "全部",
	GOODS_BOOKS:   "教材书籍",
	GOODS_DIGITAL: "数码产品",
	GOODS_CLOTHES: "服饰装备",
	GOODS_DAILY:   "生活用品",
	GOODS_OTHERS:  "其他",
}

const (
	GOODS_SOLD    = "sold"
	GOODS_SELLING = "selling"
)
