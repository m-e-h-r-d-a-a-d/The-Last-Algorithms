package recursive

func compare(a *binaryNode[int], b *binaryNode[int]) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.value != b.value {
		return false
	}

	return compare(a.left, b.left) && compare(a.right, b.right)
}
