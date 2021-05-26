package zsys

import (
	"os"
)

func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return os.IsExist(err)
}
