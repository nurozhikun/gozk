/*
 * @Author: wuzhikun zhikun.wu@firstack.com
 * @Date: 2023-04-21 16:36:55
 * @LastEditors: wuzhikun zhikun.wu@firstack.com
 * @LastEditTime: 2023-04-23 13:34:48
 * @Description:
 * Copyright (c) 2023 by Firstack, All Rights Reserved.
 */
package zarchive

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/nurozhikun/gozk/zlogger"
)

func UnzipToFolder(zipFileName, outFolder string) error {
	rzip, err := zip.OpenReader(zipFileName)
	if err != nil {
		zlogger.Error(err)
		return err
	}
	defer rzip.Close()
	//unzip
	var errRet error
	for _, f := range rzip.File {
		if f.FileInfo().IsDir() {
			dirPath := path.Join(outFolder, f.Name)
			fmt.Printf("try create path: %s\n", dirPath)
			os.MkdirAll(dirPath, os.ModeType)
		} else {
			fmt.Printf("try unzip file %s\n", f.Name)
			if rc, err := f.Open(); err == nil {
				dstFileName := path.Join(outFolder, f.Name)
				dirPath := filepath.Dir(dstFileName)
				os.Mkdir(dirPath, os.ModeType)
				// zlogger.Println(dstFileName)
				if df, err := os.Create(dstFileName); err == nil || os.IsNotExist(err) {
					io.Copy(df, rc)
					df.Close()
				} else {
					zlogger.Error(err)
					errRet = err
				}
				rc.Close()
			} else {
				zlogger.Error(err)
				errRet = err
			}
			if errRet != err {
				break
			}
		}
	}
	return errRet
}
