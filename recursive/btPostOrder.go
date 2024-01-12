package recursive

func PostOrderSearch(head *binaryNode[int]) []int {
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
