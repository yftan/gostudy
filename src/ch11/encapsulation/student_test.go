package encapsulation

import (
	"fmt"
	"reflect"
	"testing"
)

type Student struct {
	id   int
	name string
	sex  string
}

func (stu *Student) ToString() {
	fmt.Printf("id:%d,name:%s,sex:%s", stu.id, stu.name, stu.sex)
}

func TestStu(t *testing.T) {
	stu := Student{
		id:   1,
		name: "tyf",
		sex:  "man",
	}
	stu1 := &Student{
		id:   2,
		name: "mx",
		sex:  "women",
	}
	fmt.Println(reflect.TypeOf(&stu))
	fmt.Println(reflect.TypeOf(*stu1))
	stu.ToString()
}

func TestPoint(t *testing.T) {
	var i int
	i = 4
	fmt.Println(&i)
	fmt.Println(*&i)
}
