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

	// ✅ CORS (для фронта)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Public routes
	r.Route("/api", func(api chi.Router) {
		api.Mount("/auth", dep.AuthHandler)
	})

	// Protected routes
	r.Route("/api", func(api chi.Router) {
		api.Use(middleware.JWT(dep.JWTSecret))

		// ✅ Точки: owner/admin
api.Route("/branches", func(br chi.Router) {
    br.Use(middleware.RequireRoles("owner"))
    br.Mount("/", dep.BranchHandler)
})

		// ✅ Пользователи:
		// owner/admin создают сотрудников в любых точках
		// manager создаёт в своей точке
api.Route("/users", func(u chi.Router) {
    u.Use(middleware.RequireRoles("owner", "admin")) 
    u.Mount("/", dep.UserHandler)
})

	})

	return r
}
