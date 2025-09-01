package saleDtos

type PaymentDto struct {
	SaleId     uint `json:"SaleId" gorm:"not null" binding:"required"`
	CustomerId uint `json:"CustomerId" gorm:"not null" binding:"required"`
}
