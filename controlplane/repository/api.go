package repository

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"ssff.com.br/openway/controlplane/db"
	"ssff.com.br/openway/controlplane/mapper"
	"ssff.com.br/openway/controlplane/model"
)

type ApiRepository struct {
	db *gorm.DB
}

func NewApiRepository() *ApiRepository {
	return &ApiRepository{db.GetConnection()}
}

func (r *ApiRepository) Create(api *model.Api) (uuid.UUID, error) {
	entity := mapper.ApiToEntity(api)

	err := r.db.Create(&entity).Error
	if err != nil {
		logrus.Error("failed to create api", err)
		return uuid.Nil, err
	}

	return entity.ID, nil
}
