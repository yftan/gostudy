package type_test_my

import "testing"

type MyInt int64

func sum(a MyInt, b MyInt)  MyInt {
	return a+b;
}

func TestSum(t *testing.T)  {
	var a MyInt
	var c int64
	a = 1
	b := &a
	c = int64(a)
	t.Log(a, *b, c)
	t.Logf("%T %T", a, *b)
}