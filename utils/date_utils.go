package main

import (
	"time"
)

const DATE_FORMAT_DEFAULT = "2006-01-02 15:04:05"
const DAY_FORMAT_DEFAULT = "2006-01-02"

func FormatTimeByTime(t time.Time) string {
	return t.Format(DATE_FORMAT_DEFAULT)
}

func FormatDayByTime(t time.Time) string {
	return t.Format(DAY_FORMAT_DEFAULT)
}

func FormatTimeByMs(t int64) string {
	time := time.Unix(t/1000, t%1000*1000000)
	return FormatTimeByTime(time)
}

func FormatDayByMs(t int64) string {
	time := time.Unix(t/1000, t%1000*1000000)
	return FormatDayByTime(time)
}

func ParseTimeByDay(s string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local")
	return time.ParseInLocation(DAY_FORMAT_DEFAULT, s, loc)
}

func ParseTimeByTime(s string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local")
	return time.ParseInLocation(DATE_FORMAT_DEFAULT, s, loc)
}
