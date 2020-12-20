package chronos

type Job struct {
	Schedule Schedule
	Exec     func()
}

func NewJob(schedule Schedule, exec func()) Job {
	return Job{schedule, exec}
}

func (j Job) Start() {
	// do the good stuff
}
