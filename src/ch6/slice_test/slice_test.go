package slice_test

import (
	"fmt"
	"testing"
)

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	//t.Log(s2[0], s2[1], s2[2], s2[3], s2[4]) //panic: runtime error: index out of range
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))
	//s2[4] = 5 //panic: runtime error: index out of range [recovered]
	//t.Log(len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknow"
	t.Log(Q2)
}

func TestSliceCompare(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	//invalid operation: a == b (slice can only be compared to nil)
	//if a == b {
	//	t.Log("equal")
	//}
	if a != nil {
		t.Log(a)
	}
	if b != nil {
		t.Log(b)
	}

	c := [2]int{1, 2}
	d := [2]int{1, 2}
	if c == d {
		t.Log("c == b")
	}
}

/**
测试切片传递的可能出现的问题
*/
func TestParam(t *testing.T) {
	// test1
	//a := []int{0, 1}
	//foo1(a)
	//fmt.Printf("%p, %p\n", &a, a)

	// test2
	//b := make([]int, 2, 10)
	//b[0] = 0
	//b[1] = 1
	//foo2(b)
	//fmt.Printf("%p, %p, %v\n", &b, b, b)

	// test3 如果希望返回的值是修改后的值
	c := []int{0, 1}
	c = foo3(c)
	fmt.Printf("%p, %p, %v\n", &c, c, c)
}

func foo1(a []int) {
	// 如果加上这个追加，就不会影响传入的数据原始的a[0]还是0，不会变成100。原因是容量不够，从新开辟了新的地址
	//a = append(a, 5,6,7,8)
	a[0] = 100
	fmt.Printf("%p, %p, %v\n", &a, a, a)
}

func foo2(a []int) {
	// 如果使用的是make创建的，设置的初始cap，append的数量不大于cap，就不会重新分配地址。
	a = append(a, 5, 6, 7, 8)
	a[0] = 100
	fmt.Printf("%p, %p, %v\n", &a, a, a)
}

func foo3(a []int) []int {
	a = append(a, 5, 6, 7, 8)
	a[0] = 100
	fmt.Printf("%p, %p, %v\n", &a, a, a)
	return a
}
