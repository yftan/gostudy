package obj_cache

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}

	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	//runtime.GC()	//GC 会清除sync.pool中缓存的对象
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
	v2, _ := pool.Get().(int)
	fmt.Println(v2)
}

func TestSyncStringPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("创建一个新的字符串实例")
			return "test"
		},
	}
	pool.Put("test")
	pool.Put(123)
	pool.Put("test")
	for i := 0; i < 4; i++ {
		str := pool.Get()
		fmt.Println(str)
	}
}

func TestSyncPoolInMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Printf("i: %d, value: %d\n", id, pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
