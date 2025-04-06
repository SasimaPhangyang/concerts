package handler

import (
	"concerts/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ConcertHandler struct {
	service service.ConcertService
}

func NewConcertHandler(service service.ConcertService) *ConcertHandler {
	return &ConcertHandler{service: service}
}

func (h *ConcertHandler) GetAllConcerts(c *gin.Context) {
	concerts, err := h.service.GetAllConcerts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, concerts)
}

func (h *ConcertHandler) GetConcertByID(c *gin.Context) {
	// แปลงจาก string เป็น int
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	concert, err := h.service.GetConcertByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, concert)
}

// เพิ่ม handler สำหรับค้นหาคอนเสิร์ต
func (h *ConcertHandler) SearchConcerts(c *gin.Context) {
	query := c.DefaultQuery("query", "") // รับค่า query จาก URL query parameter
	concerts, err := h.service.SearchConcerts(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, concerts)
}
