package main

import (
	"fmt"
	"log"
	"net/http"

	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/router"
	"backend/internal/service"
)

func main() {
	cfg := config.Load()
	pool := db.New(cfg.DatabaseURL)
	defer pool.Close()

	// Repos
	orgRepo := repository.NewOrgRepo(pool)
	branchRepo := repository.NewBranchRepo(pool)
	userRepo := repository.NewUserRepo(pool)

	// Services
	authService := service.NewAuthService(userRepo, orgRepo, branchRepo, cfg.JWTSecret)
	branchService := service.NewBranchService(branchRepo)
	userService := service.NewUserService(userRepo)

	// Handlers
	authHandler := handler.NewAuthHandler(authService)
	branchHandler := handler.NewBranchHandler(branchService)
	userHandler := handler.NewUserHandler(userService)

	// Router
	r := router.New(router.Dependencies{
		JWTSecret:     cfg.JWTSecret,
		AuthHandler:   authHandler,
		BranchHandler: branchHandler,
		UserHandler:   userHandler,
	})

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Println("API running on", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
