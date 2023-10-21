package mapper

import (
	"ssff.com.br/openway/common/entity"
	"ssff.com.br/openway/common/model"
)

func EndpointToEntity(endpoint *model.Endpoint) *entity.Endpoint {
	return &entity.Endpoint{
		ID:          endpoint.ID,
		Description: endpoint.Description,
		Path:        endpoint.Path,
		Method:      endpoint.Method,
		ApiID:       endpoint.ApiID,
	}
}

func EndpointToModel(endpoint *entity.Endpoint) *model.Endpoint {
	return &model.Endpoint{
		ID:          endpoint.ID,
		Description: endpoint.Description,
		Path:        endpoint.Path,
		Method:      endpoint.Method,
		ApiID:       endpoint.ApiID,
	}
}
