package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"ssff.com.br/openway/controlplane/handler/in"
	"ssff.com.br/openway/controlplane/handler/out"
	"ssff.com.br/openway/controlplane/model"
	"ssff.com.br/openway/controlplane/service"
)

type ApiHandler struct {
	apiService *service.ApiService
}

func NewApiHandler() *ApiHandler {
	return &ApiHandler{
		service.NewApiService(),
	}
}

func (h *ApiHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req in.CreateApiRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logrus.Error("CreateApi bad request", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	api := &model.Api{
		Name:        req.Name,
		Description: req.Description,
		Host:        req.Host,
		Prefix:      req.Prefix,
	}

	id, err := h.apiService.Create(api)
	if err != nil {
		logrus.Error("CreateApi failed", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	res := &out.CreateApiResponse{ID: id}
	json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusCreated)
}
