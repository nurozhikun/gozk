package zsort

import (
	"fmt"
	"testing"
	// "github.com/nurozhikun/gozk/zsort"
)

func TestStringSort(t *testing.T) {
	ss := StringSlice{"abc", "bcd", "ddy", "kill"}
	fmt.Println(ss.Index("abc"))
	fmt.Println(ss.Index("kill"))
	fmt.Println(ss.Index("dddd"))
}
