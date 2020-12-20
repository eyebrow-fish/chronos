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
		{name: "minute 1 and 2", schedule: Schedule{Minute: MultiUnit(1, 2)}, want: "1,2 * * * *"},
	}
	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.schedule.String(), test.want)
		})
	}
}

func TestNewSchedule(t *testing.T) {
	tests := []struct {
		name    string
		cron    string
		want    *Schedule
		wantErr bool
	}{
		{name: "any", cron: "* * * * *", want: &Schedule{}},
		{name: "exact minute", cron: "10 * * * *", want: &Schedule{Minute: ExactUnit(10)}},
		{name: "recurring minute", cron: "*/10 * * * *", want: &Schedule{Minute: RecurUnit(10)}},
		{name: "multi minute", cron: "1,2,3 * * * *", want: &Schedule{Minute: MultiUnit(1, 2, 3)}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out, err := NewSchedule(test.cron)
			assert.Equal(t, out, test.want)
			assert.Equal(t, err != nil, test.wantErr)
		})
	}
}
