package zutils

import (
	"fmt"
	"runtime"
)

type Error struct {
	Code      int //错误代码
	FileName  string
	Lines     int
	ErrString string
}

func NewError(code int, s string) error {
	err := &Error{Code: code, ErrString: s}
	_, err.FileName, err.Lines, _ = runtime.Caller(1)
	// err.FileName = filepath.Base(err.FileName)
	return err
}

func ErrorCode(err error) int {
	e, ok := err.(*Error)
	if ok {
		return e.Code
	}
	return 0
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s [%d] %s", e.FileName, e.Lines, e.ErrString)
}
