package model

import (
	"github.com/google/uuid"
)

type Api struct {
	ID          uuid.UUID
	Name        string
	Description string
	Host        string
	Prefix      string
}
