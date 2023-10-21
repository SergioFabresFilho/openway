package service

import (
	"ssff.com.br/openway/common/model"
	"ssff.com.br/openway/dataplane/repository"
)

type ApiService struct {
	apiRepository *repository.ApiRepository
}

func NewApiService() *ApiService {
	return &ApiService{
		repository.NewApiRepository(),
	}
}

func (s *ApiService) FindByPrefix(prefix string) (*model.Api, error) {
	return s.apiRepository.FindByPrefix(prefix)
}
