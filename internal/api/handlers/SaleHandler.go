package handlers

import (
	"Rent-And-Buy-App/internal/dtos/saleDtos"
	"Rent-And-Buy-App/internal/service"
	"Rent-And-Buy-App/pkg/Response"
	"github.com/gin-gonic/gin"
)

type SaleHandler struct {
	service *service.SaleService
}

func NewSaleHandler(service *service.SaleService) *SaleHandler {
	return &SaleHandler{service: service}
}

func (sh *SaleHandler) GetSaleAll(c *gin.Context) {
	var SaleDtos []*saleDtos.SaleDto
	var err error
	SaleDtos, err = sh.service.GetAll()
	if err != nil {
		Response.Error(c, "Could not get sales")
		return
	}

	Response.OK(c, SaleDtos)
}

func (sh *SaleHandler) GetSaleById(c *gin.Context) {
	saleId := c.Param("id")
	sale, err := sh.service.GetById(saleId)
	if err != nil {
		Response.Error(c, "Sale not found")
		return
	}

	Response.OK(c, sale)
}

func (sh *SaleHandler) CreateSale(c *gin.Context) {
	var createDto *saleDtos.CreateSaleDto
	if err := c.ShouldBindJSON(&createDto); err != nil {
		Response.Error(c, "Can not bind request body")
		return
	}

	sale, err := sh.service.CreateSale(createDto)
	if err != nil {
		Response.Error(c, err.Error())
		return
	}

	Response.OK(c, sale)
}

func (sh *SaleHandler) ConfirmSale(c *gin.Context) {
	saleId := c.Param("id")
	sale, err := sh.service.GetById(saleId)
	if err != nil {
		Response.Error(c, "Sale not found")
		return
	}

	paymentDto := &saleDtos.PaymentDto{
		SaleId:     sale.ID,
		CustomerId: sale.CustomerId,
	}

	saleDto, err := sh.service.ConfirmPayment(paymentDto)
	if err != nil {
		Response.Error(c, "Sale not found")
		return
	}
	Response.OK(c, saleDto)
}

func (sh *SaleHandler) DeleteSale(c *gin.Context) {
	saleId := c.Param("id")
	err := sh.service.Delete(saleId)
	if err != nil {
		Response.Error(c, "Sale not found")
		return
	}

	Response.OK(c, "Sale deleted")
}
