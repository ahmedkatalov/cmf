package handler

import (
	"backend/internal/middleware"
	"backend/internal/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AuthHandler struct {
	auth *service.AuthService
}

func NewAuthHandler(auth *service.AuthService) http.Handler {
	h := &AuthHandler{auth: auth}
	r := chi.NewRouter()

	// ✅ PUBLIC
	r.Post("/register-root", h.registerRoot)
	r.Post("/login", h.login)

	return r
}

// ✅ PROTECTED HANDLER (будем монтировать отдельно)
func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.CtxUserID).(string)
	orgID := r.Context().Value(middleware.CtxOrgID).(string)
	role := r.Context().Value(middleware.CtxRole).(string)

	branchID, _ := r.Context().Value(middleware.CtxBranchID).(string)

	// email мы кладём в токен, но не сохраняем в контекст.
	// Чтобы достать email, лучше либо:
	// 1) положить его в контекст в middleware.JWT
	// 2) либо вернуть без email
	// Я сделаю вариант 1 ниже (в middleware.JWT добавим CtxEmail)

	email, _ := r.Context().Value(middleware.CtxEmail).(string)

	resp := map[string]any{
		"user_id": userID,
		"org_id":  orgID,
		"role":    role,
		"email":   email,
	}

	if branchID != "" {
		resp["branch_id"] = branchID
	}

	json.NewEncoder(w).Encode(resp)
}

type registerRootRequest struct {
	OrganizationName string `json:"organization_name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
}

func (h *AuthHandler) registerRoot(w http.ResponseWriter, r *http.Request) {
	var req registerRootRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	token, err := h.auth.RegisterRoot(r.Context(), req.OrganizationName, req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	token, err := h.auth.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}