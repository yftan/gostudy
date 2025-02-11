package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// TODO：这个地方为什么要关闭ch
		close(ch)
		//ch <- 11
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 11; i++ {
			//if data,ok := <-ch; ok {
			//	fmt.Println(data)
			//} else {
			//	break
			//}
			data := <-ch
			fmt.Println(data)
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	//wg.Add(1)
	//dataReceiver(ch, &wg)
	wg.Wait()
}
