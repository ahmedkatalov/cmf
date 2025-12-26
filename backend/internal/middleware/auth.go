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

			claims := token.Claims.(jwt.MapClaims)

			ctx := r.Context()
			ctx = context.WithValue(ctx, CtxUserID, claims["user_id"].(string))
			ctx = context.WithValue(ctx, CtxOrgID, claims["org_id"].(string))
			ctx = context.WithValue(ctx, CtxRole, claims["role"].(string))

			if v, ok := claims["branch_id"]; ok {
				ctx = context.WithValue(ctx, CtxBranchID, v.(string))
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
