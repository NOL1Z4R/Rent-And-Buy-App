package service

import (
	"Rent-And-Buy-App/internal/dtos/saleDtos"
	"Rent-And-Buy-App/internal/entity"
	"Rent-And-Buy-App/internal/repository"
	"Rent-And-Buy-App/pkg/Converter"
	"errors"
	"time"
)

type SaleService struct {
	repo        *repository.SaleRepository
	postService *PostService
	userService *UserService
}

func NewSaleService(repo *repository.SaleRepository, userService *UserService, postService *PostService) *SaleService {
	return &SaleService{repo: repo, userService: userService, postService: postService}
}

func (ss *SaleService) GetAll() ([]*saleDtos.SaleDto, error) {
	var sales []*entity.Sale
	var SaleDtos []*saleDtos.SaleDto
	var err error
	sales, err = ss.repo.GetAll()
	if err != nil {
		return nil, err
	}

	for _, sale := range sales {
		var tempSaleDto *saleDtos.SaleDto
		tempSaleDto, err = ss.GetSaleDto(sale)
		if err != nil {
			return nil, err
		}

		SaleDtos = append(SaleDtos, tempSaleDto)
	}
	return SaleDtos, nil
}

func (ss *SaleService) GetById(id string) (*saleDtos.SaleDto, error) {
	saleId := Converter.StringToUint(id)
	sale, err := ss.repo.GetById(saleId)
	if err != nil {
		return nil, err
	}

	saleDto, err := ss.GetSaleDto(sale)
	if err != nil {
		return nil, err
	}

	return saleDto, nil
}

func (ss *SaleService) CreateSale(sale *saleDtos.CreateSaleDto) (*saleDtos.SaleDto, error) {
	if sale.CustomerId == 0 {
		return nil, errors.New("Customer can not be null")
	}

	userId := Converter.UintToString(sale.CustomerId)
	user, err := ss.userService.GetByID(userId)

	if err != nil {
		return nil, err
	} else if *user.IsSeller { // Customer, not seller
		return nil, errors.New("Only Customers can buy")
	}

	saleEntity := &entity.Sale{
		CustomerId: sale.CustomerId,
		SaleDate:   time.Now(),
		PostId:     sale.PostId,
	}

	newSale, err := ss.repo.Create(saleEntity)
	if err != nil {
		return nil, err
	}

	saleDto, err := ss.GetSaleDto(newSale)
	if err != nil {
		return nil, err
	}

	return saleDto, nil
}

//TODO might not need for update part
//func (ss *SaleService) UpdateSale(sale *entity.Sale) (*saleDtos.SaleDto, error) {
//	userId := Converter.UintToString(sale.CustomerId)
//	user, err := ss.userService.GetByID(userId)
//	if !*user.IsSeller {
//		return nil, errors.New("Only Customers can buy")
//	}
//}

// this function is only Confirm the Payment. Will add checking payment type in future
func (ss *SaleService) ConfirmPayment(paymentDto *saleDtos.PaymentDto) (*saleDtos.SaleDto, error) {
	sale, err := ss.repo.GetById(paymentDto.SaleId)
	if err != nil {
		return nil, err
	}

	if sale.CustomerId != paymentDto.CustomerId {
		return nil, errors.New("This user is not the same customer on this payment")
	}

	sale.SaleDate = time.Now()
	saleDto, err := ss.GetSaleDto(sale)
	if err != nil {
		return nil, err
	}

	return saleDto, nil
}

func (ss *SaleService) Delete(id string) error {
	saleId := Converter.StringToUint(id)
	sale, err := ss.repo.GetById(saleId)
	if err != nil {
		return err
	}

	userId := Converter.UintToString(sale.CustomerId)
	user, err := ss.userService.GetByID(userId)

	// Customer, not Seller
	if *user.IsSeller {
		return errors.New("Only Customers can buy")
	} else if err != nil {
		return err
	}

	err = ss.repo.Delete(sale)
	return err
}

func (ss *SaleService) GetSaleDto(sale *entity.Sale) (*saleDtos.SaleDto, error) {
	postId := Converter.UintToString(sale.PostId)
	post, err := ss.postService.GetById(postId)
	if err != nil {
		return nil, err
	}

	saleDto := &saleDtos.SaleDto{
		ID:         sale.ID,
		CustomerId: sale.CustomerId,
		SaleDate:   sale.SaleDate,
		Post:       *post,
	}

	return saleDto, nil
}
