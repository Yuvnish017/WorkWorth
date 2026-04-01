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

	api := router.Group("/api")
	api.Use(auth.JwtAuthMiddleware(cfg.JWTSecret))
	api.POST("/singup", authHandler.Signup)
	api.POST("/login", authHandler.Login)
}
