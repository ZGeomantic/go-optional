package optional

import (
	"reflect"
)

func Inspect(ptrStruct interface{}) (t reflect.Type) {
	t = reflect.TypeOf(ptrStruct)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return
}

func Optional(i interface{}) interface{} {
	t := Inspect(i)
	return reflect.Zero(t).Interface()
}
