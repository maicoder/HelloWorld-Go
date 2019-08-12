package pipefilter

func NewStraigtPipeline(name string, filters ...Filter) *StraigtPipeline {
	return &StraigtPipeline{
		Name:    name,
		Filters: &filters,
	}
}

type StraigtPipeline struct {
	Name    string
	Filters *[]Filter
}

func (f *StraigtPipeline) Process(data Request) (Response, error) {
	var ret interface{}
	var err error
	for _, filter := range *f.Filters {
		ret, err = filter.Process(data)
		if err != nil {
			return ret, err
		}
		data = ret
	}
	return ret, err

}
