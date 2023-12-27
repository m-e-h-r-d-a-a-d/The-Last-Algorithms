package strct

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppend(t *testing.T) {
	var d DoublyLinkedList[int]
	d.Append(5)
	require.Equal(t, d.tail.value, 5)
	require.Equal(t, d.head.value, 5)

	d.Append(12)
	require.Equal(t, d.tail.value, 12)
	require.Equal(t, d.head.value, 5)

	d.Append(-34)
	require.Equal(t, d.tail.value, -34)
	require.Equal(t, d.head.value, 5)

	d.Append(0)
	require.Equal(t, d.tail.value, 0)
	require.Equal(t, d.head.value, 5)
}

func TestPrepend(t *testing.T) {
	var d DoublyLinkedList[int]
	d.Prepend(5)
	require.Equal(t, d.head.value, 5)
	require.Equal(t, d.tail.value, 5)

	d.Prepend(12)
	require.Equal(t, d.head.value, 12)
	require.Equal(t, d.tail.value, 5)

	d.Prepend(-34)
	require.Equal(t, d.head.value, -34)
	require.Equal(t, d.tail.value, 5)

	d.Prepend(0)
	require.Equal(t, d.head.value, 0)
	require.Equal(t, d.tail.value, 5)
}

func TestAll(t *testing.T) {
	var d DoublyLinkedList[int]
	d.InsertAt(43, 0)
	require.Equal(t, d.head.value, 43)
	require.Equal(t, d.tail.value, 43)

	d.InsertAt(12, 1)
	require.Equal(t, d.head.value, 43)
	require.Equal(t, d.tail.value, 12)

	d.InsertAt(-23, 1)
	require.Equal(t, d.head.value, 43)
	require.Equal(t, d.tail.value, 12)

	d.InsertAt(222, 2)
	require.Equal(t, d.head.value, 43)
	require.Equal(t, d.tail.value, 12)

	// 43,-23,222,12
	v, err := d.Get(0)
	require.NoError(t, err)
	require.Equal(t, v, 43)

	v, err = d.Get(1)
	require.NoError(t, err)
	require.Equal(t, v, -23)

	v, err = d.Get(2)
	require.NoError(t, err)
	require.Equal(t, v, 222)

	v, err = d.Get(3)
	require.NoError(t, err)
	require.Equal(t, v, 12)

	_, err = d.Get(4)
	require.Equal(t, err, errors.New("index out of range"))

	v, err = d.Remove(43)
	require.NoError(t, err)
	require.Equal(t, v, 43)

	v, err = d.Remove(222)
	require.NoError(t, err)
	require.Equal(t, v, 222)

	_, err = d.Remove(50)
	require.Equal(t, err, errors.New("item not exist"))
}
