/*
 * @Author: wuzhikun zhikun.wu@firstack.com
 * @Date: 2023-04-18 10:23:10
 * @LastEditors: wuzhikun zhikun.wu@firstack.com
 * @LastEditTime: 2023-04-23 09:01:08
 * @Description:
 * Copyright (c) 2023 by Firstack, All Rights Reserved.
 */
package zsys

import (
	"os"
	"path/filepath"
	"strings"
)

func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return os.IsExist(err)
}

// func FileOpen(filename string) {
// 	//TODO:
// }

func RootPath() string {
	return filepath.Dir(os.Args[0])
}

func RootFileExSuffix() string {
	return strings.TrimSuffix(os.Args[0], filepath.Ext(os.Args[0]))
}

// func Absolute(path string) string {

// }
