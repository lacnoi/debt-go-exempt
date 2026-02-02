package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/lacnoi/debt-go-exempt/internal/db"
	"github.com/lacnoi/debt-go-exempt/internal/handler"
	"github.com/lacnoi/debt-go-exempt/internal/repo"
	"github.com/lacnoi/debt-go-exempt/internal/service"
)

func NewRouter(logger *zap.Logger, database *db.DB) http.Handler {
	r := chi.NewRouter()

	healthH := handler.NewHealthHandler()
	exemptRepo := repo.NewExemptRepo(database)
	exemptSvc := service.NewExemptService(exemptRepo, logger)
	exemptH := handler.NewExemptHandler(exemptSvc)

	r.Get("/health", healthH.Readiness)
	r.Get("/ready", healthH.Readiness)

	r.Route("/api/v1", func(api chi.Router) {
		api.Route("/exempts", func(er chi.Router) {
			er.Post("/", exemptH.Create)
			er.Get("/{id}", exemptH.GetByID)
		})
	})

	return r
}
