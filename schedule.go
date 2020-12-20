package chronos

import (
	"fmt"
	"strconv"
	"strings"
)

type Schedule struct {
	Minute,
	Hour,
	DayOfMonth,
	Month,
	DayOfWeek ScheduleUnit
}

func NewSchedule(cron string) (*Schedule, error) {
	cronUnit := func(s string) (*ScheduleUnit, error) {
		if s == "*" {
			return &ScheduleUnit{}, nil
		} else if strings.Contains(s, "/") {
			num := s[2:]
			if v, err := strconv.Atoi(num); err != nil {
				return nil, err
			} else {
				cast := uint8(v)
				return &ScheduleUnit{Value: &cast}, nil
			}
		} else {
			if v, err := strconv.Atoi(s); err != nil {
				return nil, err
			} else {
				cast := uint8(v)
				return &ScheduleUnit{IsExact: true, Value: &cast}, nil
			}
		}
	}
	units := strings.Split(cron, " ")
	minute, err := cronUnit(units[0])
	if err != nil {
		return nil, err
	}
	hour, err := cronUnit(units[1])
	if err != nil {
		return nil, err
	}
	dayOfMonth, err := cronUnit(units[2])
	if err != nil {
		return nil, err
	}
	month, err := cronUnit(units[3])
	if err != nil {
		return nil, err
	}
	dayOfWeek, err := cronUnit(units[4])
	if err != nil {
		return nil, err
	}
	return &Schedule{
		Minute: *minute,
		Hour: *hour,
		DayOfMonth: *dayOfMonth,
		Month: *month,
		DayOfWeek: *dayOfWeek,
	}, nil
}

func (s Schedule) String() string {
	cronVal := func(u ScheduleUnit) string {
		if u.Value == nil {
			return "*"
		} else if u.IsExact {
			return strconv.Itoa(int(*u.Value))
		} else {
			return "*/" + strconv.Itoa(int(*u.Value))
		}
	}
	return fmt.Sprintf(
		"%s %s %s %s %s",
		cronVal(s.Minute),
		cronVal(s.Hour),
		cronVal(s.DayOfMonth),
		cronVal(s.Month),
		cronVal(s.DayOfWeek),
	)
}

type ScheduleUnit struct {
	IsExact bool
	Value   *uint8
}

func ExactUnit(x uint8) ScheduleUnit {
	return ScheduleUnit{IsExact: true, Value: &x}
}

func RecurUnit(x uint8) ScheduleUnit {
	return ScheduleUnit{Value: &x}
}

func NoneUnit() ScheduleUnit {
	return ScheduleUnit{}
}
