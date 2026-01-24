package routes

import (
	"kasir-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	api := r.Group("/api")

	api.GET("/products", handler.GetProducts)
	api.POST("/products", handler.CreateProduct)
	api.GET("/products/:id", handler.GetProductByID)
	api.PUT("/products/:id", handler.UpdateProduct)
	api.DELETE("/products/:id", handler.DeleteProduct)
}
