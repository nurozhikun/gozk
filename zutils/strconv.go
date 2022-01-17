package zutils

import "strconv"

func ToInt32(s string) int32 {
	n, _ := strconv.Atoi(s)
	return int32(n)
}

func Itoa(v interface{}) string {
	if i, ok := InterfaceToInt(v); ok {
		return strconv.FormatInt(i, 10)
	}
	if u, ok := InterfaceToUint(v); ok {
		return strconv.FormatUint(u, 10)
	}
	return ""
}
