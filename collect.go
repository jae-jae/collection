package collect

type data  struct {
	value interface{}
}

func New(value interface{})  *data{
	return &data{value}
}

