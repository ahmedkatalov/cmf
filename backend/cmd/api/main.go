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

	// ---------- Repos ----------
	orgRepo := repository.NewOrgRepo(pool)
	branchRepo := repository.NewBranchRepo(pool)
	userRepo := repository.NewUserRepo(pool)
	txRepo := repository.NewTransactionRepo(pool)

	// ---------- Services ----------
	authService := service.NewAuthService(userRepo, orgRepo, branchRepo, cfg.JWTSecret)
	branchService := service.NewBranchService(branchRepo, userRepo)
	userService := service.NewUserService(userRepo)
	txService := service.NewTransactionService(txRepo)
	summaryService := service.NewSummaryService(pool)

	// ---------- Handlers ----------
	authHandler := handler.NewAuthHandler(authService)
	branchHandler := handler.NewBranchHandler(branchService)
	userHandler := handler.NewUserHandler(userService)
	txHandler := handler.NewTransactionHandler(txService)
	summaryHandler := handler.NewSummaryHandler(summaryService)

	// meta + me
	metaHandler := handler.NewMetaHandler()
	meHandler := handler.NewMeHandler()

	// ---------- Router ----------
	r := router.New(router.Dependencies{
		JWTSecret: cfg.JWTSecret,

		AuthHandler:   authHandler,
		BranchHandler: branchHandler,
		UserHandler:   userHandler,

		TransactionHandler: txHandler,
		SummaryHandler:     summaryHandler,

		MetaHandler: metaHandler,
		MeHandler:   meHandler,
	})

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Println("API running on", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
