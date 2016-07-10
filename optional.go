package optional

import (
	"reflect"
)

// 获得非指针类型的对象，不能通过value.type()获取，因为如果是控空指针的话，对它执行的Value.Elem()之后，其Type已经变为invalid了
// 为了执行效率，在这里判断是否是zero
func inspectNonptrType(ptrStruct interface{}) (t reflect.Type, isnil bool) {
	t = reflect.TypeOf(ptrStruct)
	isnil = ptrStruct == reflect.Zero(t).Interface()
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return
}

func Deref(i interface{}) interface{} {
	t, isnil := inspectNonptrType(i)

	if isnil {
		return reflect.Zero(t).Interface()
	} else {
		return reflect.ValueOf(i).Elem().Interface()
	}
}

func DerefMust(i interface{}, defaultValue interface{}) interface{} {

	if IsZero(i) {
		return defaultValue
	} else {
		return reflect.ValueOf(i).Elem().Interface()
	}
}

func IsZero(i interface{}) bool {
	t := reflect.TypeOf(i)
	return i == reflect.Zero(t).Interface()
}
