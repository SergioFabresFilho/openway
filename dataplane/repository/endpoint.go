package repository

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"ssff.com.br/openway/common/db"
	"ssff.com.br/openway/common/entity"
	"ssff.com.br/openway/common/mapper"
	"ssff.com.br/openway/common/model"
)

type EndpointRepository struct {
	db *gorm.DB
}

func NewEndpointRepository() *EndpointRepository {
	return &EndpointRepository{db.GetConnection()}
}

func (r *EndpointRepository) FindByApiIdAndPathAndMethod(apiID uuid.UUID, path string, method string) (*model.Endpoint, error) {
	var endpoint entity.Endpoint
	if err := r.db.Where("api_id = ? AND path = ? AND method = ?", apiID, path, method).First(&endpoint).Error; err != nil {
		logrus.Error("Endpoint not found.")
		return nil, err
	}

	model := mapper.EndpointToModel(&endpoint)
	return model, nil
}
