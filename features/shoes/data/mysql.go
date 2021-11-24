package data

import (
	"justrun/features/shoes"

	"gorm.io/gorm"
)

type shoesRepository struct {
	DB *gorm.DB
}

func NewShoesRepository(DB *gorm.DB) shoes.Data {
	return &shoesRepository{DB}
}

func (sr *shoesRepository) InsertData(data shoes.Shoes) (resp shoes.Shoes, err error) {

	record := fromShoes(data)

	if err := sr.DB.Create(&record).Error; err != nil {
		return shoes.Shoes{}, err
	}

	return data, nil
}

func (sr *shoesRepository) SelectData(name string) (resp []shoes.Shoes) {

	var record []Shoes

	if err := sr.DB.Preload("TagArticle").Find(&record).Error; err != nil {
		return []shoes.Shoes{}
	}

	return toShoesList(record)
}
