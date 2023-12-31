package mapper

import (
	"ssff.com.br/openway/common/entity"
	"ssff.com.br/openway/common/model"
)

func ApiToEntity(api *model.Api) *entity.Api {
	return &entity.Api{
		ID:          api.ID,
		Name:        api.Name,
		Description: api.Description,
		Host:        api.Host,
		Prefix:      api.Prefix,
	}
}

func ApiToModel(api *entity.Api) *model.Api {
	return &model.Api{
		ID:          api.ID,
		Name:        api.Name,
		Description: api.Description,
		Host:        api.Host,
		Prefix:      api.Prefix,
	}
}
