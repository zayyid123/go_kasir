package handler

import (
	"kasir-api/internal/model"
	"kasir-api/internal/service"
	"kasir-api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(s *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: s}
}

// Checkout godoc
// @Summary Checkout product
// @Description Buat Transaksi
// @Tags Transaction
// @Accept json
// @Produce json
// @Param transaction body model.CheckoutRequest true "Transaction to create"
// @Success 201 {object} model.Transaction
// @Router /transaction/checkout [post]
func (h *TransactionHandler) Checkout(c *gin.Context) {
	var req model.CheckoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	transaction, err := h.service.Checkout(req.Items)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, transaction)
}
