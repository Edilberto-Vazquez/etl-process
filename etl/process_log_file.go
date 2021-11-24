package etl

import (
	"bufio"
	"os"
	"time"
)

func ProcessLogFile(path string) (map[time.Time]Log, error) {
	records := make(map[time.Time]Log, 0)
	file, err := os.Open(path)
	if err != nil {
		return records, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if ThereIsLightning(scanner.Text()) {
			record := Log{
				TimeStamp: NewTimeStamp("log", scanner.Text()),
				Lightning: NewLightning(),
				Distance:  NewDistance(scanner.Text()),
			}
			records[record.TimeStamp] = record
		}
	}
	return records, file.Close()
}
