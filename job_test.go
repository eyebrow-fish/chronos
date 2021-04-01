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
