package ch29

import "sync"

// ACounter represents an approximate counter.
type ACounter struct {
	global    int
	glock     sync.Mutex
	local     []int
	llock     []sync.Mutex
	threshold int
}

// New creates a new approximate counter object.
func New(threshold, numCPU int) *ACounter {
	return &ACounter{
		threshold: threshold,
		local:     make([]int, numCPU),
		llock:     make([]sync.Mutex, numCPU),
	}
}

// Get returns an approximate value of the counter.
func (c *ACounter) Get() int {
	c.glock.Lock()
	defer c.glock.Unlock()
	return c.global
}

// Update updates the approximate counter with the given amount
// on the local counter, identified by threadID.
func (c *ACounter) Update(amount, threadID int) {
	c.llock[threadID].Lock()
	defer c.llock[threadID].Unlock()
	c.local[threadID] += amount
	if c.local[threadID] >= c.threshold {
		// transfer local count value to global counter
		c.glock.Lock()
		c.global += c.local[threadID]
		c.glock.Unlock()
		c.local[threadID] = 0
	}
}

// UpdateNoLocalLock updates the approximate counter with the given amount
// on the local counter, identified by threadID. This version does not use
// a lock on the local counter, only on the global counter. This is safe
// as long as only one goroutine updates the local counter.
func (c *ACounter) UpdateNoLocalLock(amount, threadID int) {
	c.local[threadID] += amount
	if c.local[threadID] >= c.threshold {
		// transfer local count value to global counter
		c.glock.Lock()
		c.global += c.local[threadID]
		c.glock.Unlock()
		c.local[threadID] = 0
	}
}
