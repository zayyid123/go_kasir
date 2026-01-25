package routes

import (
	"kasir-api/internal/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "kasir-api/docs"
)

func Register(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Go Kasir API",
		})
	})

	// Health check route
	r.GET("/health", handler.GetHealth)

	// Product routes
	r.GET("/products", handler.GetProducts)
	r.POST("/products", handler.CreateProduct)
	r.GET("/products/:id", handler.GetProductByID)
	r.PUT("/products/:id", handler.UpdateProduct)
	r.DELETE("/products/:id", handler.DeleteProduct)

	// Category routes
	r.GET("/categories", handler.GetCategories)
	r.POST("/categories", handler.CreateCategory)
	r.GET("/categories/:id", handler.GetCategoryByID)
	r.PUT("/categories/:id", handler.UpdateCategory)
	r.DELETE("/categories/:id", handler.DeleteCategory)
}
