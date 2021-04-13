package chronos

import (
	"context"
	"errors"
	"strings"
	"time"
)

type JobFunc func(ctx context.Context) error

func Job(name, cron string, f JobFunc) error {
	return errors.New("unimplemented")
}

type cronSchedule struct {
	minute   cronUnit
	hour     cronUnit
	monthDay cronUnit
	month    cronUnit
	weekDay  cronUnit
}

func scheduleFromString(cron string) (*cronSchedule, error) {
	cronUnitTokens := strings.Split(cron, " ")

	if tokenCount := len(cronUnitTokens); tokenCount != 5 {
		if tokenCount > 5 {
			return nil, errors.New("too many time unit tokens")
		} else {
			return nil, errors.New("too few time unit tokens")
		}
	}

	var units []cronUnit
	for _, token := range cronUnitTokens {
		if unit, err := parseUnitToken(token); err == nil {
			units = append(units, *unit)
		} else {
			return nil, err
		}
	}

	//goland:noinspection GoNilness
	schedule := cronSchedule{
		units[0],
		units[1],
		units[2],
		units[3],
		units[4],
	}

	return &schedule, nil
}

func (cs cronSchedule) nextTime(from time.Time) time.Time {
	rounded := from.Round(time.Minute)

	toMinute := adjustForUnit(rounded, cs.minute, rounded.Minute(), time.Minute, time.Hour)
	toHour := adjustForUnit(toMinute, cs.hour, toMinute.Hour(), time.Hour, time.Hour * 24)

	bumped := toHour
	if cs.minute.unitType == every && bumped.Equal(rounded) {
		bumped = rounded.Add(time.Minute)
	}

	return bumped
}

func adjustForUnit(from time.Time, unit cronUnit, fromValue int, timeUnit, greaterTimeUnit time.Duration) time.Time {
	to := from

	switch unit.unitType {
	case listed:
		next := unit.values[0]
		for _, i := range unit.values {
			if fromValue < i {
				next = i
				break
			}
		}

		delta := time.Duration(next-fromValue) * timeUnit

		if delta > 0 {
			to = to.Add(delta)
		} else {
			to = to.Add(greaterTimeUnit + delta)
		}
	case ranged:
		lowerDelta := time.Duration(unit.values[0]-fromValue) * timeUnit

		if lowerDelta > 0 {
			to = to.Add(lowerDelta)
		} else if fromValue < unit.values[1] {
			to = to.Add(timeUnit)
		} else {
			to = to.Add(greaterTimeUnit + lowerDelta)
		}
	case stepped:
		step := unit.values[0]

		until := step - fromValue%step
		if until == 0 {
			to = to.Add(time.Duration(step) * timeUnit)
		} else {
			to = to.Add(time.Duration(until) * timeUnit)
		}
	}

	return to
}

type cronUnit struct {
	unitType cronUnitType
	values   []int
}

type cronUnitType uint8

const (
	every cronUnitType = iota
	listed
	ranged
	stepped
)
