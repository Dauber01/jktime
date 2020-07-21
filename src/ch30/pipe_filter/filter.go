package pipe_filter

type Request interface{}

type Reponse interface{}

type Filter interface {
	Process(data Request) (Reponse, error)
}
