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
