package utils

import (
	"fmt"
	"time"
)

func SetTimeZone() {
	tz, err := time.LoadLocation(cfg.General.TimeZone)
	if err != nil {
		fmt.Println("Error loading time zone:", err)
		return
	}
	time.Local = tz
}

func GetTodayAt(hour int, minute int) time.Time {
	return time.Date(
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
		hour, minute, 0, 0, time.Local)
}
