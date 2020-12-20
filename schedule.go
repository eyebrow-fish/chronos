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
	DayOfWeek *uint8
}

func (s Schedule) String() string {
	return fmt.Sprintf(
		"%s %s %s %s %s",
		cronNum(s.Minute),
		cronNum(s.Hour),
		cronNum(s.DayOfMonth),
		cronNum(s.Month),
		cronNum(s.DayOfWeek),
	)
}

func cronNum(n *uint8) string {
	if n == nil {
		return "*"
	} else {
		return strconv.Itoa(int(*n))
	}
}

type ScheduleBuilder struct {
	schedule Schedule
}

func NewScheduleBuilder() ScheduleBuilder {
	return ScheduleBuilder{}
}

func (s ScheduleBuilder) WithMinute(minute uint8) ScheduleBuilder {
	s.schedule.Minute = &minute
	return s
}

func (s ScheduleBuilder) WithHour(hour uint8) ScheduleBuilder {
	s.schedule.Hour = &hour
	return s
}

func (s ScheduleBuilder) WithDayOfMonth(dayOfMonth uint8) ScheduleBuilder {
	s.schedule.DayOfMonth = &dayOfMonth
	return s
}

func (s ScheduleBuilder) WithMonth(month uint8) ScheduleBuilder {
	s.schedule.Month = &month
	return s
}

func (s ScheduleBuilder) WithDayOfWeek(dayOfWeek uint8) ScheduleBuilder {
	s.schedule.DayOfWeek = &dayOfWeek
	return s
}

func (s ScheduleBuilder) Build() Schedule {
	return s.schedule
}
