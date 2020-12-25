package chronos

import (
	"context"
	"time"
)

type Job struct {
	schedule
	Exec
	Last time.Time
}

func NewJob(exec Exec) Job {
	return JobFrom(time.Now(), exec)
}

func JobFrom(from time.Time, exec Exec) Job {
	return Job{Exec: exec, Last: from}
}

func (j Job) Run() {
	rootCtx := context.Background()
	for i := 0; ; i++ {
		ticker := time.NewTicker(time.Until(j.NextRun()))
		select {
		case <-ticker.C:
			j.Last = time.Now()
			ctx, cancel := context.WithDeadline(rootCtx, j.NextRun())
			go func() {
				if err := j.Exec(ctx); err != nil {
					cancel()
				}
				ticker.Stop()
			}()
		}
	}
}

func (j Job) NextRun() time.Time {
	return j.schedule.nextDate(j.Last)
}

func (j Job) Seconds(seconds int) Job {
	j.schedule.unit = second
	j.schedule.value = seconds
	return j
}

func (j Job) Minutes(minutes int) Job {
	j.schedule.unit = minute
	j.schedule.value = minutes
	return j
}

func (j Job) Hours(hours int) Job {
	j.schedule.unit = hour
	j.schedule.value = hours
	return j
}

func (j Job) Days(days int) Job {
	j.schedule.unit = day
	j.schedule.value = days
	return j
}

func (j Job) Months(months int) Job {
	j.schedule.unit = month
	j.schedule.value = months
	return j
}

func (j Job) Years(years int) Job {
	j.schedule.unit = year
	j.schedule.value = years
	return j
}

func (j Job) Secondly() Job {
	j.schedule.unit = second
	j.schedule.value = 1
	return j
}

func (j Job) Minutely() Job {
	j.schedule.unit = minute
	j.schedule.value = 1
	return j
}

func (j Job) Hourly() Job {
	j.schedule.unit = hour
	j.schedule.value = 1
	return j
}

func (j Job) Daily() Job {
	j.schedule.unit = day
	j.schedule.value = 1
	return j
}

func (j Job) Monthly() Job {
	j.schedule.unit = month
	j.schedule.value = 1
	return j
}

func (j Job) Yearly() Job {
	j.schedule.unit = year
	j.schedule.value = 1
	return j
}

type Exec func(context.Context) error
