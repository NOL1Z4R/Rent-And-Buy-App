package service

import (
	"Rent-And-Buy-App/internal/dtos/houseDtos"
	"Rent-And-Buy-App/internal/entity"
	"Rent-And-Buy-App/internal/repository"
	"Rent-And-Buy-App/pkg/Converter"
	"errors"
)

type HouseService struct {
	repo *repository.HouseRepository
}

func NewHouseService(repo *repository.HouseRepository) *HouseService {
	return &HouseService{repo: repo}
}

func (hs *HouseService) GetAll() ([]*entity.House, error) {
	var houses []*entity.House
	houses, err := hs.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return houses, nil
}

func (hs *HouseService) GetById(id string) (*entity.House, error) {
	Id := Converter.StringToUint(id)
	house, err := hs.repo.GetById(Id)
	if err != nil {
		return nil, err
	}
	return house, nil
}

// TODO adress control can be added in future
func (hs *HouseService) Create(house *entity.House) (*entity.House, error) {
	if house == nil {
		return nil, errors.New("house is nil")
	}
	err := hs.repo.Create(house)
	if err != nil {
		return nil, err
	}
	return house, nil
}

func (hs *HouseService) Update(house *entity.House, updateDto *houseDtos.UpdateHouseDto) (*entity.House, error) {
	id := house.ID

	if isExist, err := hs.repo.GetById(id); err != nil || isExist == nil {
		return nil, errors.New("This house is not exists")
	}

	house.Year = updateDto.Year
	house.SquareMeter = updateDto.SquareMeter
	house.Adress = updateDto.Adress

	err := hs.repo.Update(house)
	if err != nil {
		return nil, err
	}

	return house, nil
}
func (hs *HouseService) Delete(id string) error {
	house, err := hs.GetById(id)
	if err != nil {
		return err
	}

	err = hs.repo.Delete(house)
	if err != nil {
		return errors.New("could not deleted this house")
	}
	return nil
}

func (hs *HouseService) GetHouseDto(id string) (*houseDtos.HouseDto, error) {
	house, err := hs.GetById(id)
	if err != nil {
		return nil, err
	}

	houseDto := &houseDtos.HouseDto{
		ID:          house.ID,
		Adress:      house.Adress,
		SquareMeter: house.SquareMeter,
		Year:        house.Year,
	}

	return houseDto, nil
}
