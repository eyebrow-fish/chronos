package chronos

import (
	"fmt"
	"strconv"
)

type Schedule struct {
	Minute,
	Hour,
	DayOfMonth,
	Month,
	DayOfWeek ScheduleUnit
}

func (s Schedule) String() string {
	cronVal := func(u ScheduleUnit) string {
		if u.value == nil {
			return "*"
		} else if u.IsExact {
			return strconv.Itoa(int(*u.value))
		} else {
			return "*/" + strconv.Itoa(int(*u.value))
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
	value *uint8
}

func ExactUnit(x uint8) ScheduleUnit {
	return ScheduleUnit{IsExact: true, value: &x}
}

func RecurUnit(x uint8) ScheduleUnit {
	return ScheduleUnit{value: &x}
}

func NoneUnit() ScheduleUnit {
	return ScheduleUnit{}
}
