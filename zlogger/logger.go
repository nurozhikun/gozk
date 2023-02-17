package zlogger

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

type Config struct {
}

var inforLogger = log.Default()

func init() {
	inforLogger.SetPrefix("{INFO} ")
	inforLogger.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmsgprefix)
}

func Error(e ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Println(filepath.Base(file), "[ERROR]", line, e)
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
