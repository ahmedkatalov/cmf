package router

import (
	"backend/internal/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Dependencies struct {
	JWTSecret     string
	AuthHandler   http.Handler
	BranchHandler http.Handler
	UserHandler   http.Handler
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

		// ---------- PROTECTED ----------
		api.Group(func(pr chi.Router) {
			pr.Use(middleware.JWT(dep.JWTSecret))

			// branches (owner)
			pr.Route("/branches", func(br chi.Router) {
				br.Use(middleware.RequireRoles("owner"))
				br.Mount("/", dep.BranchHandler)
			})

			// users (owner/admin)
			pr.Route("/users", func(u chi.Router) {
				u.Use(middleware.RequireRoles("owner", "admin"))
				u.Mount("/", dep.UserHandler)
			})
		})
	})

	return r
}

