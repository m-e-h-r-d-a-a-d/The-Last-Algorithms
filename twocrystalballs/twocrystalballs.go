package twocrystalballs

import (
	"math"
)

func twocrystalballs(breaks []bool) int {
	l := int(math.Sqrt(float64(len(breaks))))
	if l < 1 {
		return -1
	}

	i := l
	for ; i < len(breaks); i += l {
		if breaks[i] {
			break
		}
	}

	j := i - l
	for ; j < len(breaks); j++ {
		if breaks[j] {
			return j
		}
	}

	return -1
}
