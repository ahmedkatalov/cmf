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

type UserClaims struct {
	UserID    string  `json:"user_id"`
	OrgID     string  `json:"org_id"`
	BranchID  *string `json:"branch_id,omitempty"`
	Role      string  `json:"role"`
	Email     string  `json:"email"`
	ExpiresAt int64   `json:"exp"`
}

// ParseToken parses a JWT token string and returns the user claims.
func (s *AuthService) ParseToken(tokenStr string) (*UserClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	uc := &UserClaims{}
	if v, ok := claims["user_id"].(string); ok {
		uc.UserID = v
	} else {
		return nil, errors.New("user_id missing in token")
	}
	if v, ok := claims["org_id"].(string); ok {
		uc.OrgID = v
	} else {
		return nil, errors.New("org_id missing in token")
	}
	if v, ok := claims["role"].(string); ok {
		uc.Role = v
	}
	if v, ok := claims["email"].(string); ok {
		uc.Email = v
	}
	if v, ok := claims["branch_id"].(string); ok {
		uc.BranchID = &v
	}
	if v, ok := claims["exp"].(float64); ok {
		uc.ExpiresAt = int64(v)
	}

	return uc, nil
}
