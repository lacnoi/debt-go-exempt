package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Options struct {
	DBURL  string
	Logger *zap.Logger
}

type DB struct {
	Pool   *pgxpool.Pool
	Logger *zap.Logger
}

func New(opt Options) (*DB, error) {
	cfg, err := pgxpool.ParseConfig(opt.DBURL)
	if err != nil {
		return nil, err
	}
	cfg.MaxConns = 10
	cfg.MinConns = 2
	cfg.MaxConnLifetime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return &DB{Pool: pool, Logger: opt.Logger}, nil
}
