package tool

import (
	"fmt"
	"time"
)

// 定义标准时间
// yyyy-MM-dd HH:mm:ss
type StdTime time.Time

func (t *StdTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func NowStr() string {
	// 获取当前时间
	now := time.Now()
	// 格式化时间
	return now.Format("2006-01-02 15:04:05")
}

func TimeStr(t time.Time) string {
	layout := "2006-01-02 15:04:05"
	return t.Format(layout)
}

func StrTime(t string) time.Time {
	layout := "2006-01-02 15:04:05"
	parse, err := time.Parse(layout, t)
	if err != nil {
		return time.Now()
	}
	return parse
}

// FormatTimeInterval 根据两个时间的差值返回格式化的时间间隔字符串
func FormatTimeInterval(startTime, endTime time.Time) string {
	duration := endTime.Sub(startTime)

	if duration < time.Minute {
		return "10秒钟前"
	} else if duration < time.Hour {
		minutes := int(duration.Minutes())
		return fmt.Sprintf("%d分钟前", minutes)
	} else if duration < 24*time.Hour {
		hours := int(duration.Hours())
		return fmt.Sprintf("%d小时前", hours)
	} else if duration < 30*24*time.Hour {
		days := int(duration.Hours() / 24)
		return fmt.Sprintf("%d天前", days)
	} else {
		months := int(duration.Hours() / (24 * 30))
		return fmt.Sprintf("%d个月前", months)
	}
}
