package strct

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnqueue(t *testing.T) {
	var q Queue[int]
	q.Enqueue(5)
	require.Equal(t, q.head.value, 5)
	require.Equal(t, q.tail.value, 5)
}

func TestDequeue(t *testing.T) {
	var q Queue[int]
	q.Enqueue(5)
	q.Enqueue(10)
	q.Enqueue(15)
	r1, err := q.Dequeue()
	require.NoError(t, err)
	r2, err := q.Dequeue()
	require.NoError(t, err)
	r3, err := q.Dequeue()
	require.NoError(t, err)
	_, err = q.Dequeue()
	require.Equal(t, err, errors.New("queue is empty"))
	require.Equal(t, r1, 5)
	require.Equal(t, r2, 10)
	require.Equal(t, r3, 15)
}

func TestQueuePeek(t *testing.T) {
	var q Queue[int]
	q.Enqueue(-10)
	p := q.Peek()
	require.Equal(t, q.head.value, -10)
	require.Equal(t, p, -10)
}
