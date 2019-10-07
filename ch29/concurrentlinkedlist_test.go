package ch29

import (
	"sync"
	"testing"
)

func TestConcurrentLinkedList(t *testing.T) {
	list := &List{}
	elems := 10000

	var wg sync.WaitGroup
	wg.Add(numCPUs)
	for i := 0; i < numCPUs; i++ {
		go func(i int) {
			for k := 0; k < elems; k++ {
				e := k + elems*numCPUs
				list.Insert(e)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	wg.Add(numCPUs)
	for i := 0; i < numCPUs; i++ {
		go func(i int) {
			for k := 0; k < elems; k++ {
				e := k + elems*numCPUs
				elm := list.Lookup(e)
				if elm == nil {
					t.Errorf("Lookup(%d)=%v, expected %d", e, elm, e)
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
