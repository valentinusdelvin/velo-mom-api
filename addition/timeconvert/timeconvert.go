package addition

import (
	"fmt"
	"time"
)

type TimeConvertInterface interface {
	TimeConvert(t time.Time) string
}

func TimeConvert(t time.Time) string {
	hariIndonesia := map[string]string{
		"Sunday":    "Minggu",
		"Monday":    "Senin",
		"Tuesday":   "Selasa",
		"Wednesday": "Rabu",
		"Thursday":  "Kamis",
		"Friday":    "Jumat",
		"Saturday":  "Sabtu",
	}

	bulanIndonesia := map[string]string{
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

	hari := t.Weekday().String()
	bulan := t.Month().String()

	return fmt.Sprintf("%s, %02d %s %d", hariIndonesia[hari], t.Day(), bulanIndonesia[bulan], t.Year())
}
