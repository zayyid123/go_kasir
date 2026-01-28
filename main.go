package main

import (
	"kasir-api/internal/config"
	"kasir-api/internal/routes"

	"github.com/gin-gonic/gin"
)

// @title Go Kasir API
// @version 1.0
// @description API untuk sistem kasir
// @BasePath /
func main() {
	config.Load()
	r := gin.Default()
	routes.Register(r)
	r.Run(":" + config.Cfg.AppPort)
}
