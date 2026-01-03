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

	// ✅ Create branch
	r.Post("/", h.create)

	// ✅ List branches
	r.Get("/", h.list)

<<<<<<< HEAD
	// ✅ Get branch by id
	r.Get("/{id}", h.getByID)

	// ✅ Get users by branch id
	r.Get("/{id}/users", h.usersByBranch)

=======
	// ✅ Получить точку по ID
	r.Get("/{id}", h.getByID)

>>>>>>> 512879ce66463258d4ab363f62b7e9f08b04d422
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
