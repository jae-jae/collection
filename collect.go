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
	//统一数据格式
	vMap := make(map[reflect.Value]reflect.Value)
	//判断输入值是否是reflect.Value，主要是test会输入这种类型的值
	v, err := value.(reflect.Value)
	if !err {
		v = reflect.ValueOf(value)
	}
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
		//@TODO:处理错误类型
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
