package carDtos

type CreateCarDto struct {
	Brand    string `json:"Brand" binding:"required"`
	CarModel string `json:"CarModel" binding:"required"`
	Year     uint   `json:"Year" binding:"required"`
	Milage   uint   `json:"Milage" binding:"required"`
}
