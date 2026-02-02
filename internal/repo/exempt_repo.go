package repo

import (
	"context"

	"github.com/lacnoi/debt-go-exempt/internal/db"
)

type ExemptRepo struct {
	db *db.DB
}

func NewExemptRepo(database *db.DB) *ExemptRepo {
	return &ExemptRepo{db: database}
}

func (r *ExemptRepo) Insert(ctx context.Context, employeeID, reason string) (string, error) {
	sql := `
		INSERT INTO debt_exempt_request (employee_id, reason)
		VALUES ($1, $2)
		RETURNING id
	`
	var id string
	err := r.db.Pool.QueryRow(ctx, sql, employeeID, reason).Scan(&id)
	return id, err
}

func (r *ExemptRepo) GetByID(ctx context.Context, id string) (map[string]any, error) {
	sql := `
		SELECT id, employee_id, reason, created_at
		FROM debt_exempt_request
		WHERE id = $1
	`
	var (
		_id        string
		employeeID string
		reason     string
		createdAt  any
	)
	err := r.db.Pool.QueryRow(ctx, sql, id).Scan(&_id, &employeeID, &reason, &createdAt)
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"id":         _id,
		"employeeId": employeeID,
		"reason":     reason,
		"createdAt":  createdAt,
	}, nil
}
