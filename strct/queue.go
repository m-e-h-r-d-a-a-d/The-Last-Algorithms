package strct

import "errors"

type queueNode[T any] struct {
	value T
	next  *node[T]
}

type Queue[T any] struct {
	length int
	head   *node[T]
	tail   *node[T]
}

func (q *Queue[T]) Enqueue(v T) {
	n := node[T]{
		value: v,
	}

	q.length++
	if q.tail == nil || q.length == 1 {
		q.tail = &n
		q.head = &n
		return
	}

	q.tail.next = &n
	q.tail = &n
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.length == 0 || q.head == nil {
		var v T
		return v, errors.New("queue is empty")
	}

	q.length--
	n := q.head
	q.head = q.head.next
	n.next = nil
	return n.value, nil
}

func (q Queue[T]) Peek() T {
	return q.head.value
}
