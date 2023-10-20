package in

type CreateApiRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Host        string `json:"host"`
	Prefix      string `json:"prefix"`
}
