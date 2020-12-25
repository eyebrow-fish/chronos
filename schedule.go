package chronos

import (
	"strconv"
	"time"
)

type schedule struct {
	unit  ScheduleUnit
	value int
}

func (s schedule) nextDate(from time.Time) time.Time {
	dur := time.Duration(s.value)
	secNoise := -(time.Duration(from.Second())*time.Second + time.Duration(from.Nanosecond()))
	minNoise := -(time.Duration(from.Minute()) * time.Minute)
	hourNoise := -(time.Duration(from.Hour()) * time.Hour)
	switch s.unit {
	case minute:
		return from.Add(dur*time.Minute + secNoise)
	case hour:
		return from.Add(dur*time.Hour + minNoise + secNoise)
	case day:
		return from.Add(hourNoise+minNoise+secNoise).
			AddDate(0, 0, s.value)
	case month:
		return from.AddDate(0, s.value, -from.Day()+1).
			Add(hourNoise + minNoise + secNoise)
	case year:
		return from.AddDate(s.value, int(-from.Month()+1), -from.Day()+1).
			Add(hourNoise + minNoise + secNoise)
	default:
		panic("unknown unit: " + strconv.Itoa(int(s.unit)))
	}
}

type ScheduleUnit int

const (
	minute ScheduleUnit = iota
	hour
	day
	month
	year
)
