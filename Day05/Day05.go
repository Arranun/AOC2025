package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//args := os.Args[1:]
	file, err := os.ReadFile("part1")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\r\n")
	ranges := map[[2]int]bool{}
	i := 0
	for {
		if lines[i] == "" {
			break
		}
		rangeStrings := strings.Split(lines[i], "-")
		start, _ := strconv.Atoi(rangeStrings[0])
		end, _ := strconv.Atoi(rangeStrings[1])
		newRange := [2]int{start, end}
		for r, _ := range ranges {
			if r[0] <= newRange[0] && r[1] >= newRange[1] {
				delete(ranges, r)
				newRange = [2]int{r[0], r[1]}
			}
			if r[0] <= newRange[0] && r[1] >= newRange[0] {
				delete(ranges, r)
				newRange = [2]int{r[0], newRange[1]}
			}
			if r[0] <= newRange[1] && r[1] >= newRange[1] {
				delete(ranges, r)
				newRange = [2]int{newRange[0], r[1]}
			}
		}
		ranges[newRange] = true
		i++
	}
	for cr, _ := range ranges {
		newRange := cr
		delete(ranges, cr)
		for r, _ := range ranges {
			if r[0] <= newRange[0] && r[1] >= newRange[1] {
				delete(ranges, r)
				newRange = [2]int{r[0], r[1]}
			}
			if r[0] <= newRange[0] && r[1] >= newRange[0] {
				delete(ranges, r)
				newRange = [2]int{r[0], newRange[1]}
			}
			if r[0] <= newRange[1] && r[1] >= newRange[1] {
				delete(ranges, r)
				newRange = [2]int{newRange[0], r[1]}
			}
		}
		ranges[newRange] = true
	}
	ingredients := []int{}
	for i < len(lines) {
		id, _ := strconv.Atoi(lines[i])
		ingredients = append(ingredients, id)
		i++
	}
	part1 := 0
	for _, ingredient := range ingredients {
		for r, _ := range ranges {
			if ingredient >= r[0] && ingredient <= r[1] {
				part1++
				break
			}
		}
	}
	fmt.Println(part1)

	part2 := 0
	startEnd := map[int]int{}
	for r, _ := range ranges {
		startEnd[r[0]]++
		part2 += r[1] - r[0] + 1
	}
	fmt.Println(ranges)
	for a, b := range startEnd {
		if b > 1 {
			fmt.Println(a, b)
		}
	}
	fmt.Println(part2)
}
