package ch29

import "sync"

// Node for linked list.
type Node struct {
	key  int
	next *Node
}

// List is a linked list.
type List struct {
	head  *Node
	mutex sync.Mutex
}

// Insert inserts key into linked list.
func (l *List) Insert(key int) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	new := &Node{key: key, next: l.head}
	l.head = new
}

// Lookup returns the node whose key matches the given key.
func (l *List) Lookup(key int) *Node {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	curr := l.head
	for curr != nil {
		if curr.key == key {
			return curr
		}
		curr = curr.next
	}
	return nil
}
