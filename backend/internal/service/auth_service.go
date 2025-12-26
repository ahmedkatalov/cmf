package service

import (
	"context"
	"errors"
	"time"

	"backend/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	users    *repository.UserRepo
	orgs     *repository.OrgRepo
	branches *repository.BranchRepo

	jwtSecret string
}

func NewAuthService(users *repository.UserRepo, orgs *repository.OrgRepo, branches *repository.BranchRepo, jwtSecret string) *AuthService {
	return &AuthService{
		users: users,
		orgs: orgs,
		branches: branches,
		jwtSecret: jwtSecret,
	}
}

func (s *AuthService) RegisterRoot(ctx context.Context, orgName, email, password string) (string, error) {
	orgID, err := s.orgs.Create(ctx, orgName)
	if err != nil {
		return "", err
	}

	mainBranchID, err := s.branches.Create(ctx, orgID, "Главная точка", "")
	if err != nil {
		return "", err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	branchID := mainBranchID
	userID, err := s.users.Create(ctx, orgID, &branchID, email, string(hash), "owner")
	if err != nil {
		return "", err
	}

	return s.makeToken(userID, orgID, &branchID, "owner", email)
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	u, err := s.users.GetByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	return s.makeToken(u.ID, u.OrganizationID, u.BranchID, u.Role, u.Email)
}

func (s *AuthService) makeToken(userID, orgID string, branchID *string, role, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"org_id":  orgID,
		"role":    role,
		"email":   email,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	}

	if branchID != nil {
		claims["branch_id"] = *branchID
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret)) // ✅ без лишней скобки
}
