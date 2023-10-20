package mapper

import (
	"ssff.com.br/openway/controlplane/entity"
	"ssff.com.br/openway/controlplane/model"
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

func EncpointToModel(endpoint *entity.Endpoint) *model.Endpoint {
	return &model.Endpoint{
		ID:          endpoint.ID,
		Description: endpoint.Description,
		Path:        endpoint.Path,
		Method:      endpoint.Method,
		ApiID:       endpoint.ApiID,
	}
}
