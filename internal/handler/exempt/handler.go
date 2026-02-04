package exempt

import (
	"encoding/json"
	"net/http"

	"errors"

	"github.com/go-chi/chi/v5"
	"github.com/lacnoi/debt-go-exempt/internal/apperror"
	"github.com/lacnoi/debt-go-exempt/internal/service"
	"github.com/lacnoi/debt-go-exempt/pkg/response"
	"go.uber.org/zap"
)

type Handler struct {
	svc    *service.ExemptService
	logger *zap.Logger
}

func NewHandler(svc *service.ExemptService, logger *zap.Logger) *Handler {
	return &Handler{svc: svc, logger: logger}
}

type CreateRequest struct {
	EmployeeID string `json:"employeeId"`
	Reason     string `json:"reason"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request")
		return
	}

	id, err := h.svc.Create(r.Context(), req.EmployeeID, req.Reason)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusCreated, map[string]any{"id": id})
}

func (h *Handler) GetByBaNo(w http.ResponseWriter, r *http.Request) {
	baNo := chi.URLParam(r, "ba_no")
	h.logger.Info("get exempt by ba_no - start", zap.String("baNo", baNo))

	data, err := h.svc.GetByBaNo(r.Context(), baNo)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			response.Error(w, http.StatusNotFound, "not found")
			return
		}

		h.logger.Error("get exempt by ba_no failed",
			zap.String("baNo", baNo),
			zap.Error(err),
		)
		response.Error(w, http.StatusInternalServerError, apperror.ErrInternal.Error())
		return
	}

	response.JSON(w, http.StatusOK, data)
}
