package routes

import (
	"database/sql"
	"kasir-api/internal/handler"
	"kasir-api/internal/repository"
	"kasir-api/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "kasir-api/docs"
)

func Register(r *gin.Engine, DB *sql.DB) {
	repo := repository.NewProductRepo(DB)
	service := service.NewProductService(repo)
	handlerProduct := handler.NewProductHandler(service)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Go Kasir API",
		})
	})

	// Health check route
	r.GET("/health", handler.GetHealth)

	// Product routes
	r.GET("/products", handlerProduct.GetProducts)
	r.POST("/products", handlerProduct.CreateProduct)
	r.GET("/products/:id", handlerProduct.GetProductByID)
	r.PUT("/products/:id", handlerProduct.UpdateProduct)
	r.DELETE("/products/:id", handlerProduct.DeleteProduct)

	// Category routes
	r.GET("/categories", handler.GetCategories)
	r.POST("/categories", handler.CreateCategory)
	r.GET("/categories/:id", handler.GetCategoryByID)
	r.PUT("/categories/:id", handler.UpdateCategory)
	r.DELETE("/categories/:id", handler.DeleteCategory)
}
