package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x      int
	y      int
	symbol rune
}

type Grid struct {
	points map[[2]int]Point
}

func main() {
	args := os.Args[1:]
	file, err := os.ReadFile(args[0])
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\r\n")
	grid := getGrid(lines)
	part1 := 0
	for k, _ := range grid.points {
		if grid.checkNeighbors(k, '@') < 4 {
			part1++
		}
	}
	fmt.Println(part1)
	part2 := 0
	oldAmount := 0
	for oldAmount != len(grid.points) {
		oldAmount = len(grid.points)
		for k, _ := range grid.points {
			if grid.checkNeighbors(k, '@') < 4 {
				delete(grid.points, k)
				part2++
			}
		}
	}
	fmt.Println(part2)
}

func getGrid(lines []string) Grid {
	var grid Grid
	grid.points = make(map[[2]int]Point)
	for i, line := range lines {
		for j, char := range line {
			if char != '.' {
				grid.points[[2]int{j, i}] = Point{j, i, char}
			}
		}
	}
	return grid
}

func (g *Grid) checkNeighbors(pos [2]int, hit rune) int {
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}}
	count := 0
	for _, dir := range directions {
		newPos := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
		if g.points[newPos].symbol == hit {
			count++
		}
	}
	return count
}
