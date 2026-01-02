package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type ctxKey string

const (
	CtxUserID   ctxKey = "user_id"
	CtxEmail    ctxKey = "email"
	CtxOrgID    ctxKey = "org_id"
	CtxBranchID ctxKey = "branch_id"
	CtxRole     ctxKey = "role"
)

func JWT(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "missing token", http.StatusUnauthorized)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			tokenStr := parts[1]
			token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})
			if err != nil || !token.Valid {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "invalid token claims", http.StatusUnauthorized)
				return
			}

			// ✅ безопасные чтения (без паник)
			userID, _ := claims["user_id"].(string)
			orgID, _ := claims["org_id"].(string)
			role, _ := claims["role"].(string)
			email, _ := claims["email"].(string)

			if userID == "" || orgID == "" || role == "" {
				http.Error(w, "invalid token payload", http.StatusUnauthorized)
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, CtxUserID, userID)
			ctx = context.WithValue(ctx, CtxOrgID, orgID)
			ctx = context.WithValue(ctx, CtxRole, role)

			// ✅ добавили email в контекст
			if email != "" {
				ctx = context.WithValue(ctx, CtxEmail, email)
			}

			// ✅ branch_id опциональный
			if v, ok := claims["branch_id"]; ok {
				if branchID, ok2 := v.(string); ok2 && branchID != "" {
					ctx = context.WithValue(ctx, CtxBranchID, branchID)
				}
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
