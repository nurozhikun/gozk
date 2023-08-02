/*
 * @Author: wuzhikun zhikun.wu@firstack.com
 * @Date: 2023-08-02 14:27:09
 * @LastEditors: wuzhikun zhikun.wu@firstack.com
 * @LastEditTime: 2023-08-02 16:05:39
 * @Description:
 * Copyright (c) 2023 by Firstack, All Rights Reserved.
 */
package zfile

import (
	"bufio"
	"os"
)

func IterateLine(fileName string, fnProc func(line string)) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(file)
	var line []byte
	for nil == err {
		line, _, err = buf.ReadLine()
		fnProc(string(line))
	}
	return nil
}
