package tool

import (
	"github.com/google/uuid"
	"strings"
)

// UUID 生成没有破折号的 UUID
func Uuid() string {
	return uuid.NewString()
}

func Uid() string {
	return strings.ReplaceAll(Uuid(), "-", "")
}
