package obj_pool

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

type ReuseObj struct {
}

type ObjMyPool struct {
	objs chan *ReuseObj
}

func NewObjMyPool(size int) *ObjMyPool {
	pool := new(ObjMyPool)
	pool.objs = make(chan *ReuseObj, size)
	for i := 0; i < size; i++ {
		pool.objs <- &ReuseObj{}
	}
	return pool
}

func (pool *ObjMyPool) getMyObj(timeout time.Duration) (*ReuseObj, error) {
	select {
	case res := <-pool.objs:
		return res, nil
	case <-time.After(timeout):
		return nil, errors.New("获取对象超时，对象池中无数据\"")
	}
}

func (pool *ObjMyPool) relaeaseMyObj(obj *ReuseObj) error {
	select {
	case pool.objs <- obj:
		return nil
	default:
		return errors.New("归还失败")
	}

}

func TestMyObj(t *testing.T) {
	pool := NewObjMyPool(10)

	for i := 0; i < 20; i++ {
		if obj, err := pool.getMyObj(1000); err != nil {
			fmt.Print("获取对象失败\n")
		} else {
			fmt.Printf("%T\n", obj)
			pool.relaeaseMyObj(obj)
		}
	}
}
