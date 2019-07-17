package main

import (
	"fmt"
	"runtime"
)

func trace() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	fmt.Printf("%s:%d %s\n", file, line, f.Name())
}

func foo() {
	fmt.Println("in foo function")
	trace()
}

func main() {
	foo()
	trace()
}
