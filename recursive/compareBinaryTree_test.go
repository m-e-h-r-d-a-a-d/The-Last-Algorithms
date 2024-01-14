package recursive

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompare(t *testing.T) {
	a := &binaryNode[int]{value: 7}
	a.left = &binaryNode[int]{value: 23}
	a.right = &binaryNode[int]{value: 8}
	a.left.left = &binaryNode[int]{value: 5}
	a.left.right = &binaryNode[int]{value: 4}
	a.right.left = &binaryNode[int]{value: 21}
	a.right.right = &binaryNode[int]{value: 15}

	b := &binaryNode[int]{value: 7}
	b.left = &binaryNode[int]{value: 23}
	b.right = &binaryNode[int]{value: 8}
	b.left.left = &binaryNode[int]{value: 5}
	b.left.right = &binaryNode[int]{value: 4}
	b.right.left = &binaryNode[int]{value: 21}
	b.right.right = &binaryNode[int]{value: 15}

	result := compare(a, b)
	require.Equal(t, result, true)
	result = compare(a, a)
	require.Equal(t, result, true)
	result = compare(b, b)
	require.Equal(t, result, true)

	b.right.right = &binaryNode[int]{value: 17}
	result = compare(a, b)
	require.Equal(t, result, false)
	b = nil
	result = compare(a, b)
	require.Equal(t, result, false)
	result = compare(b, b)
	require.Equal(t, result, true)
}
