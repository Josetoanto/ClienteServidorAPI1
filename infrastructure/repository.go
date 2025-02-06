package infrastructure

import (
	"apiClienteServidor/domain"
)

type CarRepositoryImpl struct{}

func NewCarRepository() domain.CarRepository {
	return &CarRepositoryImpl{}
}

func (r *CarRepositoryImpl) GetAll() ([]domain.Car, error) {
	var cars []domain.Car
	if err := DB.Find(&cars).Error; err != nil {
		return nil, err
	}
	return cars, nil
}

func (r *CarRepositoryImpl) GetByID(id uint) (domain.Car, error) {
	var car domain.Car
	if err := DB.First(&car, id).Error; err != nil {
		return car, err
	}
	return car, nil
}

func (r *CarRepositoryImpl) Create(car domain.Car) (domain.Car, error) {
	if err := DB.Create(&car).Error; err != nil {
		return car, err
	}
	return car, nil
}

func (r *CarRepositoryImpl) Update(car domain.Car) (domain.Car, error) {
	if err := DB.Save(&car).Error; err != nil {
		return car, err
	}
	return car, nil
}

func (r *CarRepositoryImpl) Delete(id uint) error {
	if err := DB.Delete(&domain.Car{}, id).Error; err != nil {
		return err
	}
	return nil
}
