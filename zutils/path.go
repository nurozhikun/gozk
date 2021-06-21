package zutils

import (
	"os"
	"path/filepath"
	"strings"
)

func RootPath() string {
	return filepath.Dir(os.Args[0])
}

func RootFileExSuffix() string {
	return strings.TrimSuffix(os.Args[0], filepath.Ext(os.Args[0]))
}
