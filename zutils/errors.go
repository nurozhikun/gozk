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

const (
	ErrCodeHasExist = 1 + iota
	ErrCodeNotFound
	ErrCodeUserOrPassMiss
	ErrCodeCmdUnexist
	ErrCodeNullParam
)

var (
	ErrHasExist       = NewError(ErrCodeHasExist, "the value has existed")
	ErrNotFound       = NewError(ErrCodeNotFound, "the value has not found")
	ErrUserOrPassMiss = NewError(ErrCodeUserOrPassMiss, "the user isn't found or wrong password")
	ErrCmdUnexist     = NewError(ErrCodeCmdUnexist, "the command is unexist")
	ErrNullParam      = NewError(ErrCodeNullParam, "there are some null parameters")
)

func NewError(code int, s string) error {
	err := &Error{Code: code, ErrString: s}
	_, err.FileName, err.Lines, _ = runtime.Caller(1)
	// err.FileName = filepath.Base(err.FileName)
	return err
}

func ErrorCode(err error) int {
	if nil == err {
		return 0
	}
	e, ok := err.(*Error)
	if ok {
		return e.Code
	}
	return -1
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s [%d] %s", e.FileName, e.Lines, e.ErrString)
}
