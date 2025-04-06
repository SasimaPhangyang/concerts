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

func (h *ReportHandler) GetSalesReport(c *gin.Context) {
	date := c.DefaultQuery("date", "")
	eventID := c.DefaultQuery("event_id", "")

	// เรียกใช้ Service เพื่อดึงข้อมูลรายงานการขาย
	report, err := h.Service.GetSalesReport(date, eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate sales report"})
		return
	}
	// ส่งข้อมูลรายงานกลับไปในรูป JSON
	c.JSON(http.StatusOK, report)
}

func (h *ReportHandler) GetSalesBySource(c *gin.Context) {
	// ใช้พารามิเตอร์เดือนเพื่อดึงข้อมูล
	month := c.DefaultQuery("month", "")

	// ตรวจสอบว่าเดือนต้องไม่เป็นค่าว่าง
	if month == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Month parameter is required"})
		return
	}

	// เรียกใช้ Service เพื่อดึงข้อมูลรายงานการขายตามแหล่งที่มา
	report, err := h.Service.GetSalesBySourceReport(month)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate sales by source report"})
		return
	}
	// ส่งข้อมูลรายงานกลับไปในรูป JSON
	c.JSON(http.StatusOK, report)
}
