package ch29

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	count := &counter{}
	count.inc()
	got := count.get()
	if got != 1 {
		t.Errorf("count.get()=%d, expected 1", got)
	}
}

func TestConcurrentCounter(t *testing.T) {
	var wg sync.WaitGroup
	count := &counter{}
	for i := 0; i < totalCount; i++ {
		wg.Add(1)
		go func() {
			count.inc()
			wg.Done()
		}()
	}
	wg.Wait()
	got := count.get()
	if got != totalCount {
		t.Errorf("count.get()=%d, expected %d", got, totalCount)
	}
}

func TestNolockCounter(t *testing.T) {
	// does not use goroutines or locks
	count := &counter{}
	tc := totalCount * numCPUs
	for j := 0; j < tc; j++ {
		count.noLockInc()
	}
	got := count.get()
	if got != tc {
		t.Errorf("count.get()=%d, expected %d", got, tc)
	}
}

func BenchmarkNoLockCounter(b *testing.B) {
	// does not use goroutines or locks
	count := &counter{}
	tc := totalCount * numCPUs
	for k := 0; k < b.N; k++ {
		for j := 0; j < tc; j++ {
			count.noLockInc()
		}
	}
}

func BenchmarkConcurrentCounter(b *testing.B) {
	count := &counter{}
	var wg sync.WaitGroup
	for k := 0; k < b.N; k++ {
		wg.Add(numCPUs)
		for i := 0; i < numCPUs; i++ {
			go func() {
				for j := 0; j < totalCount; j++ {
					count.inc()
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkConcurrentNoLockCounter(b *testing.B) {
	count := &counter{}
	var wg sync.WaitGroup
	for k := 0; k < b.N; k++ {
		wg.Add(numCPUs)
		for i := 0; i < numCPUs; i++ {
			go func() {
				for j := 0; j < totalCount; j++ {
					count.noLockInc()
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
