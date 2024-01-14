package recursive

func dfsSearch(head *binaryNode[int], needle int) bool {
	if head == nil {
		return false
	}

	if head.value == needle {
		return true
	}

	if head.value > needle {
		return dfsSearch(head.left, needle)
	}

	return dfsSearch(head.right, needle)
}
