package search

func binarySearch(haystack []int, needle int) bool {
	low := 0
	high := len(haystack) - 1

	for {
		mid := int((low + high) / 2)
		v := haystack[mid]
		if v == needle {
			return true
		} else if v > needle {
			high = mid
		} else {
			low = mid + 1
		}

		if low >= high {
			break
		}
	}

	return false
}
