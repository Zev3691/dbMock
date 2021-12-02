package util

import (
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	count uint64
}

func (c *Counter) Incr(val uint64) {
	c.mu.Lock()
	c.count += val
	c.mu.Unlock()
}

func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
