package recursive

func PreOrderSearch(head *binaryNode[int]) []int {
	return PostOrderWalk(head, []int{})
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
