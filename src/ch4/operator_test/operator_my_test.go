package operator

import "testing"

func TestSlice(t *testing.T) {
	arr := [...]int{1,2,3,4}
	arr1 := [...]int{1,2,3,4}
	arr2 := [...]int{1,3,2,4}
	t.Log(arr == arr1)
	t.Log(arr1 == arr2)
}
