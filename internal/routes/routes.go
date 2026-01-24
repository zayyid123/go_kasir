package routes

import (
	"kasir-api/internal/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "kasir-api/docs"
)

func Register(r *gin.Engine) {
	api := r.Group("/api")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Go Kasir API",
		})
	})

	// Product routes
	api.GET("/products", handler.GetProducts)
	api.POST("/products", handler.CreateProduct)
	api.GET("/products/:id", handler.GetProductByID)
	api.PUT("/products/:id", handler.UpdateProduct)
	api.DELETE("/products/:id", handler.DeleteProduct)

	// Category routes
	api.GET("/categories", handler.GetCategories)
	api.POST("/categories", handler.CreateCategory)
	api.GET("/categories/:id", handler.GetCategoryByID)
	api.PUT("/categories/:id", handler.UpdateCategory)
	api.DELETE("/categories/:id", handler.DeleteCategory)
}
