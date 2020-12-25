package chronos

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestJob_NextRun(t *testing.T) {
	tests := []struct {
		name string
		job  Job
		want time.Time
	}{
		{
			name: "minutely",
			job:  JobFrom(time.Date(2020, 12, 25, 0, 0, 0, 0, time.FixedZone("UTC+0", 0)), nil).Minutely(),
			want: time.Date(2020, 12, 25, 0, 1, 0, 0, time.FixedZone("UTC+0", 0)),
		},
		{
			name: "hourly",
			job:  JobFrom(time.Date(2020, 12, 25, 0, 30, 0, 0, time.FixedZone("UTC+0", 0)), nil).Hourly(),
			want: time.Date(2020, 12, 25, 1, 0, 0, 0, time.FixedZone("UTC+0", 0)),
		},
		{
			name: "daily",
			job:  JobFrom(time.Date(2020, 12, 25, 5, 30, 0, 0, time.FixedZone("UTC+0", 0)), nil).Daily(),
			want: time.Date(2020, 12, 26, 0, 0, 0, 0, time.FixedZone("UTC+0", 0)),
		},
		{
			name: "monthly",
			job:  JobFrom(time.Date(2020, 12, 25, 0, 0, 0, 0, time.FixedZone("UTC+0", 0)), nil).Monthly(),
			want: time.Date(2021, 1, 1, 0, 0, 0, 0, time.FixedZone("UTC+0", 0)),
		},
		{
			name: "yearly",
			job:  JobFrom(time.Date(2020, 12, 25, 0, 0, 0, 0, time.FixedZone("UTC+0", 0)), nil).Yearly(),
			want: time.Date(2021, 1, 1, 0, 0, 0, 0, time.FixedZone("UTC+0", 0)),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := test.job.NextRun()
			assert.Equal(t, out, test.want)
		})
	}
}
