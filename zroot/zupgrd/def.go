/*
 * @Author: wuzhikun zhikun.wu@firstack.com
 * @Date: 2023-05-05 18:36:11
 * @LastEditors: wuzhikun zhikun.wu@firstack.com
 * @LastEditTime: 2023-05-05 18:50:54
 * @Description:
 * Copyright (c) 2023 by Firstack, All Rights Reserved.
 */
package zupgrd

import (
	"strconv"
	"strings"
)

func VersionCmp(a, b string) int {
	sa := strings.Split(a, ".")
	sb := strings.Split(b, ".")
	if len(sa) > len(sb) {
		return 1
	} else if len(sa) < len(sb) {
		return -1
	}
	for i, v := range sa {
		ia, _ := strconv.Atoi(v)
		ib, _ := strconv.Atoi(sb[i])
		if ia > ib {
			return 1
		} else if ia < ib {
			return -1
		}
	}
	return strings.Compare(a, b)
}
