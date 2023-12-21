package strct

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPush(t *testing.T) {
	var s Stack[int]
	s.Push(5)
	require.Equal(t, s.head.value, 5)
	s.Push(-2)
	require.Equal(t, s.head.value, -2)
	s.Push(555)
	require.Equal(t, s.head.value, 555)
}

func TestPop(t *testing.T) {
	var s Stack[int]
	s.Push(5)
	s.Push(-10)
	s.Push(15)
	r1, err := s.Pop()
	require.NoError(t, err)
	r2, err := s.Pop()
	require.NoError(t, err)
	r3, err := s.Pop()
	require.NoError(t, err)
	_, err = s.Pop()
	require.Equal(t, err, errors.New("queue is empty"))
	require.Equal(t, r1, 15)
	require.Equal(t, r2, -10)
	require.Equal(t, r3, 5)
}

func TestStackPeek(t *testing.T) {
	var s Stack[int]
	s.Push(-10)
	p := s.Peek()
	require.Equal(t, s.head.value, -10)
	require.Equal(t, p, -10)
}
