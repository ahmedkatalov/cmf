package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Branch struct {
	ID             string `json:"id"`
	OrganizationID string `json:"organization_id"`
	Name           string `json:"name"`
	Address        string `json:"address"`
}

type BranchRepo struct {
	db *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) *BranchRepo {
	return &BranchRepo{db: db}
}

func (r *BranchRepo) Create(ctx context.Context, orgID, name, address string) (string, error) {
	var id string
	err := r.db.QueryRow(ctx,
		`INSERT INTO branches (organization_id, name, address)
		 VALUES ($1, $2, $3) RETURNING id`,
		orgID, name, address,
	).Scan(&id)
	return id, err
}

// ✅ ВОТ ЭТОГО МЕТОДА У ТЕБЯ СЕЙЧАС НЕТ
func (r *BranchRepo) ListByOrg(ctx context.Context, orgID string) ([]Branch, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, organization_id, name, COALESCE(address,'')
		 FROM branches
		 WHERE organization_id = $1
		 ORDER BY created_at ASC`,
		orgID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Branch
	for rows.Next() {
		var b Branch
		if err := rows.Scan(&b.ID, &b.OrganizationID, &b.Name, &b.Address); err != nil {
			return nil, err
		}
		list = append(list, b)
	}

	return list, nil
}
