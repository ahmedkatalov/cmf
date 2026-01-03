package router

import (
	"net/http"

	"backend/internal/middleware"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Dependencies struct {
	JWTSecret string

	AuthHandler   http.Handler
	BranchHandler http.Handler
	UserHandler   http.Handler

	TransactionHandler http.Handler
	SummaryHandler     http.Handler

	MetaHandler http.Handler

	// ✅ ВАЖНО: MeHandler должен быть http.HandlerFunc (не http.Handler)
	MeHandler http.HandlerFunc
}

func New(dep Dependencies) http.Handler {
	r := chi.NewRouter()

	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Route("/api", func(api chi.Router) {

		// ---------- PUBLIC ----------
		api.Mount("/auth", dep.AuthHandler)
		api.Mount("/meta", dep.MetaHandler)

		// ---------- PROTECTED ----------
		api.Group(func(pr chi.Router) {
			pr.Use(middleware.JWT(dep.JWTSecret))

			// ✅ /auth/me
			pr.Get("/auth/me", dep.MeHandler)

			// ✅ branches
			// owner + admin могут заходить (admin может только PATCH своего филиала)
			pr.Route("/branches", func(br chi.Router) {
				br.Use(middleware.RequireRoles("owner", "admin"))
				br.Mount("/", dep.BranchHandler)
			})

			// ✅ users (owner/admin)
			pr.Route("/users", func(u chi.Router) {
				u.Use(middleware.RequireRoles("owner", "admin"))
				u.Mount("/", dep.UserHandler)
			})

			// ✅ transactions
			pr.Route("/transactions", func(tr chi.Router) {
				tr.Use(middleware.RequireRoles(
					"owner", "admin", "manager", "accountant", "security", "employee",
				))
				tr.Mount("/", dep.TransactionHandler)
			})

			// ✅ summary
			pr.Route("/summary", func(sr chi.Router) {
				sr.Use(middleware.RequireRoles("owner", "admin", "manager", "accountant"))
				sr.Mount("/", dep.SummaryHandler)
			})
		})
	})

	return r
}
