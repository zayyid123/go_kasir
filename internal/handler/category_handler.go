package handler

import (
	"fmt"
	"net/http"

	"kasir-api/internal/model"
	"kasir-api/internal/service"

	"github.com/gin-gonic/gin"

	"kasir-api/internal/utils"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(s *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: s}
}

// GetCategories godoc
// @Summary Get all categories
// @Description Ambil semua kategori
// @Tags Categories
// @Produce json
// @Success 200 {array} model.Category
// @Router /categories [get]
func (h *CategoryHandler) GetCategories(c *gin.Context) {
	data, err := h.service.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, data)
}

// GetCategoryByID godoc
// @Summary Get category by ID
// @Description Ambil kategori berdasarkan ID
// @Tags Categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} model.Category
// @Router /categories/{id} [get]
func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	idParam := c.Param("id")
	// Convert idParam to int
	var id int
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
		return
	}

	// Call service to get category by ID
	data, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Buat kategori baru
// @Tags Categories
// @Accept json
// @Produce json
// @Param category body model.CreateCategoryRequest true "Category to create"
// @Success 201 {object} model.Category
// @Router /categories [post]
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req model.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	category, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, category)
}

// UpdateCategory godoc
// @Summary Update an existing category
// @Description Perbarui kategori yang ada
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body model.UpdateCategoryRequest true "Category data to update"
// @Success 200 {object} model.Category
// @Router /categories/{id} [put]
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	idParam := c.Param("id")
	var req model.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	// Convert idParam to int
	var id int
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
		return
	}

	category, err := h.service.Update(id, req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

// DeleteCategory godoc
// @Summary Delete a category
// @Description Hapus kategori
// @Tags Categories
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]string
// @Router /categories/{id} [delete]
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	idParam := c.Param("id")
	// Convert idParam to int
	var id int
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "category deleted successfully"})
}
