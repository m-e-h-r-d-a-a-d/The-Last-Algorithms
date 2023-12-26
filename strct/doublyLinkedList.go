package strct

import (
	"errors"
)

// gonull.Nullable[
type DoublyNode[T comparable] struct {
	value    T
	previous *DoublyNode[T]
	next     *DoublyNode[T]
}

type DoublyLinkedList[T comparable] struct {
	length int
	head   *DoublyNode[T]
	tail   *DoublyNode[T]
}

func (d *DoublyLinkedList[T]) Prepend(item T) {
	d.length++
	n := DoublyNode[T]{
		value: item,
	}

	if d.length == 0 || d.head == nil || d.tail == nil {
		d.head = &n
		d.tail = &n
		return
	}

	n.next = d.head
	d.head.previous = &n
	d.head = &n
}

func (d *DoublyLinkedList[T]) Append(item T) {
	d.length++
	n := DoublyNode[T]{
		value: item,
	}

	if d.length == 0 || d.head == nil || d.tail == nil {
		d.head = &n
		d.tail = &n
		return
	}

	n.previous = d.head
	d.tail.next = &n
	d.tail = &n
}

func (d *DoublyLinkedList[T]) InsertAt(item T, idx int) error {
	if idx > d.length {
		return errors.New("index out of range")
	}

	if idx == d.length {
		d.Append(item)
		return nil
	} else if idx == 0 {
		d.Prepend(item)
	}

	d.length++
	n := DoublyNode[T]{
		value: item,
	}

	c, err := d.getAt(idx)
	if err != nil {
		return err
	}

	n.next = c
	n.previous = c.previous
	c.previous = &n

	if n.previous != nil {
		n.previous.next = &n
	}

	return nil
}

func (d *DoublyLinkedList[T]) Get(idx int) (T, error) {
	n, err := d.getAt(idx)
	if err != nil {
		var v T
		return v, err
	}

	return n.value, err
}

func (d *DoublyLinkedList[T]) getAt(idx int) (*DoublyNode[T], error) {
	if idx >= d.length {
		return nil, errors.New("index out of range")
	}

	if idx == d.length {
		return d.tail, nil
	}

	c := d.head
	for i := 0; i < idx; i++ {
		c = c.next
	}

	return c, nil
}

func (d *DoublyLinkedList[T]) Remove(item T) (T, error) {
	c := d.head

	i := 0
	for c != nil && c.value != item && i < d.length {
		i++
		c = c.next
	}

	if c == nil {
		var v T
		return v, errors.New("item not exist")
	}

	return d.removeNode(c)
}

func (d *DoublyLinkedList[T]) RemoveAt(idx int) (T, error) {
	if idx > d.length {
		var v T
		return v, errors.New("index out of range")
	}

	n, err := d.getAt(idx)
	if err != nil {
		var v T
		return v, err
	}

	return d.removeNode(n)
}

func (d *DoublyLinkedList[T]) removeNode(node *DoublyNode[T]) (T, error) {
	d.length--
	if d.length == 0 {
		v := d.head.value
		d.head = nil
		d.tail = nil
		return v, nil
	}

	if node.previous != nil {
		node.previous.next = node.next
	}

	if node.next != nil {
		node.next.previous = node.previous
	}

	if node == d.head {
		d.head = node.next
	}

	if node == d.tail {
		d.tail = node.previous
	}

	node.previous = nil
	node.next = nil
	return node.value, nil
}
