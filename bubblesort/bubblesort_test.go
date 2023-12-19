package bubblesort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLinearSearch(t *testing.T) {
	testCases := []struct {
		name     string
		unsorted []int
		sorted   []int
	}{
		{
			name:     "exist 1",
			unsorted: []int{2, 12, 15, 1, -60, 133, 12, -200},
			sorted:   []int{-200, -60, 1, 2, 12, 12, 15, 133},
		},
		{
			name:     "exist 2",
			unsorted: []int{100, 100, 100, 100, 100, 100, 100, 100},
			sorted:   []int{100, 100, 100, 100, 100, 100, 100, 100},
		},
		{
			name:     "not exist 1",
			unsorted: []int{-200, -60, 1, 2, 12, 12, 15, 133},
			sorted:   []int{-200, -60, 1, 2, 12, 12, 15, 133},
		},
		{
			name:     "not exist 2",
			unsorted: []int{2122, 500, 15, 12, -60, -133, -222, -1000},
			sorted:   []int{-1000, -222, -133, -60, 12, 15, 500, 2122},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			result := bubblesort(tc.unsorted)
			require.Equal(t, result, tc.sorted)
		})
	}
}
