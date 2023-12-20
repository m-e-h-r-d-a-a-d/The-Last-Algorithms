package stack

import "errors"

type node[T any] struct {
	value T
	next  *node[T]
}

type Stack[T any] struct {
	length int
	head   *node[T]
}

func (q *Stack[T]) push(v T) {
	n := node[T]{
		value: v,
	}

	q.length++
	if q.head == nil || q.length == 1 {
		q.head = &n
		return
	}

	n.next = q.head
	q.head = &n
}

func (q *Stack[T]) pop() (T, error) {
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

func (q *Stack[T]) peek() T {
	return q.head.value
}
