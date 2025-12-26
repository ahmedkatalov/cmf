package service

import (
	"backend/internal/repository"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	users *repository.UserRepo
}

func NewUserService(users *repository.UserRepo) *UserService {
	return &UserService{users: users}
}

// Создание сотрудника
func (s *UserService) Create(ctx context.Context, orgID string, branchID *string, email, password, role string) (string, error) {
	if role == "owner" {
		return "", errors.New("cannot create owner")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return s.users.Create(ctx, orgID, branchID, email, string(hash), role)
}
