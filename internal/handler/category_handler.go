package handler

import (
	"fmt"
	"net/http"

	"kasir-api/internal/model"
	"kasir-api/internal/service"

	"github.com/gin-gonic/gin"

	"kasir-api/internal/utils"
)

// GetCategories godoc
// @Summary Get all categories
// @Description Ambil semua kategori
// @Tags Categories
// @Produce json
// @Success 200 {array} model.Category
// @Router /api/categories [get]
func GetCategories(c *gin.Context) {
	c.JSON(http.StatusOK, service.GetCategories())
}

// GetCategoryByID godoc
// @Summary Get category by ID
// @Description Ambil kategori berdasarkan ID
// @Tags Categories
// @Produce json
// @Param id path uint true "Category ID"
// @Success 200 {object} model.Category
// @Router /api/categories/{id} [get]
func GetCategoryByID(c *gin.Context) {
	idParam := c.Param("id")
	// Convert idParam to uint
	var id uint
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
		return
	}

	// Call service to get category by ID
	category, err := service.GetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Buat kategori baru
// @Tags Categories
// @Accept json
// @Produce json
// @Param category body model.CreateCategoryRequest true "Category to create"
// @Success 201 {object} model.Category
// @Router /api/categories [post]
func CreateCategory(c *gin.Context) {
	var req model.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	category := service.AddCategory(req)
	c.JSON(http.StatusCreated, category)
}

// UpdateCategory godoc
// @Summary Update an existing category
// @Description Perbarui kategori yang ada
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path uint true "Category ID"
// @Param category body model.UpdateCategoryRequest true "Category data to update"
// @Success 200 {object} model.Category
// @Router /api/categories/{id} [put]
func UpdateCategory(c *gin.Context) {
	idParam := c.Param("id")
	var req model.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	// Convert idParam to uint
	var id uint
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
		return
	}

	_, err = service.UpdateCategory(id, req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "category updated successfully"})
}

// DeleteCategory godoc
// @Summary Delete a category
// @Description Hapus kategori
// @Tags Categories
// @Param id path uint true "Category ID"
// @Success 200 {object} map[string]string
// @Router /api/categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
	idParam := c.Param("id")
	// Convert idParam to uint
	var id uint
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
		return
	}

	err = service.DeleteCategory(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "category deleted successfully"})
}
