package recursive

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	m := []string{
		"xxxxxxxxxx x",
		"x        x  ",
		"x      xxxx ",
		"x xxxxx   x ",
		"x       x   ",
		"x xxxxxxxxxx",
	}

	p := []point{
		{x: 0, y: 10},
		{x: 1, y: 10},
		{x: 1, y: 11},
		{x: 2, y: 11},
		{x: 3, y: 11},
		{x: 4, y: 11},
		{x: 4, y: 10},
		{x: 4, y: 9},
		{x: 3, y: 9},
		{x: 3, y: 8},
		{x: 3, y: 7},
		{x: 4, y: 7},
		{x: 4, y: 6},
		{x: 4, y: 5},
		{x: 4, y: 4},
		{x: 4, y: 3},
		{x: 4, y: 2},
		{x: 4, y: 1},
		{x: 5, y: 1},
	}

	t.Run("maze", func(t *testing.T) {
		result := Solve(m, 'x', point{x: 0, y: 10}, point{x: 5, y: 1})
		require.Equal(t, result, p)
	})
}
