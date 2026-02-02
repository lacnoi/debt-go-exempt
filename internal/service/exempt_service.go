package service

import (
	"context"

	"go.uber.org/zap"

	"github.com/lacnoi/debt-go-exempt/internal/repo"
)

type ExemptService struct {
	repo   *repo.ExemptRepo
	logger *zap.Logger
}

func NewExemptService(r *repo.ExemptRepo, logger *zap.Logger) *ExemptService {
	return &ExemptService{repo: r, logger: logger}
}

func (s *ExemptService) Create(ctx context.Context, employeeID, reason string) (string, error) {
	id, err := s.repo.Insert(ctx, employeeID, reason)
	if err != nil {
		s.logger.Error("create exempt failed", zap.Error(err))
		return "", err
	}
	return id, nil
}

func (s *ExemptService) GetByID(ctx context.Context, id string) (map[string]any, error) {
	return s.repo.GetByID(ctx, id)
}
