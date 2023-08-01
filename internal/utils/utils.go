package utils

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/golang-module/carbon/v2"
)

const (
	YMD_HM = "2006/01/02 15:04"
	YMD    = "2006-01-02"
)

func GetTimeUTC7() time.Time {
	now := time.Now()
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	return now.In(loc)
}
func GenTaskId() string {
	now := carbon.Now().Format("ymdHisu")
	prefix := "FC"
	if !internal.Envs.IsProduction {
		prefix = "S" + prefix
	}
	return prefix + now
}

func IsSubSlice(parent []string, subSlice []string) bool {

	for _, c := range subSlice {
		isFound := false
		for _, p := range parent {
			if c == p {
				isFound = true
				break
			}
		}
		if !isFound {
			return false
		}
	}
	return true
}
func AppointmentToDate(date string, timeSlot string) (firstTimeSlot carbon.Carbon, secondTimeSlot carbon.Carbon) {
	hoursMinutes := strings.Split(timeSlot, " - ")
	if len(hoursMinutes) == 2 {
		format := "Y-m-dH:i"
		firstTimeSlot = carbon.ParseByFormat(date+hoursMinutes[0], format)
		secondTimeSlot = carbon.ParseByFormat(date+hoursMinutes[1], format)
	}
	return
}
func CreateTokenLocation(clientKey, secretKey string) string {
	timeString := carbon.Now().Format("Y-d-m")
	return clientKey + "::" + GetMD5Hash(clientKey+"::"+secretKey+timeString)
}
func I2Str(value interface{}) string {
	if value == nil {
		return ""
	}
	return fmt.Sprint(value)
}

func StringToFormatTime(stringType string) string {
	switch stringType {
	case "Y-D-M":
		return "2006-02-01"
	case "Y-M-D":
		return "2006-01-02"
	case "D-M-Y":
		return "02-01-2006"
	case "M-D-Y":
		return "01-02-2006"
	case "Y-D-M H:M:S":
		return "2006-02-01 15:04:05"
	case "Y-M-D H:M:S":
		return "2006-01-02 15:04:05"
	case "Y-M-D H:M:S -0700":
		return "2006-01-02 15:04:05 -0700"
	case "D/M/Y":
		return "02/01/2006"
	case "D/M/Y H:M":
		return "02/01/2006 15:04"
	case "H:M D/M/Y":
		return "15:04 02/01/2006"
	case "Y-M-DTH:M:S.000":
		return "2006-01-02T15:04:05.999999999"
	case "D/M/Y H:M:S.000":
		return "02-01-2006 15:04:05.999999999"
	case "D/M/Y H:M:S":
		return "02/01/2006 15:04:05"
	case "D/M/Y - H:M":
		return "02/01/2006 - 15:04"
	}
	return ""
}

func GetTimeUTC7FrTime(input time.Time) time.Time {
	if input.Location().String() == "UTC" {
		input = input.Add(time.Hour * -7)
	}
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	return input.In(loc)
}

func ParseTimeFrString(stringType string, timeInput string) (time.Time, error) {
	layout := StringToFormatTime(stringType)
	return time.Parse(layout, timeInput)
}

func ParseTimeFrStringV2(stringType string, timeInput string) (time.Time, error) {
	t, err := ParseTimeFrString(stringType, timeInput)
	if err != nil {
		return t, err
	}
	return GetTimeUTC7FrTime(t), nil
}

func ConvertStringToList(input string, key string) []string {
	if input == "" {
		return []string{}
	}
	return strings.Split(input, key)
}

func IsEmpty(value string) bool {
	return len(strings.TrimSpace(value)) == 0
}
func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func CreateLocalEcomToken(EcomClientKey, EcomSecretKey string) string {
	timestr := GetTimeUTC7().Format("2006-02-01")
	return GetMD5Hash(EcomClientKey + "::" + EcomSecretKey + timestr)
}

func ConvertTimeStringFromAtoB(inputTimeFormat, inputTime, toTimeFormat string) (string, error) {
	/// formatTime is like RFC3390, "2006-01-02" ...
	timeParse, err := time.Parse(inputTimeFormat, inputTime)
	if err != nil {
		return "", err
	}
	parsedTime := timeParse.Format(toTimeFormat)
	return parsedTime, nil
}

func ConvertMinuteToString(inputMinute int) string {
	if inputMinute <= 0 {
		return "0"
	}
	hour := int(inputMinute / 60)
	subMinute := inputMinute - (hour * 60)
	if hour == 0 {
		return fmt.Sprintf("%vp", subMinute)
	}
	return fmt.Sprintf("%vh%vp", hour, subMinute)
}
