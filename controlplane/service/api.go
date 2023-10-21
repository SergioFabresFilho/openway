package service

import (
	"github.com/google/uuid"
	"ssff.com.br/openway/common/model"
	"ssff.com.br/openway/controlplane/repository"
)

type ApiService struct {
	apiRepository *repository.ApiRepository
}

func NewApiService() *ApiService {
	return &ApiService{
		repository.NewApiRepository(),
	}
}

func (s *ApiService) Create(api *model.Api) (uuid.UUID, error) {
	return s.apiRepository.Create(api)
}
