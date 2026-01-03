package service

import (
	"context"

	"backend/internal/repository"
)

type BranchService struct {
	branches *repository.BranchRepo
	users    *repository.UserRepo
}

func NewBranchService(branches *repository.BranchRepo, users *repository.UserRepo) *BranchService {
	return &BranchService{
		branches: branches,
		users:    users,
	}
}

func (s *BranchService) Create(ctx context.Context, orgID, name, address string) (string, error) {
	return s.branches.Create(ctx, orgID, name, address)
}

func (s *BranchService) List(ctx context.Context, orgID string) ([]repository.Branch, error) {
	return s.branches.ListByOrg(ctx, orgID)
}

func (s *BranchService) GetByID(ctx context.Context, orgID, branchID string) (*repository.Branch, error) {
	return s.branches.GetByID(ctx, orgID, branchID)
}

func (s *BranchService) ListUsers(ctx context.Context, orgID, branchID string) ([]repository.UserPublic, error) {
	return s.users.ListByBranchPublic(ctx, orgID, branchID)
}
