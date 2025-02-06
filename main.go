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

	// Inicializar repositorios
	repoCar := infrastructure.NewCarRepository()
	repoCaja := infrastructure.NewCajaRepository() // Asegúrate de que esta función existe

	// Inicializar servicios
	serviceCar := application.NewCarService(repoCar)
	serviceCaja := application.NewCajaService(repoCaja) // Esto debe estar definido en `application/caja_service.go`

	// Inicializar handlers
	handlerCar := handlers.NewCarHandler(serviceCar)
	handlerCaja := handlers.NewCajaHandler(serviceCaja) // Esto debe estar definido en `handlers/caja_handler.go`

	// Configurar el servidor Gin
	r := gin.Default()

	// Rutas de Carros
	r.GET("/cars", handlerCar.GetCars)
	r.GET("/cars/:id", handlerCar.GetCar)
	r.POST("/cars", handlerCar.CreateCar)
	r.PUT("/cars/:id", handlerCar.UpdateCar)
	r.DELETE("/cars/:id", handlerCar.DeleteCar)

	// Ruta para obtener el dinero de la caja
	r.GET("/caja", handlerCaja.GetDinero)

	// Iniciar el servidor
	r.Run(":8080")
}
