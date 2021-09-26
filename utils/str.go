package utils

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//字符串数组转整型数组
func StringArray2IntArray(list []string) []int {
	stringArray := DeleteBlankString(list)
	idx := 0
	array := make([]int, len(stringArray))
	for _, v := range list {
		num, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		array[idx] = num
		idx += 1
	}
	if idx == len(stringArray) {
		return array
	}
	return array[:idx]
}

//字符串数组转整型数组
func IntArray2StringArray(list []int) []string {
	array := make([]string, len(list))
	for idx, v := range list {
		array[idx] = strconv.Itoa(v)
	}
	return array
}

//数组里删除空白字符
func DeleteBlankString(list []string) []string {
	f := func(v string) bool {
		return IsBlank(v)
	}
	return DeleteStringByCondition(list, f)
}

func DeleteStringByCondition(list []string, f func(v string) bool) []string {
	idx := 0
	array := make([]string, len(list))
	for _, v := range list {
		if f(v) {
			continue
		}
		array[idx] = v
		idx += 1
	}
	if idx == len(list) {
		return list
	}
	return array[:idx]
}

//获取md5
func GetMd5(s string) string {
	data := []byte(s)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

//正则比较
func IsMatchString(s string, expr string) bool {
	r, _ := regexp.Compile(expr)
	return r.MatchString(s)
}

//判断字符串是不是空白字符
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

//是否合法域名
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
