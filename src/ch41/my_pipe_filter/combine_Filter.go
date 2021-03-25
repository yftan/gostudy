package my_pipe_filter

type CombineFilter struct {
	Filters *[]Filter
}

func NewCombineFilter(filters ...Filter) *CombineFilter {
	return &CombineFilter{
		Filters: &filters,
	}
}

func (filter *CombineFilter) Process(request Request) (Response, error) {
	var ret interface{}
	var err error
	for _, f := range *filter.Filters {
		ret, err = f.Process(request)
		if err != nil {
			return nil, err
		}
		request = ret
	}
	return ret, nil
}
