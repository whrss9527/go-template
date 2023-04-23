package util

import (
	"fmt"
	"time"
)

func IncreaseNumberDays(time1 time.Time, days int) time.Time {
	return time1.AddDate(0, 0, days)
}
func GetTodayStartTime() time.Time {
	currentTime := time.Now()
	return time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())

}

func TimeValid(month uint32, day uint32) error {
	var timeStr = fmt.Sprintf("2020-%02d-%02d", month, day)
	_, err := time.Parse("2006-01-02", timeStr)
	return err
}
