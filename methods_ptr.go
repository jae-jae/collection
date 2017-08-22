package collect

func (c *Collecter) Map(callback func(item interface{}, key interface{}) interface{}) *Collecter {
	var rtArr []interface{}
	for key, item := range c.value {
		rtArr = append(rtArr, callback(item.Interface(), key.Interface()))
	}
	return New(rtArr)
}

func (c *Collecter) Filter(callback func(item interface{}, key interface{}) bool) *Collecter {
	var rtMap = make(map[interface{}]interface{})
	for key, item := range c.value {
		if callback(item.Interface(), key.Interface()) {
			rtMap[key.Interface()] = item.Interface()
		}
	}
	return New(rtMap)
}
