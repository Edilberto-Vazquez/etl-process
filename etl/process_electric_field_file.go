package etl

import (
	"bufio"
	"os"
	"strings"
	"time"

	"github.com/Edilberto-Vazquez/weathercloud-etl-process/db/models"
)

func ProcessElectricFieldMonitorFile(path string) ([]*models.ElectricField, error) {
	records := make([]*models.ElectricField, 0)
	file, err := os.Open(path)
	if err != nil {
		return records, err
	}
	scanner := bufio.NewScanner(file)
	var electricFields []string = make([]string, 0)
	var timeStamp string
	var rotorStatus string
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")
		if timeStamp == "" || timeStamp == fields[0] {
			electricFields = append(electricFields, fields[1])
			timeStamp = fields[0]
			rotorStatus = fields[2]
		} else {
			date := NewTimeStamp("efm", path+" "+timeStamp)
			records = append(records, &models.ElectricField{
				TimeStamp:      date,
				TimeStampRound: date.Round(10 * time.Minute),
				ElectricField:  ElectricFieldAvg(electricFields),
				RotorStatus:    NewRotorStatus(rotorStatus),
			})
			electricFields = make([]string, 0)
			electricFields = append(electricFields, fields[1])
			timeStamp = fields[0]
		}
	}
	return records, file.Close()
}
