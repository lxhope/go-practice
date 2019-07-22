package main

// 函数参数动态传入

import (
	"encoding/json"
	"fmt"
	"reflect"
)

const (
	testData = `{ "msg": "hello" }`
)

type EchoData struct {
	Msg string `json:"msg"`
}

func echo(data *EchoData) {
	fmt.Println("Received:", data.Msg)
}

func main() {
	fnv := reflect.ValueOf(echo)
	argt := reflect.TypeOf(echo).In(0).Elem()
	parsedData := reflect.New(argt)

	err := json.Unmarshal([]byte(testData), parsedData.Interface())

	if err == nil {
		args := []reflect.Value{parsedData}
		fnv.Call(args)
	} else {
		fmt.Println(err)
	}
}
