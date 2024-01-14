package recursive

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDfsSearch(t *testing.T) {
	head := &binaryNode[int]{value: 15}
	head.left = &binaryNode[int]{value: 8}
	head.right = &binaryNode[int]{value: 20}
	head.left.left = &binaryNode[int]{value: 5}
	head.left.right = &binaryNode[int]{value: 10}
	head.right.left = &binaryNode[int]{value: 17}
	head.right.right = &binaryNode[int]{value: 25}

	result := dfsSearch(head, 20)
	require.Equal(t, result, true)
	result = dfsSearch(head, 15)
	require.Equal(t, result, true)
	result = dfsSearch(head, 10)
	require.Equal(t, result, true)
	result = dfsSearch(head, -21)
	require.Equal(t, result, false)
	result = dfsSearch(head, 9)
	require.Equal(t, result, false)
	head = nil
	result = dfsSearch(head, 25)
	require.Equal(t, result, false)
}
