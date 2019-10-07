package ch29

import "sync"

type counter struct {
	value int
	mutex sync.Mutex
}

func (c *counter) noLockInc() {
	c.value++
}

func (c *counter) inc() {
	c.mutex.Lock()
	c.value++
	c.mutex.Unlock()
}

func (c *counter) get() int {
	return c.value
}
