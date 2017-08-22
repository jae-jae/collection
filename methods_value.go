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

func (c *Collecter) SumInt(keys string) int {
	var sum interface{}
	if len(keys) == 0 {
		sum = c.Reduce(func(carry interface{}, item interface{}, key interface{}) interface{} {
			return item.(int) + carry.(int)
		}, 0)
	} else {
		sum = c.Reduce(func(carry interface{}, item interface{}, key interface{}) interface{} {
			value := getValueByKeys(reflect.ValueOf(item), keys)
			return value.Interface().(int) + carry.(int)
		}, 0)
	}
	return sum.(int)
}
