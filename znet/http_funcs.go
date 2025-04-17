/*
 * @Author: wuzhikun zhikun.wu@firstack.com
 * @Date: 2023-04-19 11:02:01
 * @LastEditors: wuzhikun zhikun.wu@firstack.com
 * @LastEditTime: 2023-08-02 16:05:12
 * @Description:
 * Copyright (c) 2023 by Firstack, All Rights Reserved.
 */
package znet

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/nurozhikun/gozk/zlogger"
)

func HttpGetFile(urlQuery string, fn func(body io.Reader)) error {
	// zlogger.Info(urlQuery)
	u, err := url.Parse(urlQuery)
	if nil != err {
		return err
	}
	// zlogger.Info(u.String())
	res, err := http.Get(u.String())
	if nil != err {
		return err
	}
	if nil != res.Body {
		fn(res.Body)
		res.Body.Close()
	}
	return nil
}

func HttpDownload(urlQuery, cachePath string) error {
	var err error
	err = HttpGetFile(urlQuery, func(body io.Reader) {
		zlogger.Info("Dowload:", urlQuery)
		var file *os.File
		fileName := filepath.Join(cachePath, filepath.Base(urlQuery))
		zlogger.Info(fileName)
		file, err = os.Create(fileName)
		if nil != err {
			return
		}
		defer file.Close()
		_, err = io.Copy(file, body)
	})
	return err
}

func HttpUploadFilePost(urlQuery, fileName string) (*http.Response, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	formFile, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return nil, err
	}
	io.Copy(formFile, file)
	writer.Close()
	return http.Post(urlQuery, writer.FormDataContentType(), body)
}
