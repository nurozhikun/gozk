package ztest

import (
	"fmt"
	"reflect"
	"testing"

	"gitee.com/sienectagv/gozk/zdev"
	"gitee.com/sienectagv/gozk/zreflect"
)

type Person struct {
	Name    string `zdev:"name"`
	YearOld int    `zdev:"year_old"`
}

type Worker struct {
	Person
	Factory string `zdev:"factory"`
}

func TestStructTag(t *testing.T) {
	w := Worker{Person: Person{Name: "John", YearOld: 24}, Factory: "sd"}
	fmt.Println(w)
	tp := reflect.TypeOf(&w)
	if tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}
	// fmt.Println(tp.Kind(), tp.Name())
	fmt.Println(tp.FieldByName("Name"))
	// fmt.Println(tp.FieldByName("YearOld"))
	fmt.Println(tp.FieldByName("Factory"))
	vp := reflect.ValueOf(w)
	fmt.Println(vp.FieldByName("Name"))
	sf, ok := tp.FieldByName("YearOld")
	if ok {
		fmt.Println(vp.FieldByIndex(sf.Index))
	}
	// fmt.Println(vp.Elem().NumField())
}

//to test -v .\structs_test.go -run TestZDevNodes
func TestZDevNodes(t *testing.T) {
	node := &zdev.VirtualNode{
		ID:     "abc",
		Addr:   ":8000",
		Stream: zdev.NewStreamTcpListener(),
		Custom: zdev.NewCustomTcpListener(func() zdev.ICustom { return nil }),
	}
	mp := zreflect.StructFieldsByTag(node, "zdev")
	fmt.Println(mp)
	// tp := reflect.TypeOf(node)
	// if tp.Kind() == reflect.Ptr {
	// 	tp = tp.Elem()
	// }
	// for i := 0; i < tp.NumField(); i++ {
	// 	fmt.Println(tp.Field(i).Type.Kind())
	// }
}
