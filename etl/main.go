package etl

import (
	"fmt"
	"time"

	"github.com/Edilberto-Vazquez/weathercloud-etl-process/db/models"
)

var (
	WeatherCloudSlice  = make([]models.WeatherCloud, 0)
	ElectricFieldSlice = make([]*models.ElectricField, 0)
)

type Log struct {
	TimeStamp time.Time
	Lightning bool
	Distance  uint8
}

func Worker(id int, jobs <-chan string, results chan<- []*models.ElectricField, processType string) {
	for job := range jobs {
		record, _ := ProcessElectricFieldMonitorFile(job)
		fmt.Println(job)
		results <- record
	}
}

func ProcessFiles(logFiles, efmFiles, wcFiles []string) {

	// log File
	logRecords, _ := ProcessLogFile(logFiles[0])

	// Electric Field Monitor Files
	nWorkers := 8
	jobs := make(chan string, len(efmFiles))
	results := make(chan []*models.ElectricField, len(efmFiles))
	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results, "ElectricField")
	}
	for _, efmFile := range efmFiles {
		jobs <- efmFile
	}
	close(jobs)
	for i := 0; i < len(efmFiles); i++ {
		ElectricFieldSlice = append(ElectricFieldSlice, <-results...)
	}
	close(results)

	// Join log records with electric field
	for _, efRecord := range ElectricFieldSlice {
		if logRecord, ok := logRecords[efRecord.TimeStamp]; ok {
			efRecord.Lightning = logRecord.Lightning
			efRecord.Distance = logRecord.Distance
		}
	}

	// Weathercloud files
	cwc := make(chan []models.WeatherCloud, len(wcFiles))
	for _, wcFile := range wcFiles {
		go func(path string, c chan<- []models.WeatherCloud) {
			record, _ := ProcessWeatherCloudFile(path)
			c <- record
		}(wcFile, cwc)
	}
	for i := 0; i < len(wcFiles); i++ {
		WeatherCloudSlice = append(WeatherCloudSlice, <-cwc...)
	}
	close(cwc)
}
