package zlogger

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

func Error(e ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Println(filepath.Base(file), line, "{error}", e)
}

func Info(s ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), filepath.Base(file), line, "{info}", s)
}

func Println(s ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Println(filepath.Base(file), line, s)
}
