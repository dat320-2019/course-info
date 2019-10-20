package main

import (
	"fmt"
	"sync"
	"time"
)

func assert(predicate bool, msg string) {
	if !predicate {
		panic(fmt.Sprintf("predicate violated: %s", msg))
	}
}

var (
	count  = 0
	buffer int
	m      sync.Mutex
	filled = sync.NewCond(&m)
	empty  = sync.NewCond(&m)
)

func put(value int) {
	// can only put on empty buffer
	assert(count == 0, "count not zero")
	count = 1
	buffer = value
}

func get() int {
	// can only get on non-empty buffer
	assert(count == 1, "count is zero")
	count = 0
	return buffer
}

func producer(elementsToProduce int) {
	for i := 0; i < elementsToProduce; i++ {
		m.Lock()
		for count == 1 {
			empty.Wait()
		}
		put(i)
		filled.Signal()
		m.Unlock()
	}
}

func consumer() {
	for {
		m.Lock()
		for count == 0 {
			filled.Wait()
		}
		tmp := get()
		fmt.Printf("consumed %d\n", tmp)
		empty.Signal()
		m.Unlock()
	}
}

func main() {
	go producer(1000)
	go consumer()
	go consumer()
	// quick hack to wait for a second
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Waiting for one second")
	time.Sleep(1 * time.Second)
}
