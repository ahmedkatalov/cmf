package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID             string
	OrganizationID string
	BranchID       *string
	Email          string
	PasswordHash   string
	Role           string
}

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, orgID string, branchID *string, email, passwordHash, role string) (string, error) {
	var id string
	err := r.db.QueryRow(ctx,
		`INSERT INTO users (organization_id, branch_id, email, password_hash, role)
		 VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		orgID, branchID, email, passwordHash, role,
	).Scan(&id)
	return id, err
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*User, error) {
	var u User
	err := r.db.QueryRow(ctx,
		`SELECT id, organization_id, branch_id, email, password_hash, role
		 FROM users WHERE email=$1`,
		email,
	).Scan(&u.ID, &u.OrganizationID, &u.BranchID, &u.Email, &u.PasswordHash, &u.Role)

	if err != nil {
		return nil, err
	}
	return &u, nil
}
