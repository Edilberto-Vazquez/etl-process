package etl

import (
	"bufio"
	"encoding/binary"
	"os"
	"strings"

	"github.com/Edilberto-Vazquez/weathercloud-etl-process/db/models"
)

func ProcessWeatherCloudFile(path string) ([]models.WeatherCloud, error) {
	records := make([]models.WeatherCloud, 0)
	file, err := os.Open(path)
	if err != nil {
		return records, err
	}
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		utf16, _ := DecodeUtf16(scanner.Bytes(), binary.BigEndian)
		fields := strings.Split(utf16, ";")
		if len(scanner.Text()) == 1 {
			continue
		}
		if i == 0 {
			i++
			continue
		}
		records = append(records, models.WeatherCloud{
			TimeStamp: NewTimeStamp("wc", fields[0]),
			TempIn:    CommaToPoint(fields[1]),
			Temp:      CommaToPoint(fields[2]),
			Chill:     CommaToPoint(fields[3]),
			DewIn:     CommaToPoint(fields[4]),
			Dew:       CommaToPoint(fields[5]),
			HeatIn:    CommaToPoint(fields[6]),
			Heat:      CommaToPoint(fields[7]),
			Humin:     CommaToPoint(fields[8]),
			Hum:       CommaToPoint(fields[9]),
			Wspdhi:    CommaToPoint(fields[10]),
			Wspdavg:   CommaToPoint(fields[11]),
			Wdiravg:   CommaToPoint(fields[12]),
			Bar:       CommaToPoint(fields[13]),
			Rain:      CommaToPoint(fields[14]),
			RainRate:  CommaToPoint(fields[15]),
		})
	}
	return records, file.Close()
}
