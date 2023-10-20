package repository

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"ssff.com.br/openway/controlplane/db"
	"ssff.com.br/openway/controlplane/mapper"
	"ssff.com.br/openway/controlplane/model"
)

type EndpointRepository struct {
	db *gorm.DB
}

func NewEndpointRepository() *EndpointRepository {
	return &EndpointRepository{db.GetConnection()}
}

func (r *EndpointRepository) Create(endpoint *model.Endpoint) (uuid.UUID, error) {
	entity := mapper.EndpointToEntity(endpoint)

	err := r.db.Create(&entity).Error
	if err != nil {
		logrus.Error("failed to create api", err)
		return uuid.Nil, err
	}

	return entity.ID, nil
}
