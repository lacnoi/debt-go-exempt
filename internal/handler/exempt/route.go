package exempt

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/lacnoi/debt-go-exempt/internal/db"
	"github.com/lacnoi/debt-go-exempt/internal/repo"
	"github.com/lacnoi/debt-go-exempt/internal/service"
)

func Register(r chi.Router, logger *zap.Logger, database *db.DB) {
	exemptRepo := repo.NewExemptRepo(database)
	exemptSvc := service.NewExemptService(exemptRepo, logger)
	h := NewHandler(exemptSvc, logger)

	r.Route("/exempts", func(er chi.Router) {
		er.Post("/", h.Create)
		er.Get("/{ba_no}", h.GetByBaNo)
	})
}
