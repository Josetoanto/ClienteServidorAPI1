package main

import (
	"apiClienteServidor/application"
	"apiClienteServidor/handlers"
	"apiClienteServidor/infrastructure"
	"github.com/gin-gonic/gin"
)

func main() {
	// Conectar a la base de datos
	infrastructure.ConnectDatabase()

	// Inicializar el repositorio, servicio y handler
	repo := infrastructure.NewCarRepository()
	service := application.NewCarService(repo)
	handler := handlers.NewCarHandler(service)

	// Configurar el servidor Gin
	r := gin.Default()

	r.GET("/cars", handler.GetCars)
	r.GET("/cars/:id", handler.GetCar)
	r.POST("/cars", handler.CreateCar)
	r.PUT("/cars/:id", handler.UpdateCar)
	r.DELETE("/cars/:id", handler.DeleteCar)

	// Iniciar el servidor
	r.Run(":8080")
}
