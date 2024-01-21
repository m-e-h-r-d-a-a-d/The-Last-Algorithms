package graph

import "math/bits"

const (
	MaxInt int = (1<<bits.UintSize)/2 - 1
)

func hasUnvisited(seen []bool, distances []int) bool {
	for i := 0; i < len(seen); i++ {
		if seen[i] && distances[i] < MaxInt {
			return true
		}
	}

	return false
}

func getLowestUnvisited(seen []bool, distances []int) (int, bool) {
	idx := -1
	min := MaxInt
	isUnvisited := false
	for i := 0; i < len(seen); i++ {
		if !seen[i] && distances[i] < MaxInt && distances[i] < min {
			isUnvisited = true
			min = distances[i]
			idx = i
		}
	}

	return idx, isUnvisited
}

func DijkstraList(weights [][]int, source int, target int) []int {
	var seen []bool
	var distances []int
	var prev []int

	for i := 0; i < len(weights); i++ {
		seen = append(seen, false)
		prev = append(prev, -1)
		distances = append(distances, MaxInt)
	}

	distances[source] = 0
	for {
		idx, ok := getLowestUnvisited(seen, distances)
		if !ok {
			break
		}

		seen[idx] = true

		for i := 0; i < len(weights[idx]); i++ {
			if weights[idx][i] < MaxInt && weights[idx][i]+distances[idx] < distances[i] {
				distances[i] = weights[idx][i] + distances[idx]
				prev[i] = idx
			}
		}

	}

	if distances[target] >= MaxInt {
		return []int{}
	}

	var out []int
	c := prev[target]
	out = append(out, target)
	for c != source {
		out = append(out, c)
		c = prev[c]
	}
	out = append(out, source)

	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}

	return out
}
