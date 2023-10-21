package controlplane

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"ssff.com.br/openway/controlplane/handler"
)

func InitAsync() {
	logrus.Info("Starting controlplane...")
	r := mux.NewRouter()
	r.Use(contentTypeMiddleware)

	apiRoutes(r)
	endpointRoutes(r)

	go func() {
		logrus.Info("Starting controlplane on port 8080")
		logrus.Fatal(http.ListenAndServe(":8080", r))
	}()
}

func apiRoutes(r *mux.Router) {
	logrus.Info("Creating api routes...")

	h := handler.NewApiHandler()

	r.HandleFunc("/api/v1/apis", h.Create).Methods("POST")
}

func endpointRoutes(r *mux.Router) {
	logrus.Info("Creating endpoint routes...")

	h := handler.NewEndpointHandler()

	r.HandleFunc("/api/v1/apis/{apiId}/endpoints", h.Create).Methods("POST")
}
