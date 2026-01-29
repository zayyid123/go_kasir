package handler

import (
	"fmt"
	"net/http"

	"kasir-api/internal/model"
	"kasir-api/internal/service"

	"github.com/gin-gonic/gin"

	"kasir-api/internal/utils"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(s *service.ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

// GetProducts godoc
// @Summary Get all products
// @Description Ambil semua produk
// @Tags Products
// @Produce json
// @Success 200 {array} model.Product
// @Router /products [get]
func (h *ProductHandler) GetProducts(c *gin.Context) {
	data, err := h.service.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, data)
}

// GetProductByID godoc
// @Summary Get product by ID
// @Description Ambil produk berdasarkan ID
// @Tags Products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} model.Product
// @Router /products/{id} [get]
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	idParam := c.Param("id")
	// Convert idParam to int
	var id int
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	// Call service to get product by ID
	data, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Buat produk baru
// @Tags Products
// @Accept json
// @Produce json
// @Param product body model.CreateProductRequest true "Product to create"
// @Success 201 {object} model.Product
// @Router /products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req model.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	product, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
}

// UpdateProduct godoc
// @Summary Update an existing product
// @Description Perbarui produk yang ada
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body model.UpdateProductRequest true "Product data to update"
// @Success 200 {object} model.Product
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	var req model.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	// Convert idParam to int
	var id int
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	product, err := h.service.Update(id, req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Hapus produk
// @Tags Products
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]string
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	// Convert idParam to int
	var id int
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product deleted successfully"})
}
