package scheduler

type Scheduler interface {
	Stop()
	ProcessEvents()
}

const (
	itemsKey = "zset:scheduled:items"
	queueKey = "queue:%s"
)
