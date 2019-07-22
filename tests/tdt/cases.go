package tdt

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	NORMAL      = iota + 1 // 正常流程
	PARAM_CHECK            // 参数校验
)

type TestCase struct {
	Name       string
	Type       int
	In         []interface{}
	CheckItems []CheckItem
}

type CheckItem struct {
	Field    string
	Expected interface{}
}

func loadCases(filepath string) (cases []TestCase, err error) {
	file, _ := os.Open(filepath)

	data := []TestCase{}

	decoder := json.NewDecoder(file)
	decoder.UseNumber()

	err = decoder.Decode(&data)
	if err != nil {
		fmt.Printf("json unmarshal error: %s", err)
		return nil, err
	}

	cases = data

	return
}
