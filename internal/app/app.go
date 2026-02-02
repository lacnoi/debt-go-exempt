package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"github.com/lacnoi/debt-go-exempt/internal/config"
	"github.com/lacnoi/debt-go-exempt/internal/db"
)

type App struct {
	cfg    config.Config
	logger *zap.Logger
	db     *db.DB
	server *http.Server
}

func New() (*App, error) {
	_ = godotenv.Load()

	logger, _ := zap.NewProduction()

	cfg := config.Load()

	database, err := db.New(db.Options{
		DBURL:  cfg.DBURL,
		Logger: logger,
	})
	if err != nil {
		return nil, err
	}

	r := NewRouter(logger, database)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: r,
	}

	return &App{
		cfg:    cfg,
		logger: logger,
		db:     database,
		server: srv,
	}, nil
}

func (a *App) Run() error {
	a.logger.Info("starting server",
		zap.String("app", a.cfg.AppName),
		zap.String("port", a.cfg.Port),
		zap.String("env", os.Getenv("ENV")),
	)
	return a.server.ListenAndServe()
}
