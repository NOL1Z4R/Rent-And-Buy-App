package saleDtos

type CreateSaleDto struct {
	CustomerId uint `json:"CustomerId" binding:"required"`
	PostId     uint `json:"PostId" binding:"required"`
}
