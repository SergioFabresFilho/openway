package model

import "github.com/google/uuid"

type Endpoint struct {
	ID          uuid.UUID
	Description string
	Path        string
	Method      string
	ApiID       uuid.UUID
}
