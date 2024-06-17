package main

import (
	"fmt"
	"time"
)

func formatTime(t time.Time) string {

	// Convert 4 digit year to 2
	formattedYear := fmt.Sprintf("%d", t.Year())[2:]

	month := t.Month().String()[0:3]

	day := t.Day()

	// Zero prefix single digit day
	formattedDay := fmt.Sprintf("%02d", day)

	return month + formattedDay + formattedYear
}
