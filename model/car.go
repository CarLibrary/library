package model

import (
	"gorm.io/gorm"
)

// CarBand 品牌
type CarBand struct {
	gorm.Model
	Band string `gorm:"primarykey"`
	Logo string `gorm:"not null;size:255"`
}

// CarSeries 车系
type CarSeries struct {
	gorm.Model

	//品牌名
	Band string `gorm:"not null;primarykey"`


	//车系名
	CarSeries string `gorm:"not null;primarykey"`

	//车系图片
	CarPicture string `gorm:"size:255"`

	//评分
	Point string `gorm:"default:0"`

	//报价
	Price string `gorm:"default:暂无报价"`

}

// CarModel 车型
type CarModel struct {
	gorm.Model

	//车系名
	CarSeries string `gorm:"not null;primarykey"`

	//车型名
	CarModel string `gorm:"not null"`

	//价格
	Price string `gorm:"default:暂无报价"`

}

// FindALLCarBand 查看所有品牌
func FindALLCarBand() (list *[]CarBand,err error){
	if err=db.Table("car_bands").Find(list).Error;err!=nil{
		return new([]CarBand), err
	}
	return list, nil
}

// FindAllCarSeries 查看某品牌的全部车系
func FindAllCarSeries(band string)(list *[]CarSeries,err error){

	if err=db.Table("car_series").Joins("JOIN car_bands ON car_series.band = car_bands.band AND car_series.band = ?", band).Find(list).Error;err!=nil{
		return new([]CarSeries), err
	}
	return list, nil


}

// FindAllCarModel 查看某品牌的某车系的全部车型
func FindAllCarModel(band string,series string)(list *[]CarModel,err error){
	if err=db.Table("car_series").Joins("JOIN car_series ON car_models.car_series = car_series.car_series AND car_series.band = ?", band).Where("car_models.carseries = ?", series).Find(list).Error;err!=nil{
		return new([]CarModel), err
	}

	return list, nil
}