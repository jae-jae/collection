package collect

import (
	"reflect"
)

func (c *data) Count() int {
	var len int
	if reflect.TypeOf(c.value).String() == "reflect.Value"{
		len = (c.value).(reflect.Value).Len()
	}else{
		len = reflect.ValueOf(c.value).Len()
	}
	return len
}
