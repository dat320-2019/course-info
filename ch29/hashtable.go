package ch29

// HashTable for key-value lookup.
type HashTable struct {
	lists   []List
	buckets int
}

// NewHashTable returns a hashtable with size buckets.
func NewHashTable(size int) *HashTable {
	return &HashTable{lists: make([]List, size), buckets: size}
}

// Insert key into hashtable.
func (h *HashTable) Insert(key int) {
	bucket := key % h.buckets
	h.lists[bucket].Insert(key)
}

// Lookup key in hashtable.
func (h *HashTable) Lookup(key int) *Node {
	bucket := key % h.buckets
	return h.lists[bucket].Lookup(key)
}
