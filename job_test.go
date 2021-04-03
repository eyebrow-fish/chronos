package chronos

import (
	"reflect"
	"testing"
)

func Test_scheduleFromString(t *testing.T) {
	type args struct {
		cron string
	}

	tests := []struct {
		name    string
		args    args
		want    *cronSchedule
		wantErr bool
	}{
		{
			"every minute",
			args{"* * * * *"},
			&cronSchedule{
				minute:   cronUnit{},
				hour:     cronUnit{},
				monthDay: cronUnit{},
				month:    cronUnit{},
				weekDay:  cronUnit{},
			},
			false,
		},

		{
			"at minute 10",
			args{"10 * * * *"},
			&cronSchedule{
				minute:   cronUnit{listed, []uint8{10}},
				hour:     cronUnit{},
				monthDay: cronUnit{},
				month:    cronUnit{},
				weekDay:  cronUnit{},
			},
			false,
		},
		{
			"at minute 10 on day-of-month 10",
			args{"10 * 10 * *"},
			&cronSchedule{
				minute:   cronUnit{listed, []uint8{10}},
				hour:     cronUnit{},
				monthDay: cronUnit{listed, []uint8{10}},
				month:    cronUnit{},
				weekDay:  cronUnit{},
			},
			false,
		},

		{
			"at minute 10, 15, and 20",
			args{"10,15,20 * * * *"},
			&cronSchedule{
				minute:   cronUnit{listed, []uint8{10, 15, 20}},
				hour:     cronUnit{},
				monthDay: cronUnit{},
				month:    cronUnit{},
				weekDay:  cronUnit{},
			},
			false,
		},
		{
			"at minute 10, 15, and 20 on day-of-week Monday and Tuesday",
			args{"10,15,20 * * * 1,2"},
			&cronSchedule{
				minute:   cronUnit{listed, []uint8{10, 15, 20}},
				hour:     cronUnit{},
				monthDay: cronUnit{},
				month:    cronUnit{},
				weekDay:  cronUnit{listed, []uint8{1, 2}},
			},
			false,
		},

		{"too many tokens", args{"* * * * * *"}, nil, true},
		{"trailing comma", args{"1, * * * *"}, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := scheduleFromString(tt.args.cron)
			if (err != nil) != tt.wantErr {
				t.Errorf("scheduleFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scheduleFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
