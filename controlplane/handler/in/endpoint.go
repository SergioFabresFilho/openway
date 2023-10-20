package in

type CreateEndpointRequest struct {
	Description string `json:"description"`
	Path        string `json:"path"`
	Method      string `json:"method"`
}
