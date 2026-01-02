package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"backend/internal/middleware"
	"backend/internal/repository"
	"backend/internal/service"

	"github.com/go-chi/chi/v5"
)

type TransactionHandler struct {
	tx *service.TransactionService
}

func NewTransactionHandler(tx *service.TransactionService) http.Handler {
	h := &TransactionHandler{tx: tx}
	r := chi.NewRouter()

	r.Post("/", h.create)
	r.Get("/", h.list)

	return r
}

type createTransactionRequest struct {
	// owner может передать branch_id, остальные игнорируются
	BranchID    string  `json:"branch_id"`
	Type        string  `json:"type"` // income | expense_company | expense_people
	CategoryID  *string `json:"category_id"`
	Amount      int64   `json:"amount"`
	OccurredAt  string  `json:"occurred_at"` // YYYY-MM-DD
	Description *string `json:"description"`
}

func (h *TransactionHandler) create(w http.ResponseWriter, r *http.Request) {
	orgID := r.Context().Value(middleware.CtxOrgID).(string)
	role := r.Context().Value(middleware.CtxRole).(string)
	userID := r.Context().Value(middleware.CtxUserID).(string)
	branchCtx, _ := r.Context().Value(middleware.CtxBranchID).(string)

	var req createTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// parse date
	date, err := time.Parse("2006-01-02", req.OccurredAt)
	if err != nil {
		http.Error(w, "invalid occurred_at (YYYY-MM-DD)", http.StatusBadRequest)
		return
	}

	// branch logic:
	// owner -> может выбрать branch_id
	// остальные -> только своя точка
	branchID := branchCtx
	if role == "owner" {
		if req.BranchID == "" {
			http.Error(w, "branch_id required for owner", http.StatusBadRequest)
			return
		}
		branchID = req.BranchID
	}

	id, err := h.tx.Create(r.Context(), orgID, branchID, &userID, req.Type, req.CategoryID, req.Amount, date, req.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func (h *TransactionHandler) list(w http.ResponseWriter, r *http.Request) {
	orgID := r.Context().Value(middleware.CtxOrgID).(string)
	role := r.Context().Value(middleware.CtxRole).(string)
	branchCtx, _ := r.Context().Value(middleware.CtxBranchID).(string)

	branchID := branchCtx

	// owner может выбрать branch_id через query
	if role == "owner" {
		if b := r.URL.Query().Get("branch_id"); b != "" {
			branchID = b
		}
	}

	var f repository.ListTransactionsFilter

	if t := r.URL.Query().Get("type"); t != "" {
		f.Type = &t
	}
	if c := r.URL.Query().Get("category_id"); c != "" {
		f.CategoryID = &c
	}
	if from := r.URL.Query().Get("from"); from != "" {
		d, err := time.Parse("2006-01-02", from)
		if err == nil {
			f.DateFrom = &d
		}
	}
	if to := r.URL.Query().Get("to"); to != "" {
		d, err := time.Parse("2006-01-02", to)
		if err == nil {
			f.DateTo = &d
		}
	}

	list, err := h.tx.ListByBranch(r.Context(), orgID, branchID, f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(list)
}
