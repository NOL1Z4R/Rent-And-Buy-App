package service

import (
	"Rent-And-Buy-App/internal/dtos/carDtos"
	"Rent-And-Buy-App/internal/dtos/houseDtos"
	"Rent-And-Buy-App/internal/dtos/productDtos"
	"Rent-And-Buy-App/internal/entity"
	"Rent-And-Buy-App/internal/repository"
	"Rent-And-Buy-App/pkg/Converter"
	"errors"
)

type ProductService struct {
	repo         *repository.ProductRepository
	carService   *CarService
	houseService *HouseService
}

func NewProductService(repo *repository.ProductRepository, carService *CarService, houseService *HouseService) *ProductService {
	return &ProductService{repo: repo, carService: carService, houseService: houseService}
}

func (ps *ProductService) GetAll() ([]*productDtos.ProductDto, error) {
	products, err := ps.repo.GetAll()
	var productDtos []*productDtos.ProductDto

	for _, product := range products {
		productDto, err := ps.GetProductDto(product)
		if err != nil {
			return nil, err
		}
		productDtos = append(productDtos, productDto)
	}

	return productDtos, err
}

func (ps *ProductService) GetById(id string) (*productDtos.ProductDto, error) {
	Id := Converter.StringToUint(id)
	product, err := ps.repo.GetById(Id)
	if product == nil {
		return nil, errors.New("Product not found")
	}
	productDto, err := ps.GetProductDto(product)

	if err != nil {
		return nil, err
	}
	return productDto, nil
}

func (ps *ProductService) GetEntityById(id string) (*entity.Product, error) {
	Id := Converter.StringToUint(id)
	product, err := ps.repo.GetById(Id)
	if product == nil {
		return nil, errors.New("Product not found")
	}

	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) Create(createProductDto *productDtos.CreateProductDto) (*productDtos.ProductDto, error) {

	car := &entity.Car{}
	house := &entity.House{}
	product := entity.Product{}
	var err error

	if createProductDto.Car == nil && createProductDto.House == nil {
		return nil, errors.New("You must create either Car or House")
	}
	if createProductDto.Car != nil && createProductDto.House != nil {
		return nil, errors.New("You cannot create both Car and House")
	}

	if createProductDto.Car != nil {
		carEntity := entity.Car{
			Brand:    createProductDto.Car.Brand,
			CarModel: createProductDto.Car.CarModel,
			Milage:   createProductDto.Car.Milage,
			Year:     createProductDto.Car.Year,
			Plate:    createProductDto.Car.Plate,
		}

		car, err = ps.carService.CreateCar(&carEntity)
		if err != nil {
			return nil, err
		}
		product.CarId = &car.ID
		product.HouseId = nil
	}

	if createProductDto.House != nil {
		houseEntity := entity.House{
			Adress:      createProductDto.House.Adress,
			SquareMeter: createProductDto.House.SquareMeter,
			Year:        createProductDto.House.Year,
		}

		house, err = ps.houseService.Create(&houseEntity)
		if err != nil {
			return nil, err
		}
		product.CarId = nil
		product.HouseId = &house.ID
	}

	err = ps.repo.Create(&product)
	if err != nil {
		return nil, err
	}

	productDto, err := ps.GetProductDto(&product)
	if err != nil {
		return nil, err
	}

	return productDto, nil
}

func (ps *ProductService) Update(product *entity.Product, updateDto *productDtos.UpdateProductDto) (*productDtos.ProductDto, error) {

	err := ps.updateInnerMaterials(product, updateDto)
	if err != nil {
		return nil, err
	}

	productDto, err := ps.GetProductDto(product)
	if err != nil {
		return nil, err
	}

	return productDto, nil
}

func (ps *ProductService) Delete(product *entity.Product) error {
	var err error

	carId := Converter.UintToString(*product.CarId)
	err = ps.carService.DeleteCar(carId)
	if err != nil {
		return err
	}

	houseId := Converter.UintToString(*product.HouseId)
	err = ps.houseService.Delete(houseId)

	err = ps.repo.Delete(product)
	if err != nil {
		return err
	}
	return nil
}

func (ps *ProductService) GetProductDto(product *entity.Product) (*productDtos.ProductDto, error) {
	var car *carDtos.CarDto
	var house *houseDtos.HouseDto
	var err error

	if product.CarId != nil {
		carUid := *product.CarId
		carId := Converter.UintToString(carUid)
		car, err = ps.carService.GetCarDto(carId)
		if err != nil {
			return nil, err
		}
	}

	if product.HouseId != nil {
		houseUid := *product.HouseId
		houseId := Converter.UintToString(houseUid)
		house, err = ps.houseService.GetHouseDto(houseId)
		if err != nil {
			return nil, err
		}
	}

	productDto := &productDtos.ProductDto{
		ID:    product.Model.ID,
		Car:   car,
		House: house,
	}

	if err = _checkMaterials(productDto); err != nil {
		return nil, err
	}

	return productDto, nil
}

// #region Inner Funcs
func _checkMaterials(productDto *productDtos.ProductDto) error {
	car := productDto.Car
	house := productDto.House

	if car == nil && house == nil {
		return errors.New("You must add either car or house")
	}
	if car != nil && house != nil {
		return errors.New("You can only add car or house")
	}

	return nil
}

func (ps *ProductService) changeInnerMaterials(productDto *productDtos.ProductDto, updateDto *productDtos.UpdateProductDto) (*productDtos.ProductDto, error) {

	//product := productDtos.ProductDto{}
	//house := product.House
	//car := product.Car
	//updateHouse := updateDto.House
	//updateCar := updateDto.Car

	if updateDto.Car != nil {
		productDto.Car.Brand = updateDto.Car.Brand
		productDto.Car.CarModel = updateDto.Car.CarModel
		productDto.Car.Year = updateDto.Car.Year
		productDto.Car.Milage = updateDto.Car.Milage
		productDto.Car.Plate = updateDto.Car.Plate
	}

	if updateDto.House != nil {
		productDto.House.Adress = updateDto.House.Adress
		productDto.House.SquareMeter = updateDto.House.SquareMeter
		productDto.House.Year = updateDto.House.Year
	}

	//productDto.Car = car
	//productDto.House = house

	return productDto, nil
}

func (ps *ProductService) updateInnerMaterials(product *entity.Product, updateDto *productDtos.UpdateProductDto) error {

	if updateDto.Car != nil && updateDto.House != nil {
		return errors.New("You can only add car or house")
	}
	if updateDto.Car == nil && updateDto.House == nil {
		return errors.New("You must add either car or house")
	}

	if updateDto.Car != nil {
		carId := Converter.UintToString(*product.CarId)
		car, err := ps.carService.GetById(carId)
		if err != nil {
			return err
		}

		//car.Brand = updateDto.Car.Brand
		//car.CarModel = updateDto.Car.CarModel
		//car.Year = updateDto.Car.Year
		//car.Milage = updateDto.Car.Milage
		//car.Plate = updateDto.Car.Plate

		_, err = ps.carService.UpdateCar(car, updateDto.Car)
		if err != nil {
			return err
		}

	}

	if updateDto.House != nil {

		houseId := Converter.UintToString(*product.HouseId)
		house, err := ps.houseService.GetById(houseId)
		if err != nil {
			return err
		}

		//house.Adress = updateDto.House.Adress
		//house.SquareMeter = updateDto.House.SquareMeter
		//house.Year = updateDto.House.Year

		_, err = ps.houseService.Update(house, updateDto.House)
		if err != nil {
			return err
		}

	}

	return nil
}

//#endregion
