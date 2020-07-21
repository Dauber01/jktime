package pipe_filter

type StraightPipeline struct {
	Name    string
	Filters *[]Filter
}

func NewStraightPipeline(name string, filters ...Filter) *StraightPipeline {
	return &StraightPipeline{
		Name:    name,
		Filters: &filters,
	}
}

func (st *StraightPipeline) Process(data Request) (Reponse, error) {
	var ret interface{}
	var err error
	for _, filter := range *st.Filters {
		ret, err = filter.Process(data)
		if err != nil {
			return ret, err
		}
		data = ret
	}
	return ret, err
}
