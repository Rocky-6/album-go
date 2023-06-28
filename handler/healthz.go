package handler

import (
	"album/model"
	"encoding/json"
	"net/http"
)

type HealthzHandler struct{}

func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(&model.HealthzResponse{Message: "OK"})
	if err != nil {
		panic(err)
	}
}
