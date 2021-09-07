package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

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

func Ipv4ToInt64(ip string) (int64, error) {
	ipSlices := strings.Split(ip, ".")
	var result int64 = 0
	for i, s := range ipSlices {
		num, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		t := num << (8 * (3 - i))
		temp := int64(t)
		result = result | temp
	}
	return result, nil
}
