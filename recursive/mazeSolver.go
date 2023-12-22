package recursive

type point struct {
	x int
	y int
}

func walk(maze []string, wall rune, curr point, end point, seen [][]bool, path []point) (bool, []point) {
	if curr.x < 0 || curr.y < 0 {
		return false, nil
	}

	if []rune(maze[curr.x])[curr.y] == wall {
		return false, nil
	}

	if seen[curr.x][curr.y] {
		return false, nil
	}

	if curr.x == end.x && curr.y == end.y {
		path = append(path, curr)
		return true, path
	}

	seen[curr.x][curr.y] = true
	path = append(path, curr)

	if ok, newPath := walk(maze, wall, point{curr.x + -1, curr.y + 0}, end, seen, path); ok {
		return true, newPath
	}

	if ok, newPath := walk(maze, wall, point{curr.x + 0, curr.y + -1}, end, seen, path); ok {
		return true, newPath
	}

	if ok, newPath := walk(maze, wall, point{curr.x + 1, curr.y + 0}, end, seen, path); ok {
		return true, newPath
	}

	if ok, newPath := walk(maze, wall, point{curr.x + 0, curr.y + 1}, end, seen, path); ok {
		return true, newPath
	}

	return false, nil
}

func Solve(maze []string, wall rune, start point, end point) []point {
	var path []point
	var seen [][]bool
	for _, v := range maze {
		var temp []bool
		for range v {
			temp = append(temp, false)
		}
		seen = append(seen, temp)
	}

	_, path = walk(maze, wall, start, end, seen, path)
	return path
}
