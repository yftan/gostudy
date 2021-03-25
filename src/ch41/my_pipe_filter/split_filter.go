package my_pipe_filter

import (
	"errors"
	"strings"
)

var SplitStringError = errors.New("分割失败")

type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter: delimiter}
}

func (filter *SplitFilter) Process(request Request) (Response, error) {
	if str, ok := request.(string); !ok {
		return nil, SplitStringError
	} else {
		strs := strings.Split(str, filter.delimiter)
		return strs, nil
	}
}
