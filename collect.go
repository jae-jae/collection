package collect

import (
	"reflect"
	"strings"
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
		for _, key := range v.MapKeys() {
			vMap[key] = v.MapIndex(key)
		}
	default:

	}
	return &Collecter{vMap, v.Kind()}
}

func getValueByKeys(item reflect.Value, keys string) reflect.Value {
	keyArr := strings.Split(keys, ".")
	for _, key := range keyArr {
		item = item.MapIndex(reflect.ValueOf(key))
	}
	return item
}
