package service

import (
	"context"
	"errors"
	"time"

	"backend/internal/repository"
)

type TransactionService struct {
	repo *repository.TransactionRepo
}

func NewTransactionService(repo *repository.TransactionRepo) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Create(
	ctx context.Context,
	orgID, branchID string,
	createdBy *string,
	txType string,
	categoryID *string,
	amount int64,
	occurredAt time.Time,
	description *string,
) (string, error) {

	if amount <= 0 {
		return "", errors.New("amount must be > 0")
	}

	switch txType {
	case "income", "expense_company", "expense_people", "transfer_to_owner":
	default:
		return "", errors.New("invalid type")
	}

	return s.repo.Create(ctx, orgID, branchID, createdBy, txType, categoryID, amount, occurredAt, description)
}

func (s *TransactionService) ListByBranch(ctx context.Context, orgID, branchID string, f repository.ListTransactionsFilter) ([]repository.Transaction, error) {
	return s.repo.ListByBranch(ctx, orgID, branchID, f)
}

// ✅ Cancel transaction (самый правильный способ отмены)
func (s *TransactionService) Cancel(
	ctx context.Context,
	orgID, txID string,
	requesterRole string,
	requesterBranchID string,
	requesterUserID string,
	reason *string,
) error {
	tx, err := s.repo.GetByID(ctx, orgID, txID)
	if err != nil {
		return err
	}

	// ✅ Права:
	// owner -> может отменять любую транзакцию
	// остальные -> только в рамках своего филиала
	if requesterRole != "owner" && tx.BranchID != requesterBranchID {
		return errors.New("forbidden")
	}

	// если уже отменена, можно вернуть ошибку
	if tx.IsCancelled {
		return errors.New("already cancelled")
	}

	return s.repo.Cancel(ctx, orgID, txID, requesterUserID, reason)
}
