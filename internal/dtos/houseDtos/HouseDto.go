package houseDtos

type HouseDto struct {
	ID          uint   `json:"ID" gorm:"primary_key"`
	Adress      string `json:"Adress" binding:"required"`
	SquareMeter uint   `json:"SquareMeter" binding:"required"`
	Year        uint   `json:"Year" binding:"required"`
}
