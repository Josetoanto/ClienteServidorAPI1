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

	repoCar := infrastructure.NewCarRepository()
	repoCaja := infrastructure.NewCajaRepository() 

	serviceCar := application.NewCarService(repoCar)
	serviceCaja := application.NewCajaService(repoCaja) 

	handlerCar := handlers.NewCarHandler(serviceCar)
	handlerCaja := handlers.NewCajaHandler(serviceCaja) 

	r := gin.Default()

	r.GET("/cars", handlerCar.GetCars)
	r.GET("/cars/:id", handlerCar.GetCar)
	r.POST("/cars", handlerCar.CreateCar)
	r.PUT("/cars/:id", handlerCar.UpdateCar)
	r.DELETE("/cars/:id", handlerCar.DeleteCar)

	r.GET("/caja", handlerCaja.GetDinero)

	r.Run(":8080")
}
