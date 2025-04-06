package handler

import (
	"concerts/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContentTemplateHandler struct {
	Service service.ContentTemplateService
}

func NewContentTemplateHandler(service service.ContentTemplateService) *ContentTemplateHandler {
	return &ContentTemplateHandler{Service: service}
}

func (h *ContentTemplateHandler) GetContentTemplates(c *gin.Context) {
	templates, err := h.Service.GetContentTemplates()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch content templates"})
		return
	}
	c.JSON(http.StatusOK, templates)
}
