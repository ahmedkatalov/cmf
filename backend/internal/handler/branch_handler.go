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

	// owner creates branch
	r.Post("/", h.create)

	// list branches
	r.Get("/", h.list)

	// get branch by id
	r.Get("/{id}", h.getByID)

	// list users of branch
	r.Get("/{id}/users", h.usersByBranch)

	// update branch
	r.Patch("/{id}", h.update)

	// delete branch
	r.Delete("/{id}", h.delete)

	return r
}

type createBranchRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (h *BranchHandler) create(w http.ResponseWriter, r *http.Request) {
	orgID := r.Context().Value(middleware.CtxOrgID).(string)

	// ✅ Только owner может создавать филиалы
	role := r.Context().Value(middleware.CtxRole).(string)
	if role != "owner" {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

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

func (h *BranchHandler) getByID(w http.ResponseWriter, r *http.Request) {
	orgID := r.Context().Value(middleware.CtxOrgID).(string)
	branchID := chi.URLParam(r, "id")

	b, err := h.branches.GetByID(r.Context(), orgID, branchID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(b)
}

func (h *BranchHandler) usersByBranch(w http.ResponseWriter, r *http.Request) {
	orgID := r.Context().Value(middleware.CtxOrgID).(string)
	branchID := chi.URLParam(r, "id")

	users, err := h.branches.ListUsers(r.Context(), orgID, branchID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(users)
}

type updateBranchRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (h *BranchHandler) update(w http.ResponseWriter, r *http.Request) {
	orgID := r.Context().Value(middleware.CtxOrgID).(string)
	role := r.Context().Value(middleware.CtxRole).(string)
	branchCtx, _ := r.Context().Value(middleware.CtxBranchID).(string)

	branchID := chi.URLParam(r, "id")

	// ✅ admin может менять только свой филиал
	if role == "admin" && branchID != branchCtx {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	var req updateBranchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if err := h.branches.Update(r.Context(), orgID, branchID, req.Name, req.Address); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *BranchHandler) delete(w http.ResponseWriter, r *http.Request) {
	orgID := r.Context().Value(middleware.CtxOrgID).(string)

	// ✅ Только owner может удалять филиалы
	role := r.Context().Value(middleware.CtxRole).(string)
	if role != "owner" {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	branchID := chi.URLParam(r, "id")

	if err := h.branches.Delete(r.Context(), orgID, branchID); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
