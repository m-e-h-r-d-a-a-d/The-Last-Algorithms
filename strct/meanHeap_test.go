package strct

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsert(t *testing.T) {
	var m MeanHeap

	m.Insert(5)
	m.Insert(33)
	m.Insert(6)
	m.Insert(-12)
	m.Insert(91)
	m.Insert(-1)
	m.Insert(50)
	v, _ := m.Delete()
	require.Equal(t, v, -12)
	v, _ = m.Delete()
	require.Equal(t, v, -1)
	v, _ = m.Delete()
	require.Equal(t, v, 5)
	v, _ = m.Delete()
	require.Equal(t, v, 6)
	v, _ = m.Delete()
	require.Equal(t, v, 33)
	v, _ = m.Delete()
	require.Equal(t, v, 50)
	v, _ = m.Delete()
	require.Equal(t, v, 91)
	_, err := m.Delete()
	require.Equal(t, err, errors.New("there is now data in heap"))
}
