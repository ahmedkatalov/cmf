package service

import (
	"context"

	"backend/internal/repository"
)

type BranchService struct {
	branches *repository.BranchRepo
}

func NewBranchService(branches *repository.BranchRepo) *BranchService {
	return &BranchService{branches: branches}
}

func (s *BranchService) Create(ctx context.Context, orgID, name, address string) (string, error) {
	return s.branches.Create(ctx, orgID, name, address)
}

func (s *BranchService) List(ctx context.Context, orgID string) ([]repository.Branch, error) {
	return s.branches.ListByOrg(ctx, orgID)
}
