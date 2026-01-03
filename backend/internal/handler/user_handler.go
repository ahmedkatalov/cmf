package handler

import (
	"encoding/json"
	"net/http"

	"backend/internal/middleware"
	"backend/internal/service"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	users *service.UserService
}

func NewUserHandler(users *service.UserService) http.Handler {
	h := &UserHandler{users: users}
	r := chi.NewRouter()

	// ✅ Создать сотрудника (owner/admin)
	r.Post("/", h.create)

	return r
}

type createUserRequest struct {
	BranchID string `json:"branch_id"` // ✅ owner обязан указать, admin игнорирует
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	FullName string `json:"full_name"`
Phone    string `json:"phone"`


}

var allowedRoles = map[string]bool{
	"admin":      true,
	"manager":    true,
	"accountant": true,
	"security":   true,
	"employee":   true,
}

func (h *UserHandler) create(w http.ResponseWriter, r *http.Request) {
	orgID := r.Context().Value(middleware.CtxOrgID).(string)
	requesterRole := r.Context().Value(middleware.CtxRole).(string)

	branchCtx, _ := r.Context().Value(middleware.CtxBranchID).(string)

	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// ✅ Проверяем что роль корректная
	if !allowedRoles[req.Role] {
		http.Error(w, "invalid role", http.StatusBadRequest)
		return
	}

	// ✅ Запрещаем создание owner всегда
	if req.Role == "owner" {
		http.Error(w, "cannot create owner", http.StatusBadRequest)
		return
	}

	// ✅ ПРАВИЛА ПО ТВОЕМУ ТЗ:
	// OWNER -> может создать в любой точке (branch_id обязателен)
	// ADMIN -> может создать только в своей точке (branch_id игнорируем)
	var branchID *string

	if requesterRole == "owner" {
		if req.BranchID == "" {
			http.Error(w, "branch_id required", http.StatusBadRequest)
			return
		}
		branchID = &req.BranchID

		// owner может создать даже admin (если нужно)
		// но ты можешь запретить и owner тоже - скажи если надо
	} else if requesterRole == "admin" {
		// admin создаёт только сотрудников своей точки
		branchID = &branchCtx

		// ✅ admin НЕ может создавать admin
		if req.Role == "admin" {
			http.Error(w, "admin cannot create admin", http.StatusForbidden)
			return
		}
	} else {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

id, err := h.users.Create(r.Context(), orgID, branchID, req.FullName, req.Phone, req.Email, req.Password, req.Role)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"id": id})
}
