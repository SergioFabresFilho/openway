package repository

import (
	"gorm.io/gorm"
	"ssff.com.br/openway/common/db"
	"ssff.com.br/openway/common/entity"
	"ssff.com.br/openway/common/mapper"
	"ssff.com.br/openway/common/model"
)

type ApiRepository struct {
	db *gorm.DB
}

func NewApiRepository() *ApiRepository {
	return &ApiRepository{db.GetConnection()}
}

func (r *ApiRepository) FindByPrefix(prefix string) (*model.Api, error) {
	var api entity.Api
	err := r.db.Where("prefix = ?", prefix).First(&api).Error
	if err != nil {
		return nil, err
	}

	model := mapper.ApiToModel(&api)
	return model, nil
}
