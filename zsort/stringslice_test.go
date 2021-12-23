package zsort

import (
	"fmt"
	"testing"

	"gitee.com/sienectagv/gozk/zsort"
)

func TestStringSort(t *testing.T) {
	ss := zsort.StringSlice{"abc", "bcd", "ddy", "kill"}
	fmt.Println(ss.Index("abc"))
	fmt.Println(ss.Index("kill"))
	fmt.Println(ss.Index("dddd"))
}
