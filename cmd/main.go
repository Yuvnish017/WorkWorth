package main

import (
	"WorkWorth/api"
	"WorkWorth/config"
	"WorkWorth/internal/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	db, err := database.NewDB(*cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	router := gin.Default()
	api.SetupRouter(router, db, cfg)
	router.Run(":8080")
}
