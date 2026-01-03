package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Transaction struct {
	ID             string     `json:"id"`
	OrganizationID string     `json:"organization_id"`
	BranchID       string     `json:"branch_id"`
	CreatedBy      *string    `json:"created_by"`
	Type           string     `json:"type"`
	CategoryID     *string    `json:"category_id"`
	Amount         int64      `json:"amount"`
	OccurredAt     time.Time  `json:"occurred_at"`
	Description    *string    `json:"description"`
	CreatedAt      time.Time  `json:"created_at"`
	IsCancelled    bool       `json:"is_cancelled"`
	CancelledAt    *time.Time `json:"cancelled_at,omitempty"`
	CancelledBy    *string    `json:"cancelled_by,omitempty"`
	CancelReason   *string    `json:"cancel_reason,omitempty"`
}

type TransactionRepo struct {
	db *pgxpool.Pool
}

func NewTransactionRepo(db *pgxpool.Pool) *TransactionRepo {
	return &TransactionRepo{db: db}
}

func (r *TransactionRepo) Create(
	ctx context.Context,
	orgID, branchID string,
	createdBy *string,
	txType string,
	categoryID *string,
	amount int64,
	occurredAt time.Time,
	description *string,
) (string, error) {
	var id string
	err := r.db.QueryRow(ctx, `
		INSERT INTO transactions (
			organization_id, branch_id, created_by,
			type, category_id, amount, occurred_at, description
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
		RETURNING id
	`, orgID, branchID, createdBy, txType, categoryID, amount, occurredAt, description).Scan(&id)

	return id, err
}

type ListTransactionsFilter struct {
	Type       *string
	CategoryID *string
	DateFrom   *time.Time
	DateTo     *time.Time
}

func (r *TransactionRepo) ListByBranch(ctx context.Context, orgID, branchID string, f ListTransactionsFilter) ([]Transaction, error) {
	query := `
		SELECT
			id, organization_id, branch_id, created_by,
			type, category_id, amount, occurred_at, description,
			created_at,
			is_cancelled, cancelled_at, cancelled_by, cancel_reason
		FROM transactions
		WHERE organization_id = $1
		  AND branch_id = $2
		  AND is_cancelled = false
	`
	args := []any{orgID, branchID}
	i := 3

	if f.Type != nil {
		query += " AND type = $" + itoa(i)
		args = append(args, *f.Type)
		i++
	}
	if f.CategoryID != nil {
		query += " AND category_id = $" + itoa(i)
		args = append(args, *f.CategoryID)
		i++
	}
	if f.DateFrom != nil {
		query += " AND occurred_at >= $" + itoa(i)
		args = append(args, *f.DateFrom)
		i++
	}
	if f.DateTo != nil {
		query += " AND occurred_at <= $" + itoa(i)
		args = append(args, *f.DateTo)
		i++
	}

	query += " ORDER BY occurred_at DESC, created_at DESC"

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []Transaction
	for rows.Next() {
		var t Transaction
		if err := rows.Scan(
			&t.ID, &t.OrganizationID, &t.BranchID, &t.CreatedBy,
			&t.Type, &t.CategoryID, &t.Amount, &t.OccurredAt, &t.Description,
			&t.CreatedAt,
			&t.IsCancelled, &t.CancelledAt, &t.CancelledBy, &t.CancelReason,
		); err != nil {
			return nil, err
		}
		out = append(out, t)
	}

	return out, rows.Err()
}

// ✅ Нужно для отмены: получить транзакцию по id
func (r *TransactionRepo) GetByID(ctx context.Context, orgID, txID string) (*Transaction, error) {
	var t Transaction
	err := r.db.QueryRow(ctx, `
		SELECT
			id, organization_id, branch_id, created_by,
			type, category_id, amount, occurred_at, description,
			created_at,
			is_cancelled, cancelled_at, cancelled_by, cancel_reason
		FROM transactions
		WHERE organization_id=$1 AND id=$2
	`, orgID, txID).Scan(
		&t.ID, &t.OrganizationID, &t.BranchID, &t.CreatedBy,
		&t.Type, &t.CategoryID, &t.Amount, &t.OccurredAt, &t.Description,
		&t.CreatedAt,
		&t.IsCancelled, &t.CancelledAt, &t.CancelledBy, &t.CancelReason,
	)

	if err != nil {
		return nil, err
	}

	return &t, nil
}

// ✅ Отмена транзакции (не удаление)
func (r *TransactionRepo) Cancel(ctx context.Context, orgID, txID, cancelledBy string, reason *string) error {
	_, err := r.db.Exec(ctx, `
		UPDATE transactions
		SET is_cancelled=true,
		    cancelled_at=NOW(),
		    cancelled_by=$3,
		    cancel_reason=$4
		WHERE organization_id=$1
		  AND id=$2
		  AND is_cancelled=false
	`, orgID, txID, cancelledBy, reason)

	return err
}

// ✅ Нужно для удаления филиала: узнать есть ли транзакции
func (r *TransactionRepo) CountByBranch(ctx context.Context, orgID, branchID string) (int, error) {
	var count int
	err := r.db.QueryRow(ctx, `
		SELECT COUNT(*)
		FROM transactions
		WHERE organization_id=$1 AND branch_id=$2 AND is_cancelled=false
	`, orgID, branchID).Scan(&count)
	return count, err
}

// helper itoa
func itoa(i int) string {
	return fmt.Sprintf("%d", i)
}
