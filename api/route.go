package api

import (
	"WorkWorth/config"
	"WorkWorth/internal/auth"
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

	privateRouter := router.Group("")
	privateRouter.Use(auth.JwtAuthMiddleware(cfg.JWTSecret))
	privateRouter.POST("/login", authHandler.Login)
}
