package recursion

type Point struct {
	y int
	x int
}

var dir []Point = []Point{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}

func Walk(maze [][]string, wall string, curr Point, end Point,
	seen [][]bool, path *[]Point) bool {
	// off the map

	if curr.x < 0 || curr.x >= len(maze[0]) ||
		curr.y < 0 || curr.y >= len(maze) {
		return false
	}

	// on a wall
	if maze[curr.y][curr.x] == wall {
		return false
	}

	if seen[curr.y][curr.x] {
		return false
	}

	if curr.x == end.x && curr.y == end.y {
		*path = append(*path, end)
		return true
	}
	seen[curr.y][curr.x] = true
	*path = append(*path, curr)
	for i := 0; i < len(dir); i++ {
		var p Point
		p.x = curr.x + dir[i].x
		p.y = curr.y + dir[i].y

		if Walk(maze, wall, p, end, seen, path) {
			return true
		}
	}
	*path = (*path)[:len(*path)-1]
	return false
}

func Solve(maze [][]string, wall string, start Point, end Point) []Point {
	var path []Point

	seen := make([][]bool, len(maze))

	for i := 0; i < len(maze); i++ {
		seen[i] = make([]bool, len(maze[i]))
		for j := 0; j < len(maze[i]); j++ {
			seen[i][j] = false
		}
	}
	Walk(maze, wall, start, end, seen, &path)
	return path
}
