package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type lotteryScheduler interface {
	Schedule()
	Add(node)
}

type lott struct {
	head *node
}

func (l lott) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "Job: %s, Tickets: %d\n", l.head.jobName, l.head.tickets)
	tmp := l.head
	for tmp.next != nil {
		fmt.Fprintf(&b, "Job: %s, Tickets: %d\n", tmp.next.jobName, tmp.next.tickets)
		tmp = tmp.next
	}
	return fmt.Sprintf("Lottery Jobs:\n%v\n", b.String())
}

func (l *lott) Add(n *node) {
	if l.head == nil {
		l.head = n
		fmt.Println("saved the head:", n)
		return
	}
	tmp := l.head
	for tmp != nil {
		if tmp.next == nil {
			break
		}
		tmp = tmp.next
	}
	tmp.next = n
}

func (l lott) Schedule() {
	rand.Seed(time.Now().UnixNano())
	rnd := rand.Int31n(399)
	fmt.Println("rnd=", rnd)
	current := l.head
	var cnt int32
	for current != nil {
		cnt += int32(current.tickets)
		if cnt > rnd {
			fmt.Println("winner is: ", current.jobName)
			break
		}
		current = current.next
	}

}

type node struct {
	jobName string
	tickets int
	next    *node
}

func main() {
	jobA := &node{jobName: "A", tickets: 100}
	jobB := &node{jobName: "B", tickets: 50}
	jobC := &node{jobName: "C", tickets: 250}

	lottery := &lott{}
	lottery.Add(jobA)
	lottery.Add(jobB)
	lottery.Add(jobC)
	lottery.Schedule()

	fmt.Println(lottery)
}
