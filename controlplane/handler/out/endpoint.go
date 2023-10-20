package out

import "github.com/google/uuid"

type CreateEndpointResponse struct {
	ID uuid.UUID `json:"id"`
}
