package util

import (
	"github.com/google/uuid"
	"strings"
)

// GenerateStr 生成32位随机字符串
func GenerateStr() string {
	str, _ := uuid.NewUUID()
	return strings.ReplaceAll(str.String(), "-", "")
}

func Generate32Str() string {
	return GenerateStr()
}

func Generate24Str() string {
	str := Generate32Str()
	return str[0:24]
}

func Generate16Str() string {
	str := Generate32Str()
	return str[0:16]
}
