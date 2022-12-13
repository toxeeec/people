package postgres

import (
	"reflect"
)

func Flatten(elems ...any) []any {
	ret := make([]any, 0, len(elems))
	flattenInner(&ret, reflect.ValueOf(elems))
	return ret
}

func flattenInner(ret *[]any, v reflect.Value) {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			flattenInner(ret, v.Index(i))
		}
		return
	}
	*ret = append(*ret, v.Interface())
}
