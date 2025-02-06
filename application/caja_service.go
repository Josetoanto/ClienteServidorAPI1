package application

import (
	"apiClienteServidor/infrastructure"
)

// CajaService maneja la lógica de negocio de la caja
type CajaService struct {
	repo *infrastructure.CajaRepository
}

func NewCajaService(repo *infrastructure.CajaRepository) *CajaService {
	return &CajaService{repo: repo}
}

// ObtenerDinero devuelve el dinero almacenado en la caja
func (s *CajaService) ObtenerDinero() (float64, error) {
	caja, err := s.repo.ObtenerCaja()
	if err != nil {
		return 0, err
	}
	return caja.Dinero, nil
}
