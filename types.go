package scheduler

type Scheduler interface {
	Jobs() []*Job
	Schedule(job *Job) error
	Stop() error
}

const (
	jobsKey = "scheduler:jobs"
)