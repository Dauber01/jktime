package pipe_filter

import (
	"errors"
	"strings"
)

var SplitFilterWrongFormatError = errors.New("input data should be string")

type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

func (st *SplitFilter) Process(data Request) (Reponse, error) {
	v, ok := data.(string)
	if !ok {
		return nil, SplitFilterWrongFormatError
	}
	parts := strings.Split(v, st.delimiter)
	return parts, nil
}
