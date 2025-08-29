package carDtos

type CarDto struct {
	ID       uint   `json:"ID" gorm:"primary_key"`
	Brand    string `json:"Brand" binding:"required"`
	CarModel string `json:"CarModel" binding:"required"`
	Year     uint   `json:"Year" binding:"required"`
	Milage   uint   `json:"Milage" binding:"required"`
	Plate    string `json:"Plate" binding:"required" gorm:"unique"`
}
