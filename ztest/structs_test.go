package ztest

import (
	"fmt"
	"reflect"
	"testing"
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
