package client

import (
	"io"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"ssff.com.br/openway/common/model"
)

type ForwardingClient struct {
}

func NewForwardingClient() *ForwardingClient {
	return &ForwardingClient{}
}

func (c *ForwardingClient) Forward(req model.Request, host string) *model.Response {
	logrus.Infof("Forwarding request to host %s", host)
	httpReq, err := http.NewRequest(req.Method, buildPath(req, host), nil)
	if err != nil {
		return &model.Response{
			Status:    500,
			Body:      err.Error(),
			Timestamp: time.Now().UnixMilli(),
		}
	}

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return &model.Response{
			Status:    500,
			Body:      err.Error(),
			Timestamp: time.Now().UnixMilli(),
		}
	}

	return &model.Response{
		Status:    resp.StatusCode,
		Body:      getBodyAsString(*resp),
		Timestamp: time.Now().UnixMilli(),
		Headers:   getHeadersAsMap(*resp),
	}
}

func buildPath(req model.Request, host string) string {
	queryString := buildQueryParams(req.QueryParams)
	return host + req.Path + queryString
}

func buildQueryParams(params map[string]string) string {
	str := "?"
	for k, v := range params {
		str += k + "=" + v + "&"
	}

	return str
}

func getBodyAsString(resp http.Response) string {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(body)
}

func getHeadersAsMap(resp http.Response) map[string][]string {
	headers := make(map[string][]string)
	for k, v := range resp.Header {
		headers[k] = v
	}

	return headers
}
