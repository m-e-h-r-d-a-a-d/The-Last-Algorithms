package queue

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnqueue(t *testing.T) {
	var q Queue[int]
	q.enqueue(5)
	require.Equal(t, q.head.value, 5)
	require.Equal(t, q.tail.value, 5)
}

func TestDequeue(t *testing.T) {
	var q Queue[int]
	q.enqueue(5)
	q.enqueue(10)
	q.enqueue(15)
	r1, err := q.dequeue()
	require.NoError(t, err)
	r2, err := q.dequeue()
	require.NoError(t, err)
	r3, err := q.dequeue()
	require.NoError(t, err)
	_, err = q.dequeue()
	require.Equal(t, err, errors.New("queue is empty"))
	require.Equal(t, r1, 5)
	require.Equal(t, r2, 10)
	require.Equal(t, r3, 15)
}

func TestPeek(t *testing.T) {
	var q Queue[int]
	q.enqueue(-10)
	p := q.peek()
	require.Equal(t, q.head.value, -10)
	require.Equal(t, p, -10)
}
