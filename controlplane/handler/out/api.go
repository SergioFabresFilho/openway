package out

import "github.com/google/uuid"

type CreateApiResponse struct {
	ID uuid.UUID `json:"id"`
}
