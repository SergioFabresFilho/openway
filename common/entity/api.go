package entity

import (
	"github.com/google/uuid"
)

type Api struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:varchar(255);not null"`
	Host        string    `gorm:"type:varchar(255);not null"`
	Prefix      string    `gorm:"type:varchar(100);not null;unique"`
}
