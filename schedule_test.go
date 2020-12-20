package chronos

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSchedule_String(t *testing.T) {
	tt := []struct {
		name     string
		schedule Schedule
		want     string
	}{
		{name: "any", schedule: Schedule{}, want: "* * * * *"},
		{name: "minute", schedule: NewScheduleBuilder().WithMinute(1).Build(), want: "1 * * * *"},
		{name: "hour", schedule: NewScheduleBuilder().WithHour(1).Build(), want: "* 1 * * *"},
		{name: "day of month", schedule: NewScheduleBuilder().WithDayOfMonth(1).Build(), want: "* * 1 * *"},
		{name: "month", schedule: NewScheduleBuilder().WithMonth(1).Build(), want: "* * * 1 *"},
		{name: "day of week", schedule: NewScheduleBuilder().WithDayOfWeek(1).Build(), want: "* * * * 1"},
	}
	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.schedule.String(), test.want)
		})
	}
}
