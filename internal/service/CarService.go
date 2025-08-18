package service

import (
	"Rent-And-Buy-App/internal/dtos/carDtos"
	"Rent-And-Buy-App/internal/entity"
	"Rent-And-Buy-App/internal/repository"
	"Rent-And-Buy-App/pkg/Converter"
	"errors"
)

type CarService struct {
	repo repository.CarRepository
}

func NewCarService(repo repository.CarRepository) *CarService {
	return &CarService{repo: repo}
}

func (cs *CarService) GetAll() ([]*entity.Car, error) {
	var cars []*entity.Car
	cars, err := cs.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return cars, nil
}

func (cs *CarService) GetById(id string) (*entity.Car, error) {
	Id := Converter.StringToUint(id)
	car, err := cs.repo.GetById(Id)
	if err != nil {
		return nil, err
	}

	return car, nil
}

func (cs *CarService) GetByPlate(plate string) (*entity.Car, error) {
	car, err := cs.repo.GetByPlate(plate)
	if err != nil {
		return nil, err
	}
	return car, nil
}

func (cs *CarService) CreateCar(car *entity.Car) (entity.Car, error) {
	plate := car.Plate
	if isExist, err := cs.repo.GetByPlate(plate); err != nil || isExist != nil {
		errors.New("This car already exists")
	}

	err := cs.repo.Create(car)
	return *car, err
}

func (cs *CarService) UpdateCar(car *entity.Car, updateDto *carDtos.UpdateCarDto) (*entity.Car, error) {
	plate := updateDto.Plate
	if isExist, err := cs.repo.GetByPlate(plate); err != nil || isExist != nil {
		errors.New("This car already exists")
	}

	car.Brand = updateDto.Brand
	car.CarModel = updateDto.CarModel
	car.Year = updateDto.Year
	car.Milage = updateDto.Milage
	car.Plate = updateDto.Plate

	err := cs.repo.Update(car)
	return car, err
}

func (cs *CarService) DeleteCar(id string) error {
	car, err := cs.GetById(id)
	if err != nil {
		return err
	}
	err = cs.repo.Delete(car)
	if err != nil {
		return err
	}
	return nil
}
