package main

import (
	"AOC2025/helper"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type tachyonGrid struct {
	currentSplitter map[[2]int]bool
	visitedSplitter map[[2]int][][2]int
	grid            helper.Grid[[2]int]
}

func main() {
	//args := os.Args[1:]
	file, err := os.ReadFile("part1")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\r\n")
	tachyon := tachyonGrid{}
	tachyon.grid = helper.GetGrid2D(lines)
	tachyon.visitedSplitter = map[[2]int][][2]int{}
	tachyon.currentSplitter = map[[2]int]bool{}
	firstSplitter := [2]int{0, 0}
	for _, point := range tachyon.grid.Points {
		if point.Symbol == 'S' {
			for i := point.Position[1]; i < tachyon.grid.Borders[1]; i++ {
				if tachyon.grid.Points[[2]int{point.Position[0], i}].Symbol == '^' {
					tachyon.currentSplitter[[2]int{point.Position[0], i}] = true
					firstSplitter = [2]int{point.Position[0], i}
					break
				}
			}
		}
	}
	for len(tachyon.currentSplitter) > 0 {
		tachyon.step()
	}
	fmt.Println(len(tachyon.visitedSplitter))
	for s, e := range tachyon.visitedSplitter {
		newE := [][2]int{}
		for _, v := range e {
			if v != [2]int{-1, -1} {
				newE = append(newE, v)
			}
		}
		tachyon.visitedSplitter[s] = newE
	}
	amount := map[[2]int]int{}
	finished := map[[2]int]bool{}
	for s, v := range tachyon.visitedSplitter {
		amount[s] = 1
		if len(v) == 0 {
			finished[s] = true
		}
	}
	for len(finished) > 0 {
		fin := [2]int{0, 0}
		for f, _ := range finished {
			fin = f
			break
		}
		delete(finished, fin)
		for s, v := range tachyon.visitedSplitter {
			if len(tachyon.visitedSplitter[s]) == 0 {
				continue
			}
			newV := [][2]int{}
			for _, v1 := range v {
				if v1 == fin {
					amount[s] += amount[fin]

				} else {
					newV = append(newV, v1)
				}
			}
			tachyon.visitedSplitter[s] = newV
			if len(tachyon.visitedSplitter[s]) == 0 {
				finished[s] = true
			}
		}
	}
	fmt.Println(amount[firstSplitter] + 1)
}

func printAmount(tachyon tachyonGrid, amount map[[2]int]int) {
	for i := 0; i < tachyon.grid.Borders[1]; i++ {
		str := ""
		for j := 0; j < tachyon.grid.Borders[0]; j++ {
			if amount[[2]int{j, i}] > 0 {
				str += strconv.Itoa(amount[[2]int{j, i}])
			} else {
				str += "."
			}
		}
		fmt.Println(str)
	}
}

func (t *tachyonGrid) step() {
	curr := [2]int{0, 0}
	for p, _ := range t.currentSplitter {
		delete(t.currentSplitter, p)
		if t.visitedSplitter[p] == nil {
			curr = p
			break
		}
	}
	currPos := [][2]int{{curr[0] + 1, curr[1]}, {curr[0] - 1, curr[1]}}
	for _, p := range currPos {
		i := p[1]
		for i < t.grid.Borders[1] {
			if t.grid.Points[[2]int{p[0], i}].Symbol == '^' {
				t.visitedSplitter[curr] = append(t.visitedSplitter[curr], [2]int{p[0], i})
				t.currentSplitter[[2]int{p[0], i}] = true
				break
			}
			i++
		}
		if i == t.grid.Borders[1] {
			t.visitedSplitter[curr] = append(t.visitedSplitter[curr], [2]int{-1, -1})
		}
	}
}
