package pipe_filter

import "errors"

var SumFilterWrongFormatError = errors.New("input data should be []int")

type SumFilter struct {
}

func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

func (sf *SumFilter) Process(data Request) (Reponse, error) {
	nums, ok := data.([]int)
	if !ok {
		return nil, SumFilterWrongFormatError
	}
	ret := 0
	for _, v := range nums {
		ret += v
	}
	return ret, nil
}
