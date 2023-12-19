package bubblesort

func bubblesort(input []int) []int {
	l := len(input)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if input[i] > input[j] {
				input[i], input[j] = input[j], input[i]
			}
		}
	}

	return input
}
