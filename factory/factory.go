package factory

import (
	"justrun/config"
	"justrun/features/shoes/bussiness"
	"justrun/features/shoes/data"
	"justrun/features/shoes/presentation"
)

type Presenter struct {
	ShoesPresentation *presentation.ShoesHandler
}

func Init() Presenter {

	shoeData := data.NewShoesRepository(config.DB)
	shoeBussiness := bussiness.NewShoeBussiness(shoeData)
	shoePresentation := presentation.AddShoesHandler(shoeBussiness)

	return Presenter{
		ShoesPresentation: shoePresentation,
	}
}
