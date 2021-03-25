package my_pipe_filter

import "testing"

func TestCombineFilter_Process(t *testing.T) {
	strs := "3,4,1,6,9"
	splitFilter := NewSplitFilter(",")
	tointFilter := NewToIntFilter()
	sumFilter := NewSumFilter()

	combineFIlter := NewCombineFilter(splitFilter, tointFilter, sumFilter)
	if ret, err := combineFIlter.Process(strs); err != nil {
		t.Log(err)
	} else {
		t.Log(ret)
	}

}
