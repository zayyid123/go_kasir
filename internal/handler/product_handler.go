package handler

import (
	"fmt"
	"net/http"

	"kasir-api/internal/model"
	"kasir-api/internal/service"

	"github.com/gin-gonic/gin"

	"kasir-api/internal/utils"
)

// GetProducts godoc
// @Summary Get all products
// @Description Ambil semua produk
// @Tags Products
// @Produce json
// @Success 200 {array} model.Product
// @Router /api/products [get]
func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, service.GetProducts())
}

// GetProductByID godoc
// @Summary Get product by ID
// @Description Ambil produk berdasarkan ID
// @Tags Products
// @Produce json
// @Param id path uint true "Product ID"
// @Success 200 {object} model.Product
// @Router /api/products/{id} [get]
func GetProductByID(c *gin.Context) {
	idParam := c.Param("id")
	// Convert idParam to uint
	var id uint
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	// Call service to get product by ID
	product, err := service.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Buat produk baru
// @Tags Products
// @Accept json
// @Produce json
// @Param product body model.CreateProductRequest true "Product to create"
// @Success 201 {object} model.Product
// @Router /api/products [post]
func CreateProduct(c *gin.Context) {
	var req model.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	product := service.AddProduct(req)
	c.JSON(http.StatusCreated, product)
}

// UpdateProduct godoc
// @Summary Update an existing product
// @Description Perbarui produk yang ada
// @Tags Products
// @Accept json
// @Produce json
// @Param id path uint true "Product ID"
// @Param product body model.UpdateProductRequest true "Product data to update"
// @Success 200 {object} model.Product
// @Router /api/products/{id} [put]
func UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	var req model.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	// Convert idParam to uint
	var id uint
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	_, err = service.UpdateProduct(id, req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product updated successfully"})
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Hapus produk
// @Tags Products
// @Param id path uint true "Product ID"
// @Success 200 {object} map[string]string
// @Router /api/products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	// Convert idParam to uint
	var id uint
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	err = service.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product deleted successfully"})
}
