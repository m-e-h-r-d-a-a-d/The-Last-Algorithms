package graph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDijkstraList(t *testing.T) {
	weights := [][]int{{0, 1, 5, MaxInt, MaxInt}, {MaxInt, 0, 7, 3, MaxInt}, {MaxInt, MaxInt, 0, MaxInt, 1}, {MaxInt, 1, 2, 0, MaxInt}, {MaxInt, MaxInt, MaxInt, MaxInt, 0}}

	output := DijkstraList(weights, 0, 1)
	require.Equal(t, output, []int{0, 1})
	output = DijkstraList(weights, 0, 4)
	require.Equal(t, output, []int{0, 2, 4})
	output = DijkstraList(weights, 3, 4)
	require.Equal(t, output, []int{3, 2, 4})
	output = DijkstraList(weights, 3, 1)
	require.Equal(t, output, []int{3, 1})
	output = DijkstraList(weights, 1, 2)
	require.Equal(t, output, []int{1, 3, 2})
	output = DijkstraList(weights, 4, 3)
	require.Equal(t, output, []int{})
}
