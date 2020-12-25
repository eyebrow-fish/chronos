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
	nsNoise := -time.Duration(from.Nanosecond())
	secNoise := -time.Duration(from.Second()) * time.Second
	minNoise := -(time.Duration(from.Minute()) * time.Minute)
	hourNoise := -(time.Duration(from.Hour()) * time.Hour)
	switch s.unit {
	case second:
		return from.Add(dur*time.Second + nsNoise)
	case minute:
		return from.Add(dur*time.Minute + nsNoise + secNoise)
	case hour:
		return from.Add(dur*time.Hour + nsNoise + secNoise + minNoise)
	case day:
		return from.AddDate(0, 0, s.value).
			Add(nsNoise + secNoise + minNoise + hourNoise)
	case month:
		return from.AddDate(0, s.value, -from.Day()+1).
			Add(nsNoise + secNoise + minNoise + hourNoise)
	case year:
		return from.AddDate(s.value, int(-from.Month()+1), -from.Day()+1).
			Add(nsNoise + secNoise + minNoise + hourNoise)
	default:
		panic("unknown unit: " + strconv.Itoa(int(s.unit)))
	}
}

type ScheduleUnit int

const (
	second ScheduleUnit = iota
	minute
	hour
	day
	month
	year
)
