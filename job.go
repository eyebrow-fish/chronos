package chronos

import (
	"context"
	"errors"
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
	return nil, errors.New("unimplemented")
}

type cronUnit struct {
	unitType cronUnitType
	values   []uint8
}

type cronUnitType uint8

const (
	every cronUnitType = iota
	exact
	divisable
)
