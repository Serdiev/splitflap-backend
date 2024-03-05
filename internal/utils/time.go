package utils

import (
	"splitflap-backend/internal/logger"
	"time"
)

func SetTimeZone() {
	tz, err := time.LoadLocation(cfg.General.TimeZone)
	if err != nil {
		logger.Error().Err(err).Msg("Error loading time zone")
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
