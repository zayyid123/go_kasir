package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHealth godoc
// @Summary Get health status
// @Description Check API health status
// @Tags Health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "API running"})
}
