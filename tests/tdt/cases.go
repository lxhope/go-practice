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
	Name       string        `json:"name"`
	Type       int           `json:"type"`
	In         []interface{} `json:"in"`
	CheckItems []CheckItem   `json:"check_items"`
}

type CheckItem struct {
	Field    string      `json:"field"`
	Expected interface{} `json:"expected"`
}

func LoadCases(filepath string) (cases []TestCase, err error) {
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
