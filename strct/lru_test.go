package strct

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLRU(t *testing.T) {
	lru := NewLRU[string, int](2)
	lru.Update("a", 234)
	require.Equal(t, lru.head.value, 234)
	require.Equal(t, lru.tail.value, 234)
	v, _ := lru.Get("a")
	require.Equal(t, v, 234)

	lru.Update("b", 453)
	require.Equal(t, lru.head.value, 453)
	require.Equal(t, lru.tail.value, 234)
	lru.Update("a", -23)
	v, _ = lru.Get("a")
	require.Equal(t, v, -23)
	v, _ = lru.Get("b")
	require.Equal(t, v, 453)
	_, err := lru.Get("c")
	require.Equal(t, err, errors.New("With this key value not found"))

	lru.Update("c", 4)
	require.Equal(t, lru.head.value, 4)
	require.Equal(t, lru.tail.value, 453)
}
