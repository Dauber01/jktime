package pipe_filter

import (
	"errors"
	"strconv"
)

var ToIntFilterWrongFormatError = errors.New("input data should be []string")

type ToIntFilter struct {
}

func NewToIntFilter() *ToIntFilter {
	return &ToIntFilter{}
}

func (ti *ToIntFilter) Process(data Request) (Reponse, error) {
	parts, ok := data.([]string)
	if !ok {
		return nil, ToIntFilterWrongFormatError
	}
	ret := []int{}
	for _, part := range parts {
		s, error := strconv.Atoi(part)
		if error != nil {
			return nil, ToIntFilterWrongFormatError
		}
		ret = append(ret, s)
	}
	return ret, nil
}
