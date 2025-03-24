package shared

import (
	"strings"
	"time"
)

var (
	weekdays = map[string]int{
		"sunday":    0,
		"monday":    1,
		"tuesday":   2,
		"wednesday": 3,
		"thursday":  4,
		"friday":    5,
		"saturday":  6,
	}
)

func CalculateDay(day string) string {
	day = strings.ToLower(strings.TrimSpace(day))

	_, exists := weekdays[day]

	if exists {
		return CapitalizeFirstLetter((day))
	}

	currentDay := time.Now().Weekday().String()

	return currentDay
}

func CalculateDayDifference(currentDay, startDay string) int {
	currentDayInt := weekdays[strings.ToLower(currentDay)]
	startDayInt := weekdays[strings.ToLower(startDay)]
	dayDifference := (startDayInt - currentDayInt + 7) % 7

	return dayDifference
}
