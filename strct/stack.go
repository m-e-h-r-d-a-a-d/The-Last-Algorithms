package strct

import "errors"

type stackNode[T any] struct {
	value T
	next  *stackNode[T]
}

type Stack[T any] struct {
	length int
	head   *stackNode[T]
}

func (s *Stack[T]) Push(v T) {
	n := stackNode[T]{
		value: v,
	}

	s.length++
	if s.head == nil || s.length == 1 {
		s.head = &n
		return
	}

	n.next = s.head
	s.head = &n
}

func (s *Stack[T]) Pop() (T, error) {
	if s.length == 0 || s.head == nil {
		var v T
		return v, errors.New("queue is empty")
	}

	s.length--
	n := s.head
	s.head = s.head.next
	n.next = nil
	return n.value, nil
}

func (s Stack[T]) Peek() T {
	return s.head.value
}
