package infrastructure

import (
	"apiClienteServidor/domain"
	"log"
)

// CajaRepository maneja la persistencia de la caja en la base de datos
type CajaRepository struct{}

func NewCajaRepository() *CajaRepository {
	return &CajaRepository{}
}

// ObtenerCaja obtiene la caja de la base de datos
func (r *CajaRepository) ObtenerCaja() (*domain.Caja, error) {
	var caja domain.Caja
	if err := DB.First(&caja).Error; err != nil {
		log.Println("Error al obtener la caja:", err)
		return nil, err
	}
	return &caja, nil
}
