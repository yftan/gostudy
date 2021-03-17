package csp

import (
	"fmt"
	"testing"
	"time"
)

func Service() string {
	return "Sync Service"
}

func AsyncService1() chan string {
	ret := make(chan string, 1)
	ret <- "Async Service"
	return ret
}

func OtherWorks() {
	fmt.Println("Other work start...")
	time.Sleep(1000)
	fmt.Println("Other work done")
}

func TestService1(t *testing.T) {
	ret := AsyncService1()
	OtherWorks()
	fmt.Println(<-ret)
}
