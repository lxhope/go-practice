package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("setflags test")

	log.SetPrefix("[log-test] ")
	log.Println("setprefix test")

	f, err := os.OpenFile("test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Errorf("open log file failed, error:%v", err)
	}
	log.SetOutput(f)
	log.Println("set output test")
}
