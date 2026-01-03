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
	FullName       string
	Phone          string
}

type UserPublic struct {
	ID       string  `json:"id"`
	Email    string  `json:"email"`
	Role     string  `json:"role"`
	BranchID *string `json:"branch_id"`
	FullName string  `json:"full_name"`
	Phone    string  `json:"phone"`
}

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, orgID string, branchID *string, fullName, phone, email, passwordHash, role string) (string, error) {
	var id string
	err := r.db.QueryRow(ctx,
		`INSERT INTO users (organization_id, branch_id, full_name, phone, email, password_hash, role)
		 VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		orgID, branchID, fullName, phone, email, passwordHash, role,
	).Scan(&id)
	return id, err
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*User, error) {
	var u User
	err := r.db.QueryRow(ctx,
		`SELECT id, organization_id, branch_id, email, password_hash, role, full_name, phone
		 FROM users WHERE email=$1`,
		email,
	).Scan(&u.ID, &u.OrganizationID, &u.BranchID, &u.Email, &u.PasswordHash, &u.Role, &u.FullName, &u.Phone)

	if err != nil {
		return nil, err
	}
	return &u, nil
}

// ✅ Для owner панели: получить сотрудников филиала
func (r *UserRepo) ListByBranchPublic(ctx context.Context, orgID, branchID string) ([]UserPublic, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, email, role, branch_id, full_name, phone
		FROM users
		WHERE organization_id=$1 AND branch_id=$2
		ORDER BY created_at ASC
	`, orgID, branchID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []UserPublic
	for rows.Next() {
		var u UserPublic
		if err := rows.Scan(&u.ID, &u.Email, &u.Role, &u.BranchID, &u.FullName, &u.Phone); err != nil {
			return nil, err
		}
		list = append(list, u)
	}
	return list, rows.Err()
}

// ✅ Нужно для удаления филиала: узнать есть ли сотрудники
func (r *UserRepo) CountByBranch(ctx context.Context, orgID, branchID string) (int, error) {
	var count int
	err := r.db.QueryRow(ctx, `
		SELECT COUNT(*)
		FROM users
		WHERE organization_id=$1 AND branch_id=$2
	`, orgID, branchID).Scan(&count)

	return count, err
}
