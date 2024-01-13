package recursive

func bfsSearch(head *binaryNode[int], needle int) bool {
	var nodes []*binaryNode[int]
	nodes = append(nodes, head)
	var x *binaryNode[int]
	for len(nodes) > 0 {
		x, nodes = nodes[0], nodes[1:]
		if x.value == needle {
			return true
		}

		if x.left != nil {
			nodes = append(nodes, x.left)
		}

		if x.right != nil {
			nodes = append(nodes, x.right)
		}
	}

	return false
}
