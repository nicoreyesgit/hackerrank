package medium

type Point struct {
	X int
	Y int
}

var directions = []Point{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func FindPathWithRecursion(way []string, start Point, end Point) []Point {
	path := make([]Point, 0)
	seen := make([][]bool, len(way))

	for i := range seen {
		seen[i] = make([]bool, len(way[i]))
	}
	i, p := findPathRecursion(way, start, end, path, seen)
	if i {
		return p
	}

	return path
}

func findPathRecursion(way []string, curr Point, end Point, path []Point, seen [][]bool) (bool, []Point) {
	if curr.X < 0 || curr.X >= len(way[0]) ||
		curr.Y < 0 || curr.Y >= len(way) {
		return false, path
	}

	if curr.X == end.X && curr.Y == end.Y {
		path = append(path, end)
		return true, path
	}

	if seen[curr.Y][curr.X] {
		return false, path
	}

	seen[curr.Y][curr.X] = true
	path = append(path, curr)

	for i := 0; i < len(directions); i++ {
		p := directions[i]
		t, fullPath := findPathRecursion(way, Point{
			X: curr.X + p.X,
			Y: curr.Y + p.Y,
		}, end, path, seen)

		if t {
			return true, fullPath
		}
	}
	path = path[:len(path)-1]

	return false, path
}
