package main

import (
	"fmt"
	"sync"
)

// Run with:
//   go run thread_join.go
//
// Or build it first and run the binary.

var (
	done     bool
	mutex    sync.Mutex
	waitCond *sync.Cond = sync.NewCond(&mutex)
)

func thrExit() {
	mutex.Lock()
	done = true
	waitCond.Signal()
	mutex.Unlock()
}

func thrJoin() {
	mutex.Lock()
	for !done {
		waitCond.Wait()
	}
	mutex.Unlock()
}

func child() {
	fmt.Println("child")
	thrExit()
}

func main() {
	fmt.Println("parent: begin")
	go child()
	thrJoin()
	fmt.Println("parent: end")
}
