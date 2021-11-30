package services

import (
	"github.com/Edilberto-Vazquez/weathercloud-etl-process/db/conexion"
	"github.com/Edilberto-Vazquez/weathercloud-etl-process/etl"
)

type ElectricFieldService struct{}

func NewElectricFieldService() *ElectricFieldService {
	return &ElectricFieldService{}
}

func (ElectricFieldService) CreateRecords() (int64, error) {
	result := conexion.DBCon.CreateInBatches(etl.ElectricFieldSlice, 1000)
	return result.RowsAffected, result.Error
}
