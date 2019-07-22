package tdt

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
)

// 动态调用被测试函数，将测试流程代码抽象出来
func RunTest(t *testing.T, fn interface{}, cases []TestCase) {
	convertCaseIn(fn, cases)
	for _, c := range cases {
		result, err := dynamicCall(fn, c.In...)
		if err != nil {
			t.Fatalf("Check case %s error: %s\n", c.Name, err)
		}
		resp := result[0].Interface()
		checkResult, err := checkResp(resp, c.CheckItems)
		if !checkResult || err != nil {
			t.Fatalf("Check case %s error: %s, resp: %+v\n", c.Name, err, resp)
		}
	}
}

// 将测试用例中的输入转换为函数需要的类型
func convertCaseIn(fn interface{}, cases []TestCase) error {
	var paramType reflect.Type
	param := reflect.TypeOf(fn).In(0)
	if param.Kind() == reflect.Ptr {
		paramType = param.Elem()
	} else {
		paramType = param
	}
	for index, item := range cases {
		parsedData := reflect.New(paramType)
		bytes, _ := json.Marshal(item.In[0])
		err := json.Unmarshal(bytes, parsedData.Interface())
		if err != nil {
			return err
		}
		if param.Kind() == reflect.Ptr {
			cases[index].In[0] = parsedData.Interface()
		} else {
			cases[index].In[0] = parsedData.Elem().Interface()
		}
	}
	return nil
}

// 动态调用函数
func dynamicCall(fn interface{}, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(fn)
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
