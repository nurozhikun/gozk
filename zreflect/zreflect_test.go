package zreflect

import (
	"fmt"
	"testing"
	// "github.com/nurozhikun/gozk/zreflect"
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
	m := StructFieldsByTag(w, "zdev", "name", "year_old", "factory")
	fmt.Println(w, m)
}
