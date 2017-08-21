package collect

import (
	"reflect"
)

func (c *Collecter) Count() int {
	return len(c.value)
}

func (c *Collecter) Map(callback func(item reflect.Value, key reflect.Value) interface{}) *Collecter {
	var rtArr []interface{}
	for key, item := range c.value {
		rtArr = append(rtArr, callback(item, key))
	}
	return New(rtArr)
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
