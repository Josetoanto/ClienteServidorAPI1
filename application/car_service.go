package application

import (
	"apiClienteServidor/domain"
)

type CarService struct {
	repo domain.CarRepository
}

func NewCarService(repo domain.CarRepository) *CarService {
	return &CarService{repo: repo}
}

func (s *CarService) GetAllCars() ([]domain.Car, error) {
	return s.repo.GetAll()
}

func (s *CarService) GetCarByID(id uint) (domain.Car, error) {
	return s.repo.GetByID(id)
}

func (s *CarService) CreateCar(car domain.Car) (domain.Car, error) {
	return s.repo.Create(car)
}

func (s *CarService) UpdateCar(car domain.Car) (domain.Car, error) {
	return s.repo.Update(car)
}

func (s *CarService) DeleteCar(id uint) error {
	return s.repo.Delete(id)
}
