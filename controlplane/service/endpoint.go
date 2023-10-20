package service

import (
	"github.com/google/uuid"
	"ssff.com.br/openway/controlplane/model"
	"ssff.com.br/openway/controlplane/repository"
)

type EndpointService struct {
	endpointRepository *repository.EndpointRepository
}

func NewEndpointService() *EndpointService {
	return &EndpointService{
		repository.NewEndpointRepository(),
	}
}

func (s *EndpointService) Create(endpoint *model.Endpoint) (uuid.UUID, error) {
	return s.endpointRepository.Create(endpoint)
}
