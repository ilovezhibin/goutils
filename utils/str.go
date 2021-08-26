package utils

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"strings"
)

func GetMd5(s string) string {
	data := []byte(s)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

func IsIpV6(s string) bool {
	return IsMatchString(s, "^((?:[0-9A-Fa-f]{1,4}(?::[0-9A-Fa-f]{1,4})*)?)::((?:[0-9A-Fa-f]{1,4}(?::[0-9A-Fa-f]{1,4})*)?)$") || IsMatchString(s, "^(?:[0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}$")
}

func IsIpV4(s string) bool {
	r, err := regexp.Compile("^(([01]?\\d\\d?|2[0-4]\\d|25[0-5])\\.){3}([01]?\\d\\d?|2[0-4]\\d|25[0-5])$")
	if err != nil {
		fmt.Println("regext error", err)
	}
	return r.MatchString(s)
}

func IsMatchString(s string, expr string) bool {
	r, _ := regexp.Compile(expr)
	return r.MatchString(s)
}

func IsBlank(s string) bool {
	s = strings.TrimSpace(s)
	return s == ""
}

//判断给定的值，在不在后续给的数据里
func ValueInList(value interface{}, list ...interface{}) bool {
	for _, item := range list {
		if value == item {
			return true
		}
	}
	return false
}

func IsLegalDomain(domain string) bool {
	r, err := regexp.Compile("^[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+$")
	if err != nil {
		fmt.Println("regext error", err)
	}
	return r.MatchString(domain)
}

//截取指定长度字符串
func GetLimitString(s string, limit int) string {
	if len(s) > limit {
		return s[:limit]
	}
	return s
}

//把驼峰转为下划线
func HumpToUnderline(s string) string {
	array := []byte(s)
	newArray := make([]byte, 0)
	for _, item := range array {
		if item >= 65 && item <= 90 {
			newArray = append(newArray, '_', item+32)
		} else {
			newArray = append(newArray, item)
		}

	}

	return string(newArray)
}
