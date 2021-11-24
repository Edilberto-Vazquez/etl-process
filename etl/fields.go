package etl

import (
	"bufio"
	"bytes"
	"database/sql"
	"encoding/binary"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf16"
)

// -----open a file-----
func OpenFile(path string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(path)
	scanner := bufio.NewScanner(file)
	if err != nil {
		log.Println(err)
	}
	return file, scanner
}

// -----chech if the log contains a lightning
func ThereIsLightning(str string) bool {
	ok := regexList["lightning"].MatchString(str)
	if !ok {
		return ok
	}
	return ok
}

// -----date-----
func NewTimeStamp(rowType string, str string) (timeStamp time.Time) {
	var date string
	var duration string
	var split []string
	switch rowType {
	case "log":
		date = regexList["date"].FindString(str)
		duration = regexList["duration"].FindString(str)
		split = strings.Split(date, "/")
		timeStamp, _ = time.Parse(time.RFC3339, split[2]+"-"+split[0]+"-"+split[1]+"T"+duration+"Z")
		return
	case "efm":
		split = strings.Split(str, "/")
		split = strings.Split(split[len(split)-1], "-")
		date = split[1]
		duration = regexList["duration"].FindString(str)
		timeStamp, _ = time.Parse(time.RFC3339, date[4:8]+"-"+date[0:2]+"-"+date[2:4]+"T"+duration+"Z")
		return
	case "wc":
		split = strings.Split(str, " ")
		timeStamp, _ = time.Parse(time.RFC3339, split[0]+"T"+split[1]+"Z")
		return
	}
	return
}

// -----place-----
func NewPlace(str string) int {
	re := regexList["place"]
	return places[re.FindString(str)]
}

func NewLightning() bool {
	return true
}

// -----distance-----
func NewDistance(str string) uint8 {
	re := regexList["distance"]
	match := re.FindString(str)
	split := strings.Split(match, " ")
	distance, _ := strconv.ParseInt(split[1], 10, 64)
	return uint8(distance)
}

// calc the average of the electric field
func ElectricFieldAvg(electricFields []string) float32 {
	sum := 0.0
	divisor := float64(len(electricFields))
	for _, value := range electricFields {
		float, _ := strconv.ParseFloat(value, 32)
		sum += float
	}
	return float32((math.Round((sum / divisor * 100)) / 100)) // the function return a float64 if the rounding is with the function Round
	// return fmt.Sprintf("%0.2f", (sum / divisor)) // the function return a string if the roundig is with the function Sprintf

}

// electric field
func NewRotorStatus(rotorStatus string) bool {
	return rotorStatus != "0"
}

// weather cloud
func SplitString(record string) []string {
	return strings.Split(record, ";")
}

// change comma to point
func CommaToPoint(str string) sql.NullFloat64 {
	if len(str) == 0 {
		return sql.NullFloat64{}
	}
	float, _ := strconv.ParseFloat(strings.Replace(str, ",", ".", 1), 32)
	return sql.NullFloat64{
		Float64: float,
		Valid:   true,
	}
}

// -----Convert from UTF-16 to UTF-8
func DecodeUtf16(b []byte, order binary.ByteOrder) (string, error) {
	ints := make([]uint16, len(b)/2)
	if err := binary.Read(bytes.NewReader(b), order, &ints); err != nil {
		return "", err
	}
	return string(utf16.Decode(ints)), nil
}
