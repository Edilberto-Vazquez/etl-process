package main

import (
	"log"

	"github.com/Edilberto-Vazquez/weathercloud-etl-process/db/conexion"
	"github.com/Edilberto-Vazquez/weathercloud-etl-process/db/migrations"
	"github.com/Edilberto-Vazquez/weathercloud-etl-process/etl"
	"github.com/Edilberto-Vazquez/weathercloud-etl-process/services"
	"github.com/Edilberto-Vazquez/weathercloud-etl-process/utils"
)

func errorFunction(err error, table string, rows int64) {
	if err != nil {
		log.Panicln(err)
	} else {
		log.Println("Rows inserted in the table", table, rows)
	}
}

func main() {
	var rows int64
	var err error

	conexion.InitCon()
	defer conexion.CloseCon()

	migrations.Migrate()

	logRoot := "/media/potatofy/Nuevo vol/DataSets/Conjuntos-originales/medidor-campo-electrico"
	efmRoot := "/media/potatofy/Nuevo vol/DataSets/Conjuntos-originales/medidor-campo-electrico"
	wcRoot := "/media/potatofy/Nuevo vol/DataSets/Conjuntos-originales/estacion-meteorologica"

	logFiles, _ := utils.ReadDirectory(logRoot, "log")
	efmFiles, _ := utils.ReadDirectory(efmRoot, "efm")
	wcFiles, _ := utils.ReadDirectory(wcRoot, "csv")

	etl.ProcessFiles(logFiles, efmFiles, wcFiles)

	log.Println(len(etl.ElectricFieldSlice), len(etl.WeatherCloudSlice))

	efService := services.NewElectricFieldService()
	rows, err = efService.CreateRecords()
	errorFunction(err, "electric_field", rows)

	wcService := services.NewWeatherCloudService()
	rows, err = wcService.CreateRecords()
	errorFunction(err, "weather_cloud", rows)
}
