package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"backend/internal/middleware"
	"backend/internal/service"

	"github.com/go-chi/chi/v5"
)

type SummaryHandler struct {
	summary *service.SummaryService
}

func NewSummaryHandler(summary *service.SummaryService) http.Handler {
	h := &SummaryHandler{summary: summary}
	r := chi.NewRouter()

	// summary по текущей точке (или owner может выбрать branch)
	r.Get("/", h.byBranch)

	// summary по всем точкам (только owner)
	r.Get("/all", h.all)

	return r
}

func monthRange(month string) (time.Time, time.Time, error) {
	// month = YYYY-MM
	start, err := time.Parse("2006-01", month)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	end := start.AddDate(0, 1, 0)
	return start, end, nil
}

func (h *SummaryHandler) byBranch(w http.ResponseWriter, r *http.Request) {
	orgID := r.Context().Value(middleware.CtxOrgID).(string)
	role := r.Context().Value(middleware.CtxRole).(string)
	branchCtx, _ := r.Context().Value(middleware.CtxBranchID).(string)

	month := r.URL.Query().Get("month")
	if month == "" {
		http.Error(w, "month query required (YYYY-MM)", http.StatusBadRequest)
		return
	}

	from, to, err := monthRange(month)
	if err != nil {
		http.Error(w, "invalid month (YYYY-MM)", http.StatusBadRequest)
		return
	}

	branchID := branchCtx
	if role == "owner" {
		if b := r.URL.Query().Get("branch_id"); b != "" {
			branchID = b
		}
	}

	sum, err := h.summary.ByBranch(r.Context(), orgID, branchID, from, to)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(sum)
}

func (h *SummaryHandler) all(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value(middleware.CtxRole).(string)
	if role != "owner" {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	orgID := r.Context().Value(middleware.CtxOrgID).(string)

	month := r.URL.Query().Get("month")
	if month == "" {
		http.Error(w, "month query required (YYYY-MM)", http.StatusBadRequest)
		return
	}

	from, to, err := monthRange(month)
	if err != nil {
		http.Error(w, "invalid month (YYYY-MM)", http.StatusBadRequest)
		return
	}

	sum, err := h.summary.ByOrgAllBranches(r.Context(), orgID, from, to)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(sum)
}
