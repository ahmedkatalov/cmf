package service

import (
	"context"
	"time"

	"backend/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SummaryService struct {
	db *pgxpool.Pool
}

func NewSummaryService(db *pgxpool.Pool) *SummaryService {
	return &SummaryService{db: db}
}

type Summary struct {
	Income          int64 `json:"income"`
	ExpenseCompany  int64 `json:"expense_company"`
	ExpensePeople   int64 `json:"expense_people"`
	TransferToOwner int64 `json:"transfer_to_owner"`

	TotalExpenses int64 `json:"total_expenses"`   // company + people
	Net           int64 `json:"net"`              // income - total_expenses
	NetAfterOwner int64 `json:"net_after_owner"`  // net - transfer_to_owner
}

func (s *SummaryService) ByBranch(ctx context.Context, orgID, branchID string, from, to time.Time) (Summary, error) {
	var res Summary

	err := s.db.QueryRow(ctx, `
		SELECT
			COALESCE(SUM(CASE WHEN type='income' THEN amount END), 0) AS income,
			COALESCE(SUM(CASE WHEN type='expense_company' THEN amount END), 0) AS expense_company,
			COALESCE(SUM(CASE WHEN type='expense_people' THEN amount END), 0) AS expense_people,
			COALESCE(SUM(CASE WHEN type='transfer_to_owner' THEN amount END), 0) AS transfer_to_owner,
			COALESCE(SUM(CASE WHEN type IN ('expense_company','expense_people') THEN amount END), 0) AS total_expenses
		FROM transactions
		WHERE organization_id=$1 AND branch_id=$2 AND occurred_at >= $3 AND occurred_at < $4
  AND is_cancelled = false

	`, orgID, branchID, from, to).Scan(
		&res.Income,
		&res.ExpenseCompany,
		&res.ExpensePeople,
		&res.TransferToOwner,
		&res.TotalExpenses,
	)
	if err != nil {
		return Summary{}, err
	}

	res.Net = res.Income - res.TotalExpenses
	res.NetAfterOwner = res.Net - res.TransferToOwner

	return res, nil
}

// Только owner: общий итог по всем branch (агрегация)
func (s *SummaryService) ByOrgAllBranches(ctx context.Context, orgID string, from, to time.Time) (Summary, error) {
	var res Summary

	err := s.db.QueryRow(ctx, `
		SELECT
			COALESCE(SUM(CASE WHEN type='income' THEN amount END), 0) AS income,
			COALESCE(SUM(CASE WHEN type='expense_company' THEN amount END), 0) AS expense_company,
			COALESCE(SUM(CASE WHEN type='expense_people' THEN amount END), 0) AS expense_people,
			COALESCE(SUM(CASE WHEN type='transfer_to_owner' THEN amount END), 0) AS transfer_to_owner,
			COALESCE(SUM(CASE WHEN type IN ('expense_company','expense_people') THEN amount END), 0) AS total_expenses
		FROM transactions
		WHERE organization_id=$1 AND occurred_at >= $2 AND occurred_at < $3
	`, orgID, from, to).Scan(
		&res.Income,
		&res.ExpenseCompany,
		&res.ExpensePeople,
		&res.TransferToOwner,
		&res.TotalExpenses,
	)
	if err != nil {
		return Summary{}, err
	}

	res.Net = res.Income - res.TotalExpenses
	res.NetAfterOwner = res.Net - res.TransferToOwner

	return res, nil
}

// Для списков транзакций используем типы из repository
type Tx = repository.Transaction
