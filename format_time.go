package main

import (
	"fmt"
	"time"
)

func format_time() string {
	currentTime := time.Now()
	formattedYear := fmt.Sprintf("%d", currentTime.Year())
	month := currentTime.Month()
	formattedMonth := fmt.Sprintf("%02d", month)
	day := currentTime.Day()
	formattedDay := fmt.Sprintf("%02d", day)

	// formatted := fmt.Sprintf("%s%s%s", month, day, formattedYear)

	return formattedMonth + formattedDay + formattedYear
}
