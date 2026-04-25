package api

import (
	"WorkWorth/config"
	"WorkWorth/internal/auth"
	"WorkWorth/internal/calculator"
	"WorkWorth/internal/currency"
	"WorkWorth/internal/purchases"
	"WorkWorth/internal/users"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, db *sql.DB, cfg *config.Config) {
	publicRouter := router.Group("")

	privateRouter := router.Group("/purchases")
	privateRouter.Use(auth.JwtAuthMiddleware(cfg.JWTSecret))

	userRepo := users.NewUserRepository(db)

	authService := auth.NewAuthService(userRepo, cfg.JWTSecret, cfg.JWTExpiry)
	authHandler := auth.NewAuthHandler(authService)

	publicRouter.POST("/signup", authHandler.Signup)
	publicRouter.POST("/login", authHandler.Login)

	currencyRepo := currency.NewCurrencyRepository(db)
	currencyService := currency.NewCurrencyService(currencyRepo)
	conversionHandler := currency.NewExchangeProvider(currencyService)

	publicRouter.GET("/convert", conversionHandler.ConvertCurrency)

	purchaseRepo := purchases.NewPurchaseRepository(db)
	purchaseService := purchases.NewPurchaseService(purchaseRepo, currencyService)
	purchaseHanlder := purchases.NewPurchaseHandler(purchaseService)

	privateRouter.POST("/create", purchaseHanlder.CreatePurchase)
	privateRouter.GET("/fetch", purchaseHanlder.FetchPurchasesByUserId)

	calculatorService := calculator.NewCalculatorService(userRepo, *currencyService)
	calculatorHandler := calculator.NewCalculatorHandler(calculatorService)

	privateRouter.GET("suggest", calculatorHandler.Suggestor)
}
