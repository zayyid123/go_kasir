package main

import (
	"kasir-api/internal/routes"

	"github.com/gin-gonic/gin"
)

// @title Go Kasir API
// @version 1.0
// @description API untuk sistem kasir
// @BasePath /
func main() {
	r := gin.Default()
	routes.Register(r)
	r.Run(":8080")
}
