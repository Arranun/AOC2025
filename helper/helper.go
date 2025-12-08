package helper

import "math"

type Slices interface {
	~[2]int | ~[3]int
}
type Point[T Slices] struct {
	Position T
	Symbol   rune
}

type Grid[T Slices] struct {
	Points  map[T]Point[T]
	Borders T
}

func EuclidianDistance[T Slices](point1 Point[T], point2 Point[T]) float64 {
	result := 0
	for i := 0; i < len(point1.Position); i++ {
		result += (point2.Position[i] - point1.Position[i]) * (point2.Position[i] - point1.Position[i])
	}
	return math.Sqrt(float64(result))
}

func GetGrid[T Slices](points []T, symbol rune) Grid[T] {
	var grid Grid[T]
	for i := 0; i < len(grid.Borders); i++ {
		grid.Borders[i] = 0
	}
	grid.Points = make(map[T]Point[T])
	for _, p := range points {
		grid.Points[p] = Point[T]{p, symbol}
		for i := 0; i < len(p); i++ {
			if p[i] > grid.Borders[i] {
				grid.Borders[i] = p[i]
			}
		}
	}
	return grid
}

func GetGrid2D(lines []string) Grid[[2]int] {
	var grid Grid[[2]int]
	grid.Points = make(map[[2]int]Point[[2]int])
	for i, line := range lines {
		for j, char := range line {
			if char != '.' {
				grid.Points[[2]int{j, i}] = Point[[2]int]{[2]int{j, i}, char}
			}
		}
	}
	grid.Borders[0] = len(lines[0])
	grid.Borders[1] = len(lines)
	return grid
}
