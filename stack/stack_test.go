package stack

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPush(t *testing.T) {
	var q Stack[int]
	q.push(5)
	require.Equal(t, q.head.value, 5)
	q.push(-2)
	require.Equal(t, q.head.value, -2)
	q.push(555)
	require.Equal(t, q.head.value, 555)
}

func TestPop(t *testing.T) {
	var q Stack[int]
	q.push(5)
	q.push(-10)
	q.push(15)
	r1, err := q.pop()
	require.NoError(t, err)
	r2, err := q.pop()
	require.NoError(t, err)
	r3, err := q.pop()
	require.NoError(t, err)
	_, err = q.pop()
	require.Equal(t, err, errors.New("queue is empty"))
	require.Equal(t, r1, 15)
	require.Equal(t, r2, -10)
	require.Equal(t, r3, 5)
}

func TestPeek(t *testing.T) {
	var q Stack[int]
	q.push(-10)
	p := q.peek()
	require.Equal(t, q.head.value, -10)
	require.Equal(t, p, -10)
}
