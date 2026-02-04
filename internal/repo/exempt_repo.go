package repo

import (
	"context"

	"github.com/lacnoi/debt-go-exempt/internal/db"
	"github.com/lacnoi/debt-go-exempt/internal/domain/exempt"
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

func (r *ExemptRepo) GetByBaNo(ctx context.Context, baNo string) ([]exempt.Exempt, error) {
	sql := `
		SELECT
			ca_no,
			ba_no,
			mobile_num,
			mode_id,
			effective_dat,
			end_dat,
			created,
			created_by,
			last_upd,
			last_upd_by
		FROM dcc_exempt
		WHERE ba_no = $1
		ORDER BY effective_dat
	`

	rows, err := r.db.Pool.Query(ctx, sql, baNo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []exempt.Exempt

	for rows.Next() {
		var e exempt.Exempt
		if err := rows.Scan(
			&e.CANo,
			&e.BANo,
			&e.MobileNum,
			&e.ModeID,
			&e.EffectiveDat,
			&e.EndDat,
			&e.Created,
			&e.CreatedBy,
			&e.LastUpd,
			&e.LastUpdBy,
		); err != nil {
			return nil, err
		}
		result = append(result, e)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// ไม่มี row → return empty slice (ไม่ใช่ error)
	return result, nil
}
