package helper

import (
	"fmt"
	"math"
)

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

type Distance[T Slices] struct {
	Start  Point[T]
	End    Point[T]
	Length float64
}

func EuclidianDistance[T Slices](point1 Point[T], point2 Point[T]) float64 {
	result := 0
	for i := 0; i < len(point1.Position); i++ {
		result += (point2.Position[i] - point1.Position[i]) * (point2.Position[i] - point1.Position[i])
	}
	return math.Sqrt(float64(result))
}

func ManHattanDistance[T Slices](point1 Point[T], point2 Point[T]) float64 {
	result := 0
	for i := 0; i < len(point1.Position); i++ {
		val := point2.Position[i] - point1.Position[i]
		if val < 0 {
			val *= -1
		}
		result += val
	}
	return float64(result)
}

func (g *Grid[T]) GetDistance(method func(Point[T], Point[T]) float64) []Distance[T] {
	distances := []Distance[T]{}
	for _, p1 := range g.Points {
		for _, p2 := range g.Points {
			if p1.Position != p2.Position {
				dis := method(p1, p2)
				distance := Distance[T]{p1, p2, dis}
				distances = append(distances, distance)
			}
		}
	}
	return distances
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

func Print2DGrid(g Grid[[2]int]) {
	borders := g.Borders
	for i := 0; i <= borders[1]+1; i++ {
		for j := 0; j <= borders[0]+1; j++ {
			if _, ok := g.Points[[2]int{j, i}]; ok {
				fmt.Print(string(g.Points[[2]int{j, i}].Symbol))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

}
