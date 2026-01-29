package main

import (
	"kasir-api/internal/config"
	"kasir-api/internal/database"
	"kasir-api/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

// @title Go Kasir API
// @version 1.0
// @description API untuk sistem kasir
// @BasePath /
func main() {
	config.Load()

	db, err := database.InitDB(config.Cfg.DBHost)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	r := gin.Default()
	routes.Register(r, db)
	r.Run(":" + config.Cfg.AppPort)
}
