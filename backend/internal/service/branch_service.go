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
	return &BranchService{branches: branches, users: users}
}


func (s *BranchService) Create(ctx context.Context, orgID, name, address string) (string, error) {
	return s.branches.Create(ctx, orgID, name, address)
}

func (s *BranchService) List(ctx context.Context, orgID string) ([]repository.Branch, error) {
	return s.branches.ListByOrg(ctx, orgID)
}
type BranchWithUsers struct {
	Branch repository.Branch           `json:"branch"`
	Users  []repository.UserPublic     `json:"users"`
}

func (s *BranchService) GetWithUsers(ctx context.Context, orgID, branchID string) (BranchWithUsers, error) {
	b, err := s.branches.GetByID(ctx, orgID, branchID)
	if err != nil {
		return BranchWithUsers{}, err
	}

	users, err := s.users.ListByBranch(ctx, orgID, branchID)
	if err != nil {
		return BranchWithUsers{}, err
	}

	return BranchWithUsers{
		Branch: *b,
		Users:  users,
	}, nil
}
