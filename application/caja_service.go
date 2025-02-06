package application

import (
	"apiClienteServidor/infrastructure"
)

type CajaService struct {
	repo *infrastructure.CajaRepository
}

func NewCajaService(repo *infrastructure.CajaRepository) *CajaService {
	return &CajaService{repo: repo}
}

func (s *CajaService) ObtenerDinero() (float64, error) {
	caja, err := s.repo.ObtenerCaja()
	if err != nil {
		return 0, err
	}
	return caja.Dinero, nil
}
