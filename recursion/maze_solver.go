package recursion

type Point struct {
	x int
	y int
}

// [#,#,#,#,#,E,
// #,#
// #,#,S,#,#,#,]
//

func Walk(maze [][]string, wall string, curr Point, end Point, seen [][]bool) bool {
	// off the map
	if curr.x < 0 || curr.x >= len(maze[0]) ||
		curr.y < 0 || curr.y >= len(maze) {
		return false
	}

	// on a wall
	if maze[curr.x][curr.y] == wall {
		return false
	}

	if seen[curr.x][curr.y] {
		return false
	}

	if curr.x == end.x && curr.y == end.y {
		return true
	}

	return true
}

func MazeSolver(maze []string, wall string, start Point, end Point) []Point {
	var points = []Point{}

	return points

}
