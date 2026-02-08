package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "kasir-api/docs"
	"kasir-api/internal/handler"
	"kasir-api/internal/repository"
	"kasir-api/internal/service"
)

func Register(r *gin.Engine, DB *sql.DB) {
	// Initialize repositories
	repoProduct := repository.NewProductRepo(DB)
	repoCategory := repository.NewCategoryRepo(DB)
	repoTransaction := repository.NewTransactionRepo(DB)
	repoReport := repository.NewReportRepo(DB)

	// Initialize services
	serviceProduct := service.NewProductService(repoProduct)
	serviceCategory := service.NewCategoryService(repoCategory)
	serviceTransaction := service.NewTransactionService(repoTransaction)
	serviceReport := service.NewReportService(repoReport)

	// Initialize handlers
	handlerProduct := handler.NewProductHandler(serviceProduct)
	handlerCategory := handler.NewCategoryHandler(serviceCategory)
	handlerTransaction := handler.NewTransactionHandler(serviceTransaction)
	handlerReport := handler.NewReportHandler(serviceReport)

	// Setup routes
	setupSwaggerRoutes(r)
	setupGeneralRoutes(r)
	setupProductRoutes(r, handlerProduct)
	setupCategoryRoutes(r, handlerCategory)
	setupTransactionRoutes(r, handlerTransaction)
	setupReportRoutes(r, handlerReport)
}

func setupSwaggerRoutes(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func setupGeneralRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Go Kasir API",
		})
	})
	r.GET("/health", handler.GetHealth)
}

func setupProductRoutes(r *gin.Engine, h *handler.ProductHandler) {
	products := r.Group("/products")
	{
		products.GET("", h.GetProducts)
		products.POST("", h.CreateProduct)
		products.GET("/:id", h.GetProductByID)
		products.PUT("/:id", h.UpdateProduct)
		products.DELETE("/:id", h.DeleteProduct)
	}
}

func setupCategoryRoutes(r *gin.Engine, h *handler.CategoryHandler) {
	categories := r.Group("/categories")
	{
		categories.GET("", h.GetCategories)
		categories.POST("", h.CreateCategory)
		categories.GET("/:id", h.GetCategoryByID)
		categories.PUT("/:id", h.UpdateCategory)
		categories.DELETE("/:id", h.DeleteCategory)
	}
}

func setupTransactionRoutes(r *gin.Engine, h *handler.TransactionHandler) {
	transaction := r.Group("/transaction")
	{
		transaction.POST("/checkout", h.Checkout)
	}
}

func setupReportRoutes(r *gin.Engine, h *handler.ReportHandler) {
	report := r.Group("/report")
	{
		report.GET("", h.Report)
	}
}
