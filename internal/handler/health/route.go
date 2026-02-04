package health

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/lacnoi/debt-go-exempt/pkg/response"
)

func Register(r chi.Router, logger *zap.Logger) {
	r.Get("/health", func(w http.ResponseWriter, req *http.Request) {
		response.JSON(w, http.StatusOK, map[string]any{"status": "ok"})
	})
}
