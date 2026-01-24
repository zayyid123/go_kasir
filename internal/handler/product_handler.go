package handler

import (
	"fmt"
	"net/http"

	"kasir-api/internal/model"
	"kasir-api/internal/service"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, service.GetProducts())
}

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

func CreateProduct(c *gin.Context) {
	var req model.Product
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := service.AddProduct(req)
	c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	var req model.Product
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
