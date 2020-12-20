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
				return &ScheduleUnit{UnitType: Recur, Value: []uint8{uint8(v)}}, nil
			}
		} else {
			var values []uint8
			for _, v := range strings.Split(s, ",") {
				if i, err := strconv.Atoi(v); err != nil {
					return nil, err
				} else {
					values = append(values, uint8(i))
				}
			}
			var ut UnitType
			if len(values) > 1 {
				ut = Multi
			} else {
				ut = Exact
			}
			return &ScheduleUnit{UnitType: ut, Value: values}, nil
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
		Minute:     *minute,
		Hour:       *hour,
		DayOfMonth: *dayOfMonth,
		Month:      *month,
		DayOfWeek:  *dayOfWeek,
	}, nil
}

func (s Schedule) String() string {
	cronVal := func(u ScheduleUnit) string {
		if u.Value == nil {
			return "*"
		} else if u.UnitType != Recur {
			var ss []string
			for _, v := range u.Value {
				ss = append(ss, strconv.Itoa(int(v)))
			}
			return strings.Join(ss, ",")
		} else {
			return "*/" + strconv.Itoa(int(u.Value[0]))
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
	UnitType UnitType
	Value    []uint8
}

func ExactUnit(x uint8) ScheduleUnit {
	return ScheduleUnit{UnitType: Exact, Value: []uint8{x}}
}

func MultiUnit(x ...uint8) ScheduleUnit {
	return ScheduleUnit{UnitType: Multi, Value: x}
}

func RecurUnit(x uint8) ScheduleUnit {
	return ScheduleUnit{UnitType: Recur, Value: []uint8{x}}
}

func NoneUnit() ScheduleUnit {
	return ScheduleUnit{}
}

type UnitType uint8

const (
	Exact UnitType = iota
	Recur
	Multi
)
