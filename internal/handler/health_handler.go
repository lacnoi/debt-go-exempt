package handler

import (
	"net/http"

	"github.com/lacnoi/debt-go-exempt/pkg/response"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler { return &HealthHandler{} }

func (h *HealthHandler) Readiness(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, 200, map[string]any{"status": "ok"})
}
