package scheduler

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"github.com/prometheus/common/log"
)

type mutex struct {
	*redsync.Mutex
	ctx context.Context
}

func NewMutex(ctx context.Context, conn redis.UniversalClient) *mutex {
	pool := goredis.NewPool(conn)
	rs := redsync.New(pool)

	mutexName := "lock:zset:items"

	mu := &mutex{rs.NewMutex(mutexName), ctx}
	return mu
}

func (mu *mutex) Lock() {
	if err := mu.LockContext(mu.ctx); err != nil {
		log.Error("error acquiring lock: ", err)
	}
}

func (mu *mutex) Unlock() {
	if _, err := mu.UnlockContext(mu.ctx); err != nil {
		log.Error("error releasing lock: ", err)
	}
}
