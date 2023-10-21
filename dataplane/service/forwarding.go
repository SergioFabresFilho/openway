package service

import (
	"strings"

	"github.com/sirupsen/logrus"
	"ssff.com.br/openway/common/model"
	"ssff.com.br/openway/dataplane/client"
)

type ForwardingService struct {
	client          *client.ForwardingClient
	apiService      *ApiService
	endpointService *EndpointService
}

func NewForwardingService() *ForwardingService {
	return &ForwardingService{
		client.NewForwardingClient(),
		NewApiService(),
		NewEndpointService(),
	}
}

func (s *ForwardingService) Forward(r model.Request) (*model.Response, error) {
	prefix := getPrefix(r.Path)
	api, err := s.apiService.FindByPrefix(prefix)
	if err != nil {
		logrus.Error("Api not fouund:", prefix)
		return nil, err
	}
	r.Path = extractResourcePath(&r)

	_, err = s.endpointService.FindByApiIdAndPathAndMethod(api.ID, r.Path, r.Method)
	if err != nil {
		logrus.Error("Endpoint not found:", r.Path, r.Method)
		return nil, err
	}

	resp := s.client.Forward(r, api.Host)

	return resp, nil
}

func getPrefix(path string) string {
	return "/" + strings.Split(path, "/")[1]
}

func extractResourcePath(r *model.Request) string {
	path := strings.Split(r.Path, "/")
	return "/" + strings.Join(path[2:], "/")
}
