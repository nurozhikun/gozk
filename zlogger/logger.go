package zlogger

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

type config struct {
}

var defCfg = &config{}

func Error(e ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Println(filepath.Base(file), "[ERROR]", line, e)
}

func Info(s ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), filepath.Base(file), "[INFO]", line, s)
}

func Println(s ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Println(filepath.Base(file), line, s)
}
