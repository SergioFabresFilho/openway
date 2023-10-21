package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"ssff.com.br/openway/common/model"
	"ssff.com.br/openway/controlplane/handler/in"
	"ssff.com.br/openway/controlplane/handler/out"
	"ssff.com.br/openway/controlplane/service"
)

type EndpointHandler struct {
	endpointService *service.EndpointService
}

func NewEndpointHandler() *EndpointHandler {
	return &EndpointHandler{
		service.NewEndpointService(),
	}
}

func (h *EndpointHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req in.CreateEndpointRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logrus.Error("CreateEndpoint bad request ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	idString := mux.Vars(r)["apiId"]
	apiID, err := uuid.Parse(idString)
	if err != nil {
		logrus.Errorf("CreateEndpoint bad request. Invalid apiID: [%s] ", idString)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	endpoint := &model.Endpoint{
		Description: req.Description,
		Path:        req.Path,
		Method:      req.Method,
		ApiID:       apiID,
	}

	id, err := h.endpointService.Create(endpoint)
	if err != nil {
		logrus.Error("CreateEndpoint failed", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := &out.CreateEndpointResponse{ID: id}
	json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusCreated)
}
