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
		{name: "minute", schedule: Schedule{Minute: ExactUnit(1)}, want: "1 * * * *"},
		{name: "hour", schedule: Schedule{Hour: ExactUnit(1)}, want: "* 1 * * *"},
		{name: "day of month", schedule: Schedule{DayOfMonth: ExactUnit(1)}, want: "* * 1 * *"},
		{name: "month", schedule: Schedule{Month: ExactUnit(1)}, want: "* * * 1 *"},
		{name: "day of week", schedule: Schedule{DayOfWeek: ExactUnit(1)}, want: "* * * * 1"},
		{
			name: "any explicit",
			schedule: Schedule{
				Minute:     NoneUnit(),
				Hour:       NoneUnit(),
				DayOfMonth: NoneUnit(),
				Month:      NoneUnit(),
				DayOfWeek:  NoneUnit(),
			},
			want: "* * * * *",
		},
		{name: "every 10 minutes", schedule: Schedule{Minute: RecurUnit(10)}, want: "*/10 * * * *"},
	}
	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.schedule.String(), test.want)
		})
	}
}
