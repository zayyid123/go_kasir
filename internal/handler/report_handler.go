package handler

import (
	"kasir-api/internal/service"
	"kasir-api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	service *service.ReportService
}

func NewReportHandler(s *service.ReportService) *ReportHandler {
	return &ReportHandler{service: s}
}

// Report godoc
// @Summary Get reports
// @Description Ambil report penjualan
// @Tags Report
// @Produce json
// @Param start query string false "Start date (format: YYYY-MM-DD)"
// @Param end   query string false "End date (format: YYYY-MM-DD)"
// @Success 200 {array} model.AllReport
// @Failure 400 {object} map[string]string
// @Router /report [get]
func (h *ReportHandler) Report(c *gin.Context) {
	startDate, err := utils.ParseDateParam(c, "start")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	endDate, err := utils.ParseDateParam(c, "end")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if endDate != nil {
		*endDate = endDate.AddDate(0, 0, 1)
	}

	report, err := h.service.Report(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bestProduct, err := h.service.BestSeller(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := gin.H{
		"best_product": bestProduct,
	}

	resp["total_revenue"] = report.TotalRevenue
	resp["total_transaction"] = report.TotalTransaction

	c.JSON(http.StatusOK, resp)

}
