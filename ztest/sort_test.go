package ztest

import (
	"fmt"
	"sort"
	"testing"
)

func TestStringSort(t *testing.T) {
	ss := sort.StringSlice{"abc", "bcd", "ddy", "kill"}
	fmt.Println(ss.Search("abc"))
	fmt.Println(ss.Search("kill"))
	// fmt.Println(ss.Search("dddd"))
	fmt.Println("end", 'A', 'a'-'A')
	t.Log("is ok")
}
