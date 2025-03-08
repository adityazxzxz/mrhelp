package mrtime

import (
	"fmt"
	"time"
)

var monthEN = []string{
	"January", "February", "March", "April", "May", "June",
	"July", "August", "September", "October", "November", "December",
}

var monthIDN = []string{
	"Januari", "Februari", "Maret", "April", "Mei", "Juni",
	"Juli", "Agustus", "September", "Oktober", "November", "Desember",
}

func DateTime(input interface{}) (string, error) {
	layout := "2006-01-02 15:04:05"
	res, err := converter(input, layout)
	return res, err
}

func Date(input interface{}) (string, error) {
	layout := "2006-01-02"
	res, err := converter(input, layout)
	return res, err
}

func FullDateEN(input interface{}) (string, error) {
	layout := "2006-01-02"
	res, err := converter(input, layout)
	if err != nil {
		return res, err
	}
	t, _ := time.Parse("2006-01-02", res)
	day := t.Day()
	month := monthEN[t.Month()-1]
	year := t.Year()
	return fmt.Sprintf("%d %s %d", day, month, year), nil
}

func FullDateIDN(input interface{}) (string, error) {
	layout := "2006-01-02"
	res, err := converter(input, layout)
	if err != nil {
		return res, err
	}
	t, _ := time.Parse("2006-01-02", res)
	day := t.Day()
	month := monthIDN[t.Month()-1]
	year := t.Year()
	return fmt.Sprintf("%d %s %d", day, month, year), nil
}

func CurrentUnixTime() int64 {
	return time.Now().Unix()
}

func CurrentTimeString() string {
	return time.Now().Format("15:04:05")
}

func converter(input interface{}, layout string) (string, error) {
	switch v := input.(type) {
	case time.Time:
		return v.Format(layout), nil
	case string:
		parsedTime, err := time.Parse(layout, v)
		if err != nil {
			return "", err
		}
		return parsedTime.Format(layout), nil
	case int64:
		return time.Unix(v, 0).Format(layout), nil
	case float64:
		return time.Unix(int64(v), 0).Format(layout), nil
	default:
		return "", fmt.Errorf("unsupport type")
	}
}
