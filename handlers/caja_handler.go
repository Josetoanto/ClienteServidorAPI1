package handlers

import (
	"apiClienteServidor/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CajaHandler struct {
	service *application.CajaService
}

func NewCajaHandler(service *application.CajaService) *CajaHandler {
	return &CajaHandler{service: service}
}

// GetDinero maneja la solicitud GET para obtener el dinero de la caja
func (h *CajaHandler) GetDinero(c *gin.Context) {
	dinero, err := h.service.ObtenerDinero()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener el dinero"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"dinero": dinero})
}
