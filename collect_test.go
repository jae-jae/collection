package collect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	var tests []interface{}

	tests = append(tests, []struct {
		input []int
		want  int
	}{
		{[]int{11, 22, 33}, 3},
		{[]int{22, 33, 44, 55}, 4},
		{[]int{}, 0},
	})

	tests = append(tests, []struct {
		input []string
		want  int
	}{
		{[]string{"c++", "go"}, 2},
		{[]string{"jaeger"}, 1},
		{[]string{"a", "b", "c", "d", "e"}, 5},
		{[]string{}, 0},
	})

	tests = append(tests, []struct {
		input map[string]string
		want  int
	}{
		{map[string]string{"xx": "dd", "bb": "ff"}, 2},
		{map[string]string{}, 0},
	})

	for _, g := range tests {
		gv := reflect.ValueOf(g)
		l := gv.Len()
		for i := 0; i < l; i++ {
			test := gv.Index(i)
			input := test.FieldByName("input")
			want := test.FieldByName("want").Int()
			if count := New(input).Count(); count != int(want) {
				t.Errorf("Count(%v) = %v,want %v", input, count, input)
			}
		}
	}
}

func TestSum(t *testing.T) {
	var tests []interface{}
	tests = append(tests, []struct {
		input []int
		keys  string
		cb    func(sum interface{}) bool
	}{
		{[]int{11, 22, 33}, "", func(sum interface{}) bool {
			return sum.(int64) == 66
		}},
		{[]int{22, 33, 44, 55}, "", func(sum interface{}) bool {
			return sum.(int64) == 154
		}},
		{[]int{}, "", func(sum interface{}) bool {
			return sum.(int64) == 0
		}},
	})

	for _, g := range tests {
		gv := reflect.ValueOf(g)
		l := gv.Len()
		for i := 0; i < l; i++ {
			test := gv.Index(i)
			input := test.FieldByName("input")
			keys := test.FieldByName("keys").String()
			cb := test.Interface().(struct{
				input []int
				keys  string
				cb    func(sum interface{}) bool
			})
			/*if sum := New(input).Sum(keys); cb(sum) {
				t.Errorf("sum(%v) = %v", input, sum)
			}*/
			//cb.Call([]reflect.Value{reflect.ValueOf(12)})

			//rr := cb.Call([]reflect.Value{reflect.ValueOf(int(11))})

			//cb.Call([]reflect.Value{})

			fmt.Println(input, keys,cb)
		}
	}

}
