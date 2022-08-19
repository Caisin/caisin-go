package ids

import (
	"strconv"
	"sync"
	"time"
)

const (
	workerBits  uint8 = 10
	numberBits  uint8 = 12
	workerMax   int64 = -1 ^ (-1 << workerBits)
	numberMax   int64 = -1 ^ (-1 << numberBits)
	timeShift         = workerBits + numberBits
	workerShift       = numberBits
	startTime   int64 = 1525705533000 // 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID
)

type Worker struct {
	mu        sync.Mutex
	timestamp int64
	workerId  int64
	number    int64
}

func NewWorker(workerId int64) *Worker {
	if workerId < 0 || workerId > workerMax {
		return nil
	}
	// 生成一个新节点
	return &Worker{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}
}

func (c *Worker) NextId() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now().UnixNano() / 1e6
	if c.timestamp == now {
		c.number++
		if c.number > numberMax {
			for now <= c.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		c.number = 0
		c.timestamp = now
	}
	ID := (now-startTime)<<timeShift | (c.workerId << workerShift) | (c.number)
	return ID
}

// NextStr /* 生成唯一编号 */
func (c *Worker) NextStr(pre string) string {
	generate := c.NextId()
	return pre + strconv.FormatInt(generate, 10)
}
