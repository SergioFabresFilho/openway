package dataplane

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"ssff.com.br/openway/dataplane/handler"
)

func Init() {
	logrus.Info("Starting dataplane on port 8081")

	h := handler.NewForwardingHandler()

	logrus.Fatal(http.ListenAndServe(":8081", http.HandlerFunc(h.Forward)))
}
