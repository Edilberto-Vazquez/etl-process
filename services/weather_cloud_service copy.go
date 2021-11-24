package services

import (
	"github.com/Edilberto-Vazquez/weathercloud-etl-process/db/conexion"
	"github.com/Edilberto-Vazquez/weathercloud-etl-process/etl"
)

type WeatherCloudService struct{}

func NewWeatherCloudService() *WeatherCloudService {
	return &WeatherCloudService{}
}

func (WeatherCloudService) CreateRecords() (int64, error) {
	result := conexion.DBCon.CreateInBatches(etl.WeatherCloudSlice, 5000)
	return result.RowsAffected, result.Error
}
