package ch29

import (
	"fmt"
	"sync"
	"testing"
)

func BenchmarkConcurrentApproximateCounterNoLocalLock(b *testing.B) {
	for threshold := 1; threshold < 8193; threshold = 2 * threshold {
		b.Run(fmt.Sprintf("%5d", threshold), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				runLoopNoLocalLock(threshold)
			}
		})
	}
}

func runLoopNoLocalLock(threshold int) {
	c := New(threshold, numCPUs)
	var wg sync.WaitGroup
	wg.Add(numCPUs)
	for i := 0; i < numCPUs; i++ {
		go func(i int) {
			for k := 0; k < totalCount; k++ {
				c.UpdateNoLocalLock(1, i)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func BenchmarkConcurrentApproximateCounter(b *testing.B) {
	for threshold := 1; threshold < 8193; threshold = 2 * threshold {
		b.Run(fmt.Sprintf("%5d", threshold), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				runLoop(threshold)
			}
		})
	}
}

func runLoop(threshold int) {
	c := New(threshold, numCPUs)
	var wg sync.WaitGroup
	wg.Add(numCPUs)
	for i := 0; i < numCPUs; i++ {
		go func(i int) {
			for k := 0; k < totalCount; k++ {
				c.Update(1, i)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
