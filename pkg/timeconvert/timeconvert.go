package timeconvert

import (
	"fmt"
	"time"
)

type TimeConvertInterface interface {
	TimeConvert(t time.Time) string
}

func TimeConvert(t time.Time) string {
	indonesianDay := map[string]string{
		"Sunday":    "Minggu",
		"Monday":    "Senin",
		"Tuesday":   "Selasa",
		"Wednesday": "Rabu",
		"Thursday":  "Kamis",
		"Friday":    "Jumat",
		"Saturday":  "Sabtu",
	}

	indonesianMonth := map[string]string{
		"January":   "Januari",
		"February":  "Februari",
		"March":     "Maret",
		"April":     "April",
		"May":       "Mei",
		"June":      "Juni",
		"July":      "Juli",
		"August":    "Agustus",
		"September": "September",
		"October":   "Oktober",
		"November":  "November",
		"December":  "Desember",
	}

	day := t.Weekday().String()
	month := t.Month().String()

	return fmt.Sprintf("%s, %02d %s %d", indonesianDay[day], t.Day(), indonesianMonth[month], t.Year())
}
