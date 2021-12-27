package zreflect

import (
	"reflect"
	"sort"

	"gitee.com/sienectagv/gozk/zlogger"
	"gitee.com/sienectagv/gozk/zmap"
	"gitee.com/sienectagv/gozk/zsort"
	"gitee.com/sienectagv/gozk/zstrings"
)

// type StructFields struct {
// 	TopVal reflect.Value
// 	values zmap.Map
// }
type (
	Map = zmap.Map
)

func StructFieldsByTag(i interface{}, tagName string, tagVals ...string) Map {
	zr := &zreflect{in: i,
		// topType: reflect.TypeOf(i),
		// topVal:  reflect.ValueOf(i),
		m:       zmap.NewMap(),
		tagName: tagName,
		tagVals: tagVals}
	zr.tagVals.Sort()
	zr.processOneInterface(i)
	return zr.m
}

type zreflect struct {
	in      interface{}
	topType reflect.Type
	topVal  reflect.Value
	m       Map
	tagName string
	tagVals sort.StringSlice //[]string
}

func (z *zreflect) processOneInterface(i interface{}) {
	it := reflect.TypeOf(i)
	iv := reflect.ValueOf(i)
	if it.Kind() == reflect.Ptr {
		it = it.Elem()
		iv = iv.Elem()
	}
	// fmt.Println(it)
	zlogger.Info(it, it.Kind())
	if it.Kind() != reflect.Struct {
		return
	}
	z.processOneStruct(it, iv)
}

func (z *zreflect) processOneStruct(st reflect.Type, sv reflect.Value) {
	zlogger.Info(st.NumField())
	for i := 0; i < st.NumField(); i++ {
		stField := st.Field(i)
		svField := sv.Field(i)
	ONE_FIELD:
		curType := stField.Type
		zlogger.Info(i, curType, curType.Kind())
		switch curType.Kind() {
		case reflect.Ptr:
			svField = sv.Elem()
			curType = svField.Type()
			goto ONE_FIELD
		case reflect.Interface:
			if !z.tryTagField(stField, svField) {
				z.processOneInterface(sv.Interface())
			}
		case reflect.Struct:
			z.processOneStruct(curType, svField)
		case reflect.Invalid:
		case reflect.UnsafePointer:
		default:
			z.tryAddField(stField, svField)
		}
	}
}

func (z *zreflect) tryTagField(stF reflect.StructField, v reflect.Value) bool {
	name, ok := stF.Tag.Lookup(z.tagName)
	if !ok {
		return false
	}
	if z.tagVals.Len() > 0 && zsort.SSIndex(z.tagVals, name) == -1 {
		return false
	}
	z.m.Insert(name, v.Interface())
	return true
}

func (z *zreflect) tryAddField(stF reflect.StructField, v reflect.Value) {
	name, ok := stF.Tag.Lookup(z.tagName)
	if !ok {
		name = zstrings.HumpToUnderline(stF.Name)
	}
	if z.tagVals.Len() > 0 && zsort.SSIndex(z.tagVals, name) == -1 {
		return
	}
	z.m.Insert(name, v.Interface())
}

// func processOneStruct()

// func structFieldsByTag(i)
