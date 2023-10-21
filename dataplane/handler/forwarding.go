package handler

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"ssff.com.br/openway/common/model"
	"ssff.com.br/openway/dataplane/service"
)

type ForwardingHandler struct {
	forwardingService *service.ForwardingService
}

func NewForwardingHandler() *ForwardingHandler {
	return &ForwardingHandler{
		service.NewForwardingService(),
	}
}

func (h *ForwardingHandler) Forward(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Forwarding request.")
	req := model.Request{
		Method:      r.Method,
		Path:        r.URL.Path,
		Timestamp:   time.Now().UnixMilli(),
		Body:        getBodyAsString(r.Body),
		Headers:     r.Header,
		QueryParams: extractQueryParams(r.URL.RawQuery),
	}

	resp, err := h.forwardingService.Forward(req)
	if err != nil {
		logrus.Error("Error forwarding request.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logrus.Info("Request forwarded.")
	w.WriteHeader(resp.Status)
	w.Write([]byte(resp.Body))
}

func getBodyAsString(body io.ReadCloser) string {
	b, err := io.ReadAll(body)
	if err != nil {
		return ""
	}

	return string(b)
}

func extractQueryParams(queryString string) map[string]string {
	result := make(map[string]string)
	if queryString == "" {
		return result
	}

	query := strings.Split(queryString, "&")
	for _, q := range query {
		pair := strings.Split(q, "=")
		result[pair[0]] = pair[1]
	}

	return result
}
