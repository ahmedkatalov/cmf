package service

import (
	"context"
	"errors"
	"strings"

	"backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	users *repository.UserRepo
}

func NewUserService(users *repository.UserRepo) *UserService {
	return &UserService{users: users}
}

// Создание сотрудника
func (s *UserService) Create(ctx context.Context, orgID string, branchID *string, fullName, phone, email, password, role string) (string, error) {
	if role == "owner" {
		return "", errors.New("cannot create owner")
	}

	fullName = strings.TrimSpace(fullName)
	phone = strings.TrimSpace(phone)
	email = strings.TrimSpace(email)

	if fullName == "" {
		return "", errors.New("full_name required")
	}
	if phone == "" {
		return "", errors.New("phone required")
	}
	if email == "" {
		return "", errors.New("email required")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return s.users.Create(ctx, orgID, branchID, fullName, phone, email, string(hash), role)
}
