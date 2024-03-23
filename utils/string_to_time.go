package utils

import (
	"time"
)

func StringToTime(date string) (time.Time, error) {
	// stringDate := "07.26.2020" 01.02.2006
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return t, err
	}

	return t, err
}
