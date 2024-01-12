package recursive

type binaryNode[T any] struct {
	value T
	left  *binaryNode[T]
	right *binaryNode[T]
}

func InOrderDfsTraverse(head *binaryNode[int]) []int {
	return InOrderWalk(head, []int{})
}

func InOrderWalk(head *binaryNode[int], list []int) []int {
	if head == nil {
		return list
	}

	list = InOrderWalk(head.left, list)
	list = append(list, head.value)
	list = InOrderWalk(head.right, list)
	return list
}

func PreOrderDfsTraverse(head *binaryNode[int]) []int {
	return PreOrderWalk(head, []int{})
}

func PreOrderWalk(head *binaryNode[int], list []int) []int {
	if head == nil {
		return list
	}

	list = append(list, head.value)
	list = PreOrderWalk(head.left, list)
	list = PreOrderWalk(head.right, list)
	return list
}

func PostOrderDfsTraverse(head *binaryNode[int]) []int {
	return PostOrderWalk(head, []int{})
}

func PostOrderWalk(head *binaryNode[int], list []int) []int {
	if head == nil {
		return list
	}

	list = PostOrderWalk(head.left, list)
	list = PostOrderWalk(head.right, list)
	list = append(list, head.value)
	return list
}
