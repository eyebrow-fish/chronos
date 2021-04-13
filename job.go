package chronos

import (
	"context"
	"errors"
	"strings"
	"time"
)

type JobFunc func(ctx context.Context) error

func Job(name, cron string, f JobFunc) error {
	return errors.New("unimplemeneted")
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
	to := from.Round(time.Minute)

	switch cs.minute.unitType {
	case every:
		to = to.Add(time.Minute)
	case listed:
		next := cs.minute.values[0]
		for _, i := range cs.minute.values {
			if to.Minute() < i {
				next = i
				break
			}
		}

		current := to.Minute()
		delta := time.Duration(next-current) * time.Minute

		if delta > 0 {
			to = to.Add(delta)
		} else {
			to = to.Add(time.Hour + delta)
		}
	case ranged:
		current := to.Minute()

		lowerDelta := time.Duration(cs.minute.values[0]-current) * time.Minute

		if lowerDelta > 0 {
			to = to.Add(lowerDelta)
		} else if to.Minute() < cs.minute.values[1] {
			to = to.Add(time.Minute)
		} else {
			to = to.Add(time.Hour + lowerDelta)
		}
	case stepped:
		step := cs.minute.values[0]

		until := step - to.Minute()%step
		if until == 0 {
			to = to.Add(time.Duration(step) * time.Minute)
		} else {
			to = to.Add(time.Duration(until) * time.Minute)
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
