package main

import (
	"AOC2025/helper"
	"fmt"
	"os"
	"strings"
)

type tachyonGrid struct {
	currentPos map[[2]int]bool
	visited    map[[2]int]bool
	grid       helper.Grid
	splits     int
}

func main() {
	//args := os.Args[1:]
	file, err := os.ReadFile("part1")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\r\n")
	tachyon := tachyonGrid{}
	tachyon.splits = 0
	tachyon.grid = helper.GetGrid(lines)
	tachyon.visited = map[[2]int]bool{}
	tachyon.currentPos = map[[2]int]bool{}
	for _, point := range tachyon.grid.Points {
		if point.Symbol == 'S' {
			tachyon.currentPos[[2]int{point.X, point.Y}] = true
		}
	}
	part1 := 0
	for len(tachyon.currentPos) > 0 {
		if tachyon.step() {
			part1++
		}
	}
	fmt.Println(part1 - 1)
}

func (t *tachyonGrid) step() bool {
	curr := [2]int{0, 0}
	for p, _ := range t.currentPos {
		delete(t.currentPos, p)
		if !t.visited[p] {
			t.visited[p] = true
			curr = p
			break
		}
	}
	for i := curr[1]; i < t.grid.Borders[1]; i++ {
		if t.grid.Points[[2]int{curr[0], i}].Symbol == '^' {
			t.currentPos[[2]int{curr[0] + 1, i}] = true
			t.currentPos[[2]int{curr[0] + -1, i}] = true
			return true
		}
	}
	return false
}
