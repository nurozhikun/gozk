/*
 * @Author: wuzhikun zhikun.wu@firstack.com
 * @Date: 2023-04-18 10:23:10
 * @LastEditors: wuzhikun zhikun.wu@firstack.com
 * @LastEditTime: 2023-04-18 16:58:21
 * @Description:
 * Copyright (c) 2023 by Firstack, All Rights Reserved.
 */
package zlogger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type Config struct {
}

var (
	inforLogger = log.New(log.Default().Writer(), "[INFO] ", log.LstdFlags|log.Lshortfile|log.Lmsgprefix)
	debugLogger = log.New(log.Default().Writer(), "[DEBUG] ", log.LstdFlags|log.Lshortfile|log.Lmsgprefix)
	errorLogger = log.New(log.Default().Writer(), "[ERROR] ", log.LstdFlags|log.Lshortfile|log.Lmsgprefix)
)

// func init() {
// 	inforLogger.SetPrefix("{INFO} ")
// 	inforLogger.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmsgprefix)
// }

func InitLogPath(path string) {
	Info(path)
	err := os.MkdirAll(path, os.ModeType)
	if nil != err {
		Error(err)
	}
	fileName := filepath.Join(path, time.Now().Format("2006-01-02")+".txt")
	logfile, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if nil != err {
		Error(err)
	}
	inforLogger.SetOutput(logfile)
	debugLogger.SetOutput(logfile)
	errorLogger.SetOutput(logfile)
}

func Error(e ...interface{}) {
	// _, file, line, _ := runtime.Caller(1)
	// fmt.Println(filepath.Base(file), "[ERROR]", line, e)
	errorLogger.Output(2, fmt.Sprintln(e...))
}

func Info(s ...interface{}) {
	// _, file, line, _ := runtime.Caller(1)
	// fmt.Println(time.Now().Format("2006-01-02 15:04:05"), filepath.Base(file), "[INFO]", line, s)
	inforLogger.Output(2, fmt.Sprintln(s...))
}

func Println(s ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Println(filepath.Base(file), line, s)
}
