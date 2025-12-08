package helper

type Point struct {
	X      int
	Y      int
	Symbol rune
}

type Grid struct {
	Points  map[[2]int]Point
	Borders [2]int
}

func GetGrid(lines []string) Grid {
	var grid Grid
	grid.Points = make(map[[2]int]Point)
	for i, line := range lines {
		for j, char := range line {
			if char != '.' {
				grid.Points[[2]int{j, i}] = Point{j, i, char}
			}
		}
	}
	grid.Borders[0] = len(lines[0])
	grid.Borders[1] = len(lines)
	return grid
}
