package tdt

import (
	"fmt"
	"reflect"

	"github.com/thoas/go-funk"
)

func checkResp(resp interface{}, items []CheckItem) (bool, error) {
	for _, item := range items {
		result, err := checkField(resp, item.Field, item.Expected)
		if err != nil {
			return false, err
		}
		if !result {
			err = fmt.Errorf("Check path %s error", item.Field)
			return result, err
		}
	}
	return true, nil
}

func checkField(data interface{}, path string, expect interface{}) (bool, error) {
	actual := funk.Get(data, path)
	if expect == nil && isNil(actual) {
		return true, nil
	}
	if reflect.TypeOf(expect).String() == "json.Number" {
		expect = fmt.Sprintf("%s", expect)
		actual = toString(actual)
	}
	result := reflect.DeepEqual(actual, expect)
	return result, nil
}

func isNil(val interface{}) bool {
	v := reflect.ValueOf(val)
	return v.IsValid() && v.Kind() == reflect.Ptr && v.IsNil()
}

func toString(val interface{}) string {
	switch val.(type) {
	default:
		return fmt.Sprintf("%s", val)
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", val)
	case float32, float64:
		return fmt.Sprintf("%f", val)
	}
}
