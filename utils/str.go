package main

import (
	"fmt"
	"regexp"
)

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
