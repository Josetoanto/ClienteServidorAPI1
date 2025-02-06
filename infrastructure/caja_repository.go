package infrastructure

import (
	"apiClienteServidor/domain"
	"log"
)

type CajaRepository struct{}

func NewCajaRepository() *CajaRepository {
	return &CajaRepository{}
}

func (r *CajaRepository) ObtenerCaja() (*domain.Caja, error) {
	var caja domain.Caja
	if err := DB.First(&caja).Error; err != nil {
		log.Println("Error al obtener la caja:", err)
		return nil, err
	}
	return &caja, nil
}
