package collect

import (
	"reflect"
)

type Collecter struct {
	value map[reflect.Value]reflect.Value
	kind  reflect.Kind
}

func New(value interface{}) *Collecter {
	vMap := make(map[reflect.Value]reflect.Value)
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		length := v.Len()
		for i := 0; i < length; i++ {
			vMap[reflect.ValueOf(i)] = v.Index(i)
		}
	case reflect.Map:

	default:

	}
	return &Collecter{vMap, v.Kind()}
}
