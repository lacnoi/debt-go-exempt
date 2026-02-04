package app

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	"github.com/lacnoi/debt-go-exempt/internal/db"
	"github.com/lacnoi/debt-go-exempt/internal/handler/exempt"
	"github.com/lacnoi/debt-go-exempt/internal/handler/health"
)

// NewRouter เป็นตัวรวมทุก route ของ service นี้
func NewRouter(logger *zap.Logger, database *db.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	health.Register(r, logger)

	r.Route("/api/v1", func(api chi.Router) {
		exempt.Register(api, logger, database)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not found", http.StatusNotFound)
	})
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	return r
}
