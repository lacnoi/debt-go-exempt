package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/lacnoi/debt-go-exempt/internal/service"
	"github.com/lacnoi/debt-go-exempt/pkg/response"
)

type ExemptHandler struct {
	svc *service.ExemptService
}

func NewExemptHandler(svc *service.ExemptService) *ExemptHandler {
	return &ExemptHandler{svc: svc}
}

type CreateExemptReq struct {
	EmployeeID string `json:"employeeId"`
	Reason     string `json:"reason"`
}

func (h *ExemptHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateExemptReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, 400, "invalid json")
		return
	}
	if req.EmployeeID == "" || req.Reason == "" {
		response.Error(w, 400, "employeeId and reason are required")
		return
	}

	id, err := h.svc.Create(r.Context(), req.EmployeeID, req.Reason)
	if err != nil {
		response.Error(w, 500, err.Error())
		return
	}
	response.JSON(w, 201, map[string]any{"id": id})
}

func (h *ExemptHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	data, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		response.Error(w, 404, "not found")
		return
	}
	response.JSON(w, 200, data)
}
