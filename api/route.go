package api

import (
	"WorkWorth/config"
	"WorkWorth/internal/auth"
	"WorkWorth/internal/purchases"
	"WorkWorth/internal/users"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, db *sql.DB, cfg *config.Config) {
	userRepo := users.NewUserRepository(db)

	authService := auth.NewAuthService(userRepo, cfg.JWTSecret, cfg.JWTExpiry)

	authHandler := auth.NewAuthHandler(authService)

	publicRouter := router.Group("")
	publicRouter.POST("/signup", authHandler.Signup)
	publicRouter.POST("/login", authHandler.Login)

	purchaseRepo := purchases.NewPurchaseRepository(db)
	purchaseService := purchases.NewPurchaseService(purchaseRepo)
	purchaseHanlder := purchases.NewPurchaseHandler(purchaseService)

	privateRouter := router.Group("/purchases")
	privateRouter.Use(auth.JwtAuthMiddleware(cfg.JWTSecret))

	privateRouter.POST("/create", purchaseHanlder.CreatePurchase)
	privateRouter.GET("/fetch", purchaseHanlder.FetchPurchasesByUserId)
}
