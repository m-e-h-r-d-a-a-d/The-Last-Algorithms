package sort

func QuickSort(input []int) []int {
	return qSort(input, 0, len(input)-1)
}

func qSort(input []int, lo int, hi int) []int {
	if lo >= hi {
		return input
	}

	input, p := partition(input, lo, hi)
	input = qSort(input, lo, p-1)
	input = qSort(input, p+1, hi)

	return input
}

func partition(input []int, lo int, hi int) ([]int, int) {
	p := input[hi]

	idx := lo
	for i := lo; i < hi; i++ {
		if input[i] <= p {
			input[i], input[idx] = input[idx], input[i]
			idx++
		}
	}

	input[hi], input[idx] = input[idx], input[hi]
	return input, idx
}
