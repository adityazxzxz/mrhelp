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

var currentTimezone, _ = time.LoadLocation("Asia/Jakarta")

func SetTimezone(tz string) error {
	zone, err := time.LoadLocation(tz)
	if err != nil {
		return err
	}

	currentTimezone = zone
	return nil
}

// Convert datetime yang diinput ke dalam format "YYYY-mm-dd HH:MM:SS"
//
// How to use:
//
//	func main(){
//		t,_ := Date("2025-01-01")
//		fmt.Println(t)
//	}
func DateTime(input interface{}) (string, error) {
	layout := "2006-01-02 15:04:05"
	res, err := converter(input, layout)
	return res, err
}

// Convert tanggal yang diinput menjadi format "2006-01-02"
//
// How to use:
//
//	func main(){
//		t,_ := Date("2025-01-01")
//		fmt.Println(t)
//	}
func Date(input interface{}) (string, error) {
	layout := "2006-01-02"
	res, err := converter(input, layout)
	return res, err
}

// FullDateEN convert waktu yang diinput menjadi format "1 January 2022" dalam bahasa inggris
// berdasarkan timezone yang sudah diatur dengan `SetTimezone`.
//
// How to use:
//
//	func main(){
//		t,_ := FullDateEN("2025-01-01")
//		fmt.Println(t)
//	}
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

// FullDateEN convert waktu yang diinput menjadi format "1 Januari 2022" dalam bahasa indonesia
// berdasarkan timezone yang sudah diatur dengan `SetTimezone`.
//
// Parameters:
//   - input (string,int,dateTime): contoh: 2025-01-02 atau 171233122.
//
// Returns:
//   - string: DateTime seperti "1 January 2022".
//   - error: Mengambalikan error apabila format input tidak sesuai
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
	return time.Now().In(currentTimezone).Unix()
}

func CurrentTimeString() string {
	return time.Now().In(currentTimezone).Format("15:04:05")
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
		return time.Unix(v, 0).In(currentTimezone).Format(layout), nil
	case float64:
		return time.Unix(int64(v), 0).In(currentTimezone).Format(layout), nil
	default:
		return "", fmt.Errorf("unsupport type")
	}
}
