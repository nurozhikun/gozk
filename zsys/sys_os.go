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

func RootPath() string {
	return filepath.Dir(os.Args[0])
}

func RootFileExSuffix() string {
	return strings.TrimSuffix(os.Args[0], filepath.Ext(os.Args[0]))
}

// func Absolute(path string) string {

// }
