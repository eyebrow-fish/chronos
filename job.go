package chronos

import (
	"context"
	"errors"
	"fmt"
	"strings"
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

	schedule := cronSchedule{
		units[0],
		units[1],
		units[2],
		units[3],
		units[4],
	}

	return &schedule, nil
}

func parseUnitToken(cronToken string) (*cronUnit, error) {
	if cronToken == "*" {
		return &cronUnit{}, nil
	}

	return nil, fmt.Errorf("unknown token: %s", cronToken)
}

type cronUnit struct {
	unitType cronUnitType
	values   []uint8
}

type cronUnitType uint8

const (
	every cronUnitType = iota
	listed
	ranged
	stepped
)
