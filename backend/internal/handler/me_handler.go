package handler

import (
	"backend/internal/middleware"
	"encoding/json"
	"net/http"
)

func NewMeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(middleware.CtxUserID).(string)
		orgID := r.Context().Value(middleware.CtxOrgID).(string)
		role := r.Context().Value(middleware.CtxRole).(string)

		branchID, _ := r.Context().Value(middleware.CtxBranchID).(string)
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
}
