package scheduler

import "fmt"

func (t *TaskScheduler) GetTopFromSortedSet() (string, error) {
	payload, err := t.conn.ZRangeWithScores(t.ctx, itemsKey, 0, 0).Result()
	if err != nil {
		return "", err
	}

	if len(payload) == 0 {
		return "", nil
	} else if payload[0].Member == nil {
		return "", nil
	} else {
		return payload[0].Member.(string), nil
	}
}

func (t *TaskScheduler) DeleteItemFromSet(item string) error {
	if err := t.conn.ZRem(t.ctx, itemsKey, item).Err(); err != nil {
		return err
	}
	return nil
}

func (t *TaskScheduler) AddItemToQueue(key, payload string) error {
	if err := t.conn.RPush(t.ctx, fmt.Sprintf(queueKey, key), payload).Err(); err != nil {
		return err
	}
	return nil
}

