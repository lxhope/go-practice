package main

import "fmt"
import "reflect"
import "errors"

func foo() {
	fmt.Println("we are running foo")
}

func bar(a, b, c int) {
	fmt.Println("we are running bar", a, b, c)
}

func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

func main() {
	// nota bene: for perfect score: use reflection to build this map
	funcs := map[string]interface{}{
		"foo": foo,
		"bar": bar,
	}

	Call(funcs, "foo")
	Call(funcs, "bar", 1, 2, 3)
}
