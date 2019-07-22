package example

import (
	"reflect"
	"testing"

	"github.com/lxhope/go-practice/tests/tdt"
)

var person = &Person{Name: "Jack"}

func TestPersonWalk(t *testing.T) {
	fnName := "Walk"
	caseFile := "cases/walk.json"
	method := reflect.ValueOf(person).MethodByName(fnName).Interface()
	cases, _ := tdt.LoadCases(caseFile)
	tdt.RunTest(t, method, cases)
}

func TestPersonSpeak(t *testing.T) {
	fnName := "Speak"
	caseFile := "cases/speak.json"
	method := reflect.ValueOf(person).MethodByName(fnName).Interface()
	cases, _ := tdt.LoadCases(caseFile)
	tdt.RunTest(t, method, cases)
}

func TestPersonEat(t *testing.T) {
	fnName := "Eat"
	caseFile := "cases/eat.json"
	method := reflect.ValueOf(person).MethodByName(fnName).Interface()
	cases, _ := tdt.LoadCases(caseFile)
	tdt.RunTest(t, method, cases)
}
