package scheduler

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/prometheus/common/log"
	"strconv"
	"strings"
	"time"
)

type TaskScheduler struct {
	ctx    context.Context
	cancel context.CancelFunc
	conn   redis.UniversalClient
	mu     *mutex
}

func (t *TaskScheduler) ProcessEvents() {
	for {
		t.mu.Lock()
		value, err := t.GetTopFromSortedSet()
		if err != nil {
			log.Error("error on event processing: ", err)
			t.mu.Unlock()
			continue
		}

		if len(value) == 0 {
			t.mu.Unlock()
			time.Sleep(time.Second)
			continue
		}

		if v := strings.Split(value, ":"); len(v) == 2 {
			timestampString := v[1]

			timestamp, err := strconv.ParseInt(timestampString, 10, 64)
			if err != nil {
				t.mu.Unlock()
				log.Error("error processing timestamp: ", err)
				continue
			}

			startTime := time.Unix(timestamp, 0)

			if time.Since(startTime) < 0 {
				if err = t.DeleteItemFromSet(value); err != nil {
					t.mu.Unlock()
					log.Error("error removing item from set: ", err)
					continue
				}

				if err = t.AddItemToQueue(strconv.Itoa(int(startTime.Unix())), v[0]); err != nil {
					t.mu.Unlock()
					log.Error("error adding item to queue: ", err)
					continue
				}
			}

		}

		t.mu.Unlock()
	}
}


func (t *TaskScheduler) Stop() {
	t.cancel()
}
