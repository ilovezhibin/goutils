package utils

import "strings"

const COMMA string = ","

//数组转逗号分隔，去除空白字符
func ToCommaString(list []string) string {
	return strings.Join(DeleteBlankString(list), COMMA)
}

//数组转逗号分隔，去除空白字符
func ToCommaStringByInt(list []int) string {
	return strings.Join(IntArray2StringArray(list), COMMA)
}

//从逗号分隔的字符串转成数组，去除空白字符
func ToListFromCommaString(s string) []string {
	return DeleteBlankString(strings.Split(s, COMMA))
}

//从逗号分隔的字符串转成数组，去除空白字符
func ToIntListFromCommaString(s string) []int {
	array := ToListFromCommaString(s)
	return StringArray2IntArray(array)
}
