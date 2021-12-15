package zutils

import "strconv"

func ToInt32(s string) int32 {
	n, _ := strconv.Atoi(s)
	return int32(n)
}
