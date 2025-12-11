package main

import (
	"AOC2025/helper"
	"fmt"
	"os"
	"sort"
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
	numbers := [][2]int{}
	for _, line := range lines {
		numberStrings := strings.Split(line, ",")
		number := [2]int{}
		number[0], _ = strconv.Atoi(numberStrings[0])
		number[1], _ = strconv.Atoi(numberStrings[1])
		numbers = append(numbers, number)
	}
	currNumber := numbers[0]
	downLines, rightLines := getLines(numbers, currNumber)
	grid := helper.GetGrid(numbers, '#')
	distances := grid.GetDistance(helper.ManHattanDistance[[2]int])
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Length < distances[j].Length
	})
	for i := len(distances) - 1; i > -1; i-- {
		checkCurveDistance := getCheckCurveDistance(distances, i)
		if checkIfInCurve(checkCurveDistance, downLines, rightLines) {
			fmt.Println(getArea(distances[i]))
			break
		}
	}

}

func getCheckCurveDistance(distances []helper.Distance[[2]int], i int) helper.Distance[[2]int] {
	checkCurveDistance := distances[i]
	for n := 0; n < 2; n++ {
		if checkCurveDistance.Start.Position[n] < checkCurveDistance.End.Position[n] {
			checkCurveDistance.Start.Position[n]++
			checkCurveDistance.End.Position[n]--
		}
		if checkCurveDistance.Start.Position[n] > checkCurveDistance.End.Position[n] {
			checkCurveDistance.Start.Position[n]--
			checkCurveDistance.End.Position[n]++
		}
	}
	return checkCurveDistance
}

func getLines(numbers [][2]int, currNumber [2]int) (map[[3]int]bool, map[[3]int]bool) {
	downLines := map[[3]int]bool{}
	rightLines := map[[3]int]bool{}
	numbers = append(numbers, currNumber)
	for _, number := range numbers[1:] {
		diff := [2]int{number[0] - currNumber[0], number[1] - currNumber[1]}
		if diff[0] == 0 {
			downLines[[3]int{number[0], number[1], currNumber[1]}] = true
		} else {
			rightLines[[3]int{number[1], number[0], currNumber[0]}] = true
		}
		currNumber = number
	}
	for dL, _ := range downLines {
		delete(downLines, dL)
		downLines[orderLine(dL)] = true
	}
	for rL, _ := range rightLines {
		delete(rightLines, rL)
		rightLines[orderLine(rL)] = true

	}

	return downLines, rightLines
}

func orderLine(line [3]int) [3]int {
	if line[1] > line[2] {
		return [3]int{line[0], line[2], line[1]}
	}
	return line
}

func checkIfInCurve(distance helper.Distance[[2]int], downLines map[[3]int]bool, rightLines map[[3]int]bool) bool {
	var dLs [][3]int
	dLs = append(dLs, orderLine([3]int{distance.Start.Position[0], distance.Start.Position[1], distance.End.Position[1]}))
	dLs = append(dLs, orderLine([3]int{distance.End.Position[0], distance.Start.Position[1], distance.End.Position[1]}))
	var rLs [][3]int
	rLs = append(rLs, orderLine([3]int{distance.Start.Position[1], distance.Start.Position[0], distance.End.Position[0]}))
	rLs = append(rLs, orderLine([3]int{distance.End.Position[1], distance.Start.Position[0], distance.End.Position[0]}))
	if checkForHits(dLs, rightLines) {
		return false
	}
	if checkForHits(rLs, downLines) {
		return false
	}
	return true
}

func checkForHits(dLs [][3]int, rightLines map[[3]int]bool) bool {
	for _, dl := range dLs {
		affectedLines := [][3]int{}
		for rL, _ := range rightLines {
			if dl[0] >= rL[1] && dl[0] <= rL[2] {
				affectedLines = append(affectedLines, rL)
			}
		}
		hitLines := [][3]int{}

		for _, al := range affectedLines {
			if al[0] >= dl[1] && al[0] <= dl[2] {
				hitLines = append(hitLines, al)
			}
		}

		hits := len(hitLines)
		if hits > 0 {
			return true
		}
	}
	return false
}

func part1(distances []helper.Distance[[2]int]) {
	highestDistance := distances[len(distances)-1]
	fmt.Println(highestDistance)
	fmt.Println(getArea(highestDistance))
}

func getArea(highestDistance helper.Distance[[2]int]) int {
	width := highestDistance.End.Position[0] - highestDistance.Start.Position[0]
	if width < 0 {
		width *= -1
	}
	height := highestDistance.End.Position[1] - highestDistance.Start.Position[1]
	if height < 0 {
		height *= -1
	}
	return (width + 1) * (height + 1)
}
