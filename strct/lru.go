package strct

import (
	"errors"
)

type LRU[K comparable, V comparable] struct {
	length        int
	capacity      int
	head          *DoublyNode[V]
	tail          *DoublyNode[V]
	lookup        map[K]*DoublyNode[V]
	reverseLookup map[*DoublyNode[V]]K
}

func NewLRU[K comparable, V comparable](capacity int) *LRU[K, V] {
	return &LRU[K, V]{
		length:        0,
		capacity:      capacity,
		head:          nil,
		tail:          nil,
		lookup:        make(map[K]*DoublyNode[V]),
		reverseLookup: make(map[*DoublyNode[V]]K),
	}
}

func CreateValue[V comparable](v V) *DoublyNode[V] {
	return &DoublyNode[V]{
		value: v,
	}
}

func (l *LRU[K, V]) Update(key K, value V) {
	node, ok := l.lookup[key]
	if ok {
		node.value = value
		l.detach(node)
		l.prepend(node)
		return
	}

	node = CreateValue[V](value)
	l.prepend(node)
	l.length++
	l.lookup[key] = node
	l.reverseLookup[node] = key
	l.trimCache()
}

func (l *LRU[K, V]) Get(key K) (V, error) {
	node, ok := l.lookup[key]
	if !ok {
		var v V
		return v, errors.New("With this key value not found")
	}

	l.detach(node)
	l.prepend(node)
	return node.value, nil
}

func (l *LRU[K, V]) detach(node *DoublyNode[V]) {
	if node.previous != nil {
		node.previous.next = node.next
	}

	if node.next != nil {
		node.next.previous = node.previous
	}

	if l.head == node {
		l.head = node.next
	}

	if l.tail == node {
		l.tail = node.previous
	}

	node.next = nil
	node.previous = nil
}

func (l *LRU[K, V]) prepend(node *DoublyNode[V]) {
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	node.next = l.head
	l.head.previous = node
	l.head = node
}

func (l *LRU[K, V]) trimCache() {
	if l.length <= l.capacity {
		return
	}

	tail := l.tail
	l.detach(l.tail)
	key := l.reverseLookup[tail]
	delete(l.lookup, key)
	delete(l.reverseLookup, tail)
}
