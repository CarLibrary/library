package serializer

import (
	"CarLibrary/carlibrary/model"
	pb "github.com/CarLibrary/proto/carlibrary"
)


func BuildCarBand(carband *model.CarBand) *pb.CarBand {

	return &pb.CarBand{
		CarBand:    carband.Band,
		CarBandLogo: carband.Logo,
	}
}

func BuildCarSeries(carseries *model.CarSeries) *pb.CarSeries{
	return &pb.CarSeries{
		CarSeries: carseries.CarSeries,
		CarPicture: carseries.CarPicture,
		CarPoint:   carseries.Point,
		CarPrice:   carseries.Price,
	}
}

func BuildCarModel(carmodel *model.CarModel) *pb.CarModel  {
	return &pb.CarModel{
		CarSeries:     carmodel.CarSeries,
		CarModel:      carmodel.CarModel,
		CarModelPrice: carmodel.Price,
	}
}