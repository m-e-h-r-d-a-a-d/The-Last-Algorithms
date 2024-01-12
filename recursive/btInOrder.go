package recursive

type binaryNode[T any] struct {
	value T
	left  *binaryNode[T]
	right *binaryNode[T]
}

func InOrderSearch(head *binaryNode[int]) []int {
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
