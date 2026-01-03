package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrgRepo struct {
	db *pgxpool.Pool
}

func NewOrgRepo(db *pgxpool.Pool) *OrgRepo {
	return &OrgRepo{db: db}
}

func (r *OrgRepo) Create(ctx context.Context, name string) (string, error) {
	var id string
	err := r.db.QueryRow(ctx,
		`INSERT INTO organizations (name) VALUES ($1) RETURNING id`,
		name,
	).Scan(&id)
	return id, err
}
func (r *BranchRepo) Update(ctx context.Context, orgID, branchID, name, address string) error {
	_, err := r.db.Exec(ctx, `
		UPDATE branches
		SET name=$1, address=$2
		WHERE organization_id=$3 AND id=$4
	`, name, address, orgID, branchID)
	return err
}

func (r *BranchRepo) Delete(ctx context.Context, orgID, branchID string) error {
	_, err := r.db.Exec(ctx, `
		DELETE FROM branches
		WHERE organization_id=$1 AND id=$2
	`, orgID, branchID)
	return err
}
