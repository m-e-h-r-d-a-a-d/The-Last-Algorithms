package graph

type AdjacencyMatrix struct {
	weights [][]int
}

func bfsSearch(weights [][]int, source int, needle int) []int {
	var seen []bool
	var prev []int
	var q []int
	found := false

	for i := 0; i < len(weights); i++ {
		seen = append(seen, false)
		prev = append(prev, -1)
	}

	q = append(q, source)
	seen[source] = true

	for len(q) > 0 {
		c := q[0]
		q = q[1:]
		if c == needle {
			found = true
			break
		}

		for i := 0; i < len(weights[c]); i++ {
			if weights[c][i] == 0 || seen[i] {
				continue
			}
			seen[i] = true
			prev[i] = c
			q = append(q, i)
		}
	}

	if !found {
		return []int{}
	}

	var out []int
	c := prev[needle]
	out = append(out, needle)
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
