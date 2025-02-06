package handlers

import (
	"apiClienteServidor/application"
	"apiClienteServidor/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CarHandler maneja las peticiones HTTP
type CarHandler struct {
	service *application.CarService
}

func NewCarHandler(service *application.CarService) *CarHandler {
	return &CarHandler{service: service}
}

func (h *CarHandler) GetCars(c *gin.Context) {
	cars, err := h.service.GetAllCars()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cars)
}

func (h *CarHandler) GetCar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	car, err := h.service.GetCarByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}
	c.JSON(http.StatusOK, car)
}

func (h *CarHandler) CreateCar(c *gin.Context) {
	var car domain.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newCar, err := h.service.CreateCar(car)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newCar)
}

func (h *CarHandler) UpdateCar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var car domain.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	car.ID = uint(id)
	updatedCar, err := h.service.UpdateCar(car)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedCar)
}

func (h *CarHandler) DeleteCar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.DeleteCar(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Car deleted successfully"})
}
