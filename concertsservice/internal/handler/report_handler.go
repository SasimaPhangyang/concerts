package handler

import (
	"concerts/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	Service service.ReportService
}

func NewReportHandler(service service.ReportService) *ReportHandler {
	return &ReportHandler{Service: service}
}

// GetSalesReport ดึงข้อมูลยอดขายตาม product
func (h *ReportHandler) GetSalesReport(c *gin.Context) {
	product := c.DefaultQuery("product", "")

	if product == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product parameter is required"})
		return
	}

	report, err := h.Service.GetSalesReport(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate sales report"})
		return
	}

	c.JSON(http.StatusOK, report)
}

// GetSalesBySource ดึงข้อมูลยอดขายแยกตาม source
func (h *ReportHandler) GetSalesBySource(c *gin.Context) {
	report, err := h.Service.GetSalesBySourceReport()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate sales by source report"})
		return
	}

	c.JSON(http.StatusOK, report)
}
