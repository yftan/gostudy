package my_pipe_filter

import "errors"

var SumError = errors.New("求和出错")

type SumFilter struct {
}

func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

func (filter *SumFilter) Process(request Request) (Response, error) {
	ret := 0
	if vals, ok := request.([]int); !ok {
		return nil, SumError
	} else {
		for _, val := range vals {
			ret += val
		}
		return ret, nil
	}
}
