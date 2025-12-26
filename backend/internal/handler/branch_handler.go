package handler

import (
	"backend/internal/middleware"
	"backend/internal/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type BranchHandler struct {
	branches *service.BranchService
}

func NewBranchHandler(branches *service.BranchService) http.Handler {
	h := &BranchHandler{branches: branches}
	r := chi.NewRouter()

	// ✅ Создать точку (owner/admin)
	r.Post("/", h.create)

	// ✅ Список точек (owner/admin)
	r.Get("/", h.list)

	return r
}

type createBranchRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (h *BranchHandler) create(w http.ResponseWriter, r *http.Request) {
	orgID := r.Context().Value(middleware.CtxOrgID).(string)

	var req createBranchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	id, err := h.branches.Create(r.Context(), orgID, req.Name, req.Address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func (h *BranchHandler) list(w http.ResponseWriter, r *http.Request) {
	orgID := r.Context().Value(middleware.CtxOrgID).(string)

	list, err := h.branches.List(r.Context(), orgID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(list)
}
