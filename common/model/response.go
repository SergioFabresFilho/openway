package model

type Response struct {
	Timestamp int64
	Body      string
	Headers   map[string][]string
	Status    int
}
