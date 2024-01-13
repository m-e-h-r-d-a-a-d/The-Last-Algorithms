package recursive

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBfsSearch(t *testing.T) {
	head := &binaryNode[int]{value: 7}
	head.left = &binaryNode[int]{value: 23}
	head.right = &binaryNode[int]{value: 8}
	head.left.left = &binaryNode[int]{value: 5}
	head.left.right = &binaryNode[int]{value: 4}
	head.right.left = &binaryNode[int]{value: 21}
	head.right.right = &binaryNode[int]{value: 15}

	output := bfsSearch(head, 8)
	require.Equal(t, output, true)
	output = bfsSearch(head, 15)
	require.Equal(t, output, true)
	output = bfsSearch(head, 21)
	require.Equal(t, output, true)
	output = bfsSearch(head, 33)
	require.Equal(t, output, false)
	output = bfsSearch(head, -2)
	require.Equal(t, output, false)
}
