package houseDtos

type UpdateHouseDto struct {
	Adress      string `json:"Adress" binding:"required"`
	SquareMeter uint   `json:"SquareMeter" binding:"required"`
	Year        uint   `json:"Year" binding:"required"`
}
