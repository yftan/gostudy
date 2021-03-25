package my_pipe_filter

type Filter interface {
	Process(request Request) (Response, error)
}

type Response interface {
}

type Request interface {
}
