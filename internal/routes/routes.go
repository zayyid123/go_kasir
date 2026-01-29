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
	repoProduct := repository.NewProductRepo(DB)
	serviceProduct := service.NewProductService(repoProduct)
	handlerProduct := handler.NewProductHandler(serviceProduct)

	repoCategory := repository.NewCategoryRepo(DB)
	serviceCategory := service.NewCategoryService(repoCategory)
	handlerCategory := handler.NewCategoryHandler(serviceCategory)

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
	r.GET("/categories", handlerCategory.GetCategories)
	r.POST("/categories", handlerCategory.CreateCategory)
	r.GET("/categories/:id", handlerCategory.GetCategoryByID)
	r.PUT("/categories/:id", handlerCategory.UpdateCategory)
	r.DELETE("/categories/:id", handlerCategory.DeleteCategory)
}
