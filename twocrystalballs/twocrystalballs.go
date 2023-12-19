package twocrystalballs

import (
	"math"
)

func twocrystalballs(breaks []bool) int {
	l := int(math.Sqrt(float64(len(breaks))))
	if l < 1 {
		return 0
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
			break
		}
	}

	return j
}
