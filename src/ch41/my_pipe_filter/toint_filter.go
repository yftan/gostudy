package my_pipe_filter

import (
	"errors"
	"strconv"
)

var IntConvertError = errors.New("整形转换错误")

type ToIntFilter struct {
}

func NewToIntFilter() *ToIntFilter {
	return &ToIntFilter{}
}

func (filter *ToIntFilter) Process(request Request) (Response, error) {
	if strs, ok := request.([]string); !ok {
		return nil, IntConvertError
	} else {
		var rets []int
		for _, str := range strs {
			val, _ := strconv.Atoi(str)
			rets = append(rets, val)
		}
		return rets, nil
	}
}
