package chronos

import (
	"reflect"
	"testing"
	"time"
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

		{
			"at every minute 5 through 10",
			args{"5-10 * * * *"},
			&cronSchedule{
				minute:   cronUnit{ranged, []uint8{5, 10}},
				hour:     cronUnit{},
				monthDay: cronUnit{},
				month:    cronUnit{},
				weekDay:  cronUnit{},
			},
			false,
		},

		{
			"at every 10th minutes",
			args{"*/10 * * * *"},
			&cronSchedule{
				minute:   cronUnit{stepped, []uint8{10}},
				hour:     cronUnit{},
				monthDay: cronUnit{},
				month:    cronUnit{},
				weekDay:  cronUnit{},
			},
			false,
		},

		{"too many tokens", args{"* * * * * *"}, nil, true},
		{"trailing comma", args{"1, * * * *"}, nil, true},
		{"open ranged", args{"1- * * * *"}, nil, true},
		{"excessive range", args{"5-10-15 * * * *"}, nil, true},
		{"non-standard step", args{"1/10 * * * *"}, nil, true},
		{"too many steps", args{"*/10/15 * * * *"}, nil, true},
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

func Test_cronSchedule_nextTime(t *testing.T) {
	type fields struct {
		minute   cronUnit
		hour     cronUnit
		monthDay cronUnit
		month    cronUnit
		weekDay  cronUnit
	}

	type args struct {
		from time.Time
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   time.Time
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := cronSchedule{
				minute:   tt.fields.minute,
				hour:     tt.fields.hour,
				monthDay: tt.fields.monthDay,
				month:    tt.fields.month,
				weekDay:  tt.fields.weekDay,
			}
			if got := cs.nextTime(tt.args.from); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cronSchedule.nextTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
