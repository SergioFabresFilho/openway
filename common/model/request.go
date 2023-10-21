package model

type Request struct {
	Timestamp   int64
	Body        string
	Path        string
	Headers     map[string][]string
	QueryParams map[string]string
	Method      string
}
