package zmap

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := NewMap()
	m.Insert(1, 100)
	m.Insert(2, 200)
	m.Insert("abc", "abc100")
	fmt.Println(m)
	fmt.Println(m.GetInt(1))
	fmt.Println(m.GetInt(2))
	fmt.Println(m.GetInt("abc"))
}
