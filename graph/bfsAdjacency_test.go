package graph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBfsSearch(t *testing.T) {
	weights := [][]int{{0, 1, 2, 5, 0}, {1, 0, 0, 0, 0}, {0, 0, 0, 2, 0}, {0, 0, 0, 0, 5}, {0, 0, 0, 0, 0}}

	output := bfsSearch(weights, 0, 1)
	require.Equal(t, output, []int{0, 1})
	output = bfsSearch(weights, 0, 4)
	require.Equal(t, output, []int{0, 3, 4})
	output = bfsSearch(weights, 3, 4)
	require.Equal(t, output, []int{3, 4})
	output = bfsSearch(weights, 3, 1)
	require.Equal(t, output, []int{})
	output = bfsSearch(weights, 1, 2)
	require.Equal(t, output, []int{1, 0, 2})
}
