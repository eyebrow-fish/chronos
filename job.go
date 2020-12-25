package chronos

import (
	"context"
	"time"
)

// Job defines a scheduled task to be executed.
//
// A Job can be as granular as one second or as long as
// you want. One constraint put in place is on the context.Context
// accessible on the Exec func. There is a deadline which ends when
// the next job is scheduled to run.
//
// By default the Job schedule is Job.Secondly. To change this
// use one of the schedule defining methods, such as Job.Daily
// or Job.Yearly.
//
// Example:
//
//  NewJob(func(ctx context.Context) error {
//    fmt.Println("Pay rent!")
//    return nil
//  }).Monthly().Run()
type Job struct {
	schedule
	Exec
	Last time.Time
}

// NewJob constructs a Job.
//
// Exec is a function which is passed in as the only parameter.
// This function is executed on the defined schedule.
//
// The "Last" field of Job is set to time.Now() by default.
//
// By default the Job schedule is Job.Secondly. To change this
// use one of the schedule defining methods, such as Job.Daily
// or Job.Yearly.
func NewJob(exec Exec) Job {
	return JobFrom(time.Now(), exec)
}

// JobFrom constructs a Job.
//
// JobFrom allows for a start time to passed in along with the
// Exec field.
//
// See NewJob for more information about constructing a Job.
func JobFrom(from time.Time, exec Exec) Job {
	return Job{Exec: exec, Last: from}
}

// Run is a method for launching the Job.
//
// Each Exec is run in a separate goroutine and executes
// once the scheduled time is reached. These times are
// rounded to their respective units (i.e. minutes are
// rounded to the nearest minute).
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

// NextRun is a method which returns the next scheduled time.Time.
//
// These times are rounded to their respective units (i.e. minutes are
// rounded to the nearest minute).
func (j Job) NextRun() time.Time {
	return j.schedule.nextDate(j.Last)
}

// Seconds is a method for creating a Job which runs every X seconds.
func (j Job) Seconds(seconds int) Job {
	j.schedule.unit = second
	j.schedule.value = seconds
	return j
}

// Minutes is a method for creating a Job which runs every X minutes.
//
// The exact time it runs at will be rounded to the nearest minute.
func (j Job) Minutes(minutes int) Job {
	j.schedule.unit = minute
	j.schedule.value = minutes
	return j
}

// Hours is a method for creating a Job which runs every X hours.
//
// The exact time it runs at will be rounded to the nearest hour.
func (j Job) Hours(hours int) Job {
	j.schedule.unit = hour
	j.schedule.value = hours
	return j
}

// Days is a method for creating a Job which runs every X days.
//
// The exact time it runs at will be rounded to the nearest day.
func (j Job) Days(days int) Job {
	j.schedule.unit = day
	j.schedule.value = days
	return j
}

// Months is a method for creating a Job which runs every X months.
//
// The exact time it runs at will be rounded to the nearest month.
func (j Job) Months(months int) Job {
	j.schedule.unit = month
	j.schedule.value = months
	return j
}

// Years is a method for creating a Job which runs every X years.
//
// The exact time it runs at will be rounded to the nearest year.
func (j Job) Years(years int) Job {
	j.schedule.unit = year
	j.schedule.value = years
	return j
}

// Secondly is a method for creating a Job which runs every second.
//
// The exact time it runs at will be rounded to the nearest second.
func (j Job) Secondly() Job {
	j.schedule.unit = second
	j.schedule.value = 1
	return j
}

// Minutely is a method for creating a Job which runs every minute.
//
// The exact time it runs at will be rounded to the nearest minute.
func (j Job) Minutely() Job {
	j.schedule.unit = minute
	j.schedule.value = 1
	return j
}

// Hourly is a method for creating a Job which runs every hour.
//
// The exact time it runs at will be rounded to the nearest hour.
func (j Job) Hourly() Job {
	j.schedule.unit = hour
	j.schedule.value = 1
	return j
}

// Daily is a method for creating a Job which runs every day.
//
// The exact time it runs at will be rounded to the nearest day.
func (j Job) Daily() Job {
	j.schedule.unit = day
	j.schedule.value = 1
	return j
}

// Monthly is a method for creating a Job which runs every month.
//
// The exact time it runs at will be rounded to the nearest month.
func (j Job) Monthly() Job {
	j.schedule.unit = month
	j.schedule.value = 1
	return j
}

// Yearly is a method for creating a Job which runs every year.
//
// The exact time it runs at will be rounded to the nearest year.
func (j Job) Yearly() Job {
	j.schedule.unit = year
	j.schedule.value = 1
	return j
}

// Exec is a func which takes a context.Context and returns an error.
//
// This Exec is what type Job uses to execute jobs.
type Exec func(context.Context) error
