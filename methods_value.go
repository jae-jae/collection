package collect

import (
	"reflect"
)

func (c *Collecter) Count() int {
	return len(c.value)
}

func (c *Collecter) All() interface{} {
	switch c.kind {
	case reflect.Array, reflect.Slice:
		/*var tArr []interface{}
		for k,v := range c.value{
			tArr[k.Int()] = v.Interface()
		}*/
	case reflect.Map:

	}
	return nil
}

func (c *Collecter) avg() {

}

func (c *Collecter) Reduce(callback func(carry interface{}, item interface{}, key interface{}) interface{}, init interface{}) interface{} {
	var carryData = init
	for key, item := range c.value {
		carryData = callback(carryData, item.Interface(), key.Interface())
	}
	return carryData
}

func (c *Collecter) Sum(keys string) interface{} {
	var sum interface{}
	if len(keys) == 0 {
		sum = c.Reduce(func(carry interface{}, item interface{}, key interface{}) interface{} {
			return sumKind(carry, item)
		}, 0)
	} else {
		sum = c.Reduce(func(carry interface{}, item interface{}, key interface{}) interface{} {
			value := getValueByKeys(reflect.ValueOf(item), keys)
			return sumKind(carry, value.Interface())
		}, 0)
	}
	return sum
}

func sumKind(carry interface{}, item interface{}) interface{} {
	var sum interface{}
	itemValue := reflect.ValueOf(item)
	carryValue := reflect.ValueOf(carry)
	switch itemValue.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		sum = carryValue.Int() + itemValue.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if carryValue.Kind() == reflect.Int {
			sum = uint64(carryValue.Int()) + itemValue.Uint()
		} else {
			sum = carryValue.Uint() + itemValue.Uint()
		}
	case reflect.Float32, reflect.Float64:
		if carryValue.Kind() == reflect.Int {
			sum = float64(carryValue.Int()) + itemValue.Float()
		} else {
			sum = carryValue.Float() + itemValue.Float()
		}
	}
	rt := reflect.ValueOf(&item).Elem()
	rt.Set(reflect.ValueOf(sum))
	return rt.Interface()
}
