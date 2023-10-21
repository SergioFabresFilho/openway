package service

import (
	"github.com/google/uuid"
	"ssff.com.br/openway/common/model"
	"ssff.com.br/openway/dataplane/repository"
)

type EndpointService struct {
	endpointRepository *repository.EndpointRepository
}

func NewEndpointService() *EndpointService {
	return &EndpointService{
		repository.NewEndpointRepository(),
	}
}

func (s *EndpointService) FindByApiIdAndPathAndMethod(apiID uuid.UUID, path string, method string) (*model.Endpoint, error) {
	return s.endpointRepository.FindByApiIdAndPathAndMethod(apiID, path, method)
}
