package _select

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func AsyncService() chan string {
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
	}()
	return retCh
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}

func TestMySelect(t *testing.T) {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	c3 := make(chan string, 1)

	go func() {
		time.Sleep(4 * time.Second)
		c1 <- "c1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "c2"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		c3 <- "c3"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println(msg1)
	case msg2 := <-c2:
		fmt.Println(msg2)
	case msg3 := <-c3:
		fmt.Println(msg3)
	}

}
