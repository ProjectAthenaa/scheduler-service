package scheduler

import "github.com/go-redis/redis/v8"

type TaskScheduler struct {
	conn redis.UniversalClient

}

func (t TaskScheduler) Jobs() []*Job {
	panic("implement me")
}

func (t TaskScheduler) Schedule(job *Job) error {
	panic("implement me")
}

func (t TaskScheduler) Stop() error {
	panic("implement me")
}