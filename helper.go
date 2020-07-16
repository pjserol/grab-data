package main

import (
	"fmt"
	"time"
)

func convertTime(input string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05Z", input)
}

// getFilePath when hour block is empty default value '0'
// e.g format generated '/cors/rinex/2017/257/nybp/nybp257b.17o.gz'
// hoursBlock: a refers to the hour from 12am-1am (GPS time), x refers to 11pm-12am, and so on
func getFilePath(baseStationID, hourBlock string, t time.Time) string {
	year := t.Year()
	yearDay := t.YearDay()

	if hourBlock == "" {
		hourBlock = "0"
	}

	return fmt.Sprintf("/cors/rinex/%d/%03d/%s/%s%03d%s.%do.gz", year, yearDay, baseStationID, baseStationID, yearDay, hourBlock, lastTwoDigits(year))
}

// TODO: add unit test
func lastTwoDigits(num int) int {
	return num % 100
}
