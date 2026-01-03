package service

import (
	"context"
	"errors"
	"strings"

	"backend/internal/repository"
)

type BranchService struct {
	branches *repository.BranchRepo
	users    *repository.UserRepo
	txs      *repository.TransactionRepo
}

func NewBranchService(branches *repository.BranchRepo, users *repository.UserRepo, txs *repository.TransactionRepo) *BranchService {
	return &BranchService{
		branches: branches,
		users:    users,
		txs:      txs,
	}
}

func (s *BranchService) Create(ctx context.Context, orgID, name, address string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return "", errors.New("branch name required")
	}
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

// ✅ Update branch (owner или admin своего филиала - проверяется в handler)
func (s *BranchService) Update(ctx context.Context, orgID, branchID, name, address string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return errors.New("name required")
	}
	return s.branches.Update(ctx, orgID, branchID, name, address)
}

// ✅ Delete branch (owner only) + строгие проверки
func (s *BranchService) Delete(ctx context.Context, orgID, branchID string) error {
	// временно запрещаем удаление, если есть users или transactions
	userCount, err := s.users.CountByBranch(ctx, orgID, branchID)
	if err != nil {
		return err
	}
	if userCount > 0 {
		return errors.New("cannot delete branch: branch has users")
	}

	txCount, err := s.txs.CountByBranch(ctx, orgID, branchID)
	if err != nil {
		return err
	}
	if txCount > 0 {
		return errors.New("cannot delete branch: branch has transactions")
	}

	// ✅ В будущем сюда добавим проверку contracts (active contracts)
	// activeContractsCount := ...
	// if activeContractsCount > 0 { return errors.New("branch has active contracts") }

	return s.branches.Delete(ctx, orgID, branchID)
}
