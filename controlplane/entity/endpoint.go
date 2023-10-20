package entity

import "github.com/google/uuid"

type Endpoint struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Description string    `gorm:"type:varchar(255);not null"`
	Path        string    `gorm:"type:varchar(255);not null"`
	Method      string    `gorm:"type:varchar(15);not null"`
	ApiID       uuid.UUID `gorm:"type:uuid;not null"`
}
