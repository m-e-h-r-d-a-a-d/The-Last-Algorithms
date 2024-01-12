package recursive

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPreOrderDfsTraverse(t *testing.T) {
	head := &binaryNode[int]{value: 7}
	head.left = &binaryNode[int]{value: 23}
	head.right = &binaryNode[int]{value: 8}
	head.left.left = &binaryNode[int]{value: 5}
	head.left.right = &binaryNode[int]{value: 4}
	head.right.left = &binaryNode[int]{value: 21}
	head.right.right = &binaryNode[int]{value: 15}

	preOrderExpect := []int{7, 23, 5, 4, 8, 21, 15}
	preOrderOutput := PreOrderDfsTraverse(head)
	require.Equal(t, preOrderExpect, preOrderOutput)
}

func TestInOrderDfsTraverse(t *testing.T) {
	head := &binaryNode[int]{value: 7}
	head.left = &binaryNode[int]{value: 23}
	head.right = &binaryNode[int]{value: 8}
	head.left.left = &binaryNode[int]{value: 5}
	head.left.right = &binaryNode[int]{value: 4}
	head.right.left = &binaryNode[int]{value: 21}
	head.right.right = &binaryNode[int]{value: 15}

	inOrderExpect := []int{5, 23, 4, 7, 21, 8, 15}
	inOrderOutput := InOrderDfsTraverse(head)
	require.Equal(t, inOrderExpect, inOrderOutput)
}

func TestPostOrderDfsTraverse(t *testing.T) {
	head := &binaryNode[int]{value: 7}
	head.left = &binaryNode[int]{value: 23}
	head.right = &binaryNode[int]{value: 8}
	head.left.left = &binaryNode[int]{value: 5}
	head.left.right = &binaryNode[int]{value: 4}
	head.right.left = &binaryNode[int]{value: 21}
	head.right.right = &binaryNode[int]{value: 15}

	postOrderExpect := []int{5, 4, 23, 21, 15, 8, 7}
	postOrderOutput := PostOrderDfsTraverse(head)
	require.Equal(t, postOrderExpect, postOrderOutput)
}
