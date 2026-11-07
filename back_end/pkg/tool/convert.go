package tool

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

// base64 -string
func Base64ToStr(base64Str string) string {
	// 解码 Base64 字符串
	decodedBytes, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		fmt.Printf("Base64ToStr Err: 解码失败: %v\n", err)
		return ""
	}

	// 将解码后的字节切片转换为字符串
	return string(decodedBytes)
}

// 16进制转字符串
func HexToStr(hex string) {

}

// StringToIntSlice 把字符串转为整数切片
func StringToIntSlice(s string) ([]int, error) {
	// 去掉首尾的方括号
	s = strings.Trim(s, "[]")
	// 用逗号分隔字符串
	parts := strings.Split(s, ",")
	result := make([]int, 0, len(parts))
	for _, part := range parts {
		// 去掉可能存在的空白字符
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		// 把字符串转为整数
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}
	return result, nil
}

// IntSliceToString 把整数切片转为字符串
func IntSliceToString(nums []int) string {
	var parts []string
	for _, num := range nums {
		parts = append(parts, strconv.Itoa(num))
	}
	return "[" + strings.Join(parts, ",") + "]"
}

// AppendAndDeduplicate 追加数字并去重
func AppendAndDeduplicate(s *string, newNums ...int) (string, error) {
	nums, err := StringToIntSlice(*s)
	if err != nil {
		return "", err
	}
	numMap := make(map[int]bool)
	for _, num := range nums {
		numMap[num] = true
	}
	for _, newNum := range newNums {
		if !numMap[newNum] {
			nums = append(nums, newNum)
			numMap[newNum] = true
		}
	}
	return IntSliceToString(nums), nil
}

func RemoveDuplicates(s *string) (string, error) {
	nums, err := StringToIntSlice(*s)
	if err != nil {
		return "", err
	}
	numMap := make(map[int]bool)
	var uniqueNums []int
	for _, num := range nums {
		if !numMap[num] {
			uniqueNums = append(uniqueNums, num)
			numMap[num] = true
		}
	}
	return IntSliceToString(uniqueNums), nil
}

func StringToStringSlice(s string) []string {
	// 去除两端的方括号
	s = strings.Trim(s, "[]")
	// 以逗号为分隔符拆分字符串
	parts := strings.Split(s, ",")
	result := make([]string, 0, len(parts))

	for _, part := range parts {
		// 去除两端的引号
		part = strings.Trim(part, "\"")
		if part != "" {
			result = append(result, part)
		}
	}

	return result
}

func StringSliceToString(slice []string) string {
	var parts []string
	for _, str := range slice {
		parts = append(parts, fmt.Sprintf("\"%s\"", str))
	}
	result := "[" + strings.Join(parts, ",") + "]"
	return result
}
