package main

import (
	"kasir-api/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.Register(r)
	r.Run(":8080")
}
