package main

import (
	"AOC2025/helper"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Distance struct {
	start  helper.Point[[3]int]
	end    helper.Point[[3]int]
	length float64
}

func main() {
	//args := os.Args[1:]
	file, err := os.ReadFile("part1")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\r\n")
	grid := getGrid(lines)
	distances := getDistances(grid)
	n := 10000000
	part1(distances, len(grid.Points), n)
}

func part1(distances []Distance, pointAmount int, n int) {
	posToNetworks := map[[3]int]int{}
	posToNetworks[distances[0].start.Position] = 1
	posToNetworks[distances[0].end.Position] = 1
	networkCounter := 2
	if n > len(distances) {
		n = len(distances)
	}
	part2 := Distance{}
	for _, d := range distances[1:n] {
		if posToNetworks[d.start.Position] != 0 && posToNetworks[d.end.Position] != 0 {
			newNetwork := posToNetworks[d.start.Position]
			oldNetwork := posToNetworks[d.end.Position]
			for pos, nw := range posToNetworks {
				if nw == oldNetwork {
					posToNetworks[pos] = newNetwork
				}
			}
		}
		if posToNetworks[d.start.Position] != 0 {
			posToNetworks[d.end.Position] = posToNetworks[d.start.Position]
		}
		if posToNetworks[d.end.Position] != 0 {
			posToNetworks[d.start.Position] = posToNetworks[d.end.Position]
		}
		if posToNetworks[d.start.Position] == 0 && posToNetworks[d.end.Position] == 0 {
			posToNetworks[d.start.Position] = networkCounter
			posToNetworks[d.end.Position] = networkCounter
			networkCounter++
		}
		if checkAllNetworkConnected(posToNetworks, pointAmount) {
			part2 = d
			break
		}
	}
	netWorkCount := map[int][][3]int{}
	for p, nw := range posToNetworks {
		netWorkCount[nw] = append(netWorkCount[nw], p)
	}

	sliceNetWorkCount := []int{}
	for _, v := range netWorkCount {
		sliceNetWorkCount = append(sliceNetWorkCount, len(v))
	}
	if part2.length != 0 {
		fmt.Println(part2.start.Position[0] * part2.end.Position[0])
	} else {
		fmt.Println(sliceNetWorkCount[len(sliceNetWorkCount)-1] * sliceNetWorkCount[len(sliceNetWorkCount)-2] * sliceNetWorkCount[len(sliceNetWorkCount)-3])
	}
}

func checkAllNetworkConnected(posToNetworks map[[3]int]int, pointAmount int) bool {
	firstNw := 0
	if len(posToNetworks) != pointAmount {
		return false
	}
	for _, nw := range posToNetworks {
		firstNw = nw
		break
	}
	for _, nw := range posToNetworks {
		if nw != firstNw {
			return false
		}
	}
	return true
}

func getGrid(lines []string) helper.Grid[[3]int] {
	var points [][3]int
	for _, line := range lines {
		numberStrings := strings.Split(line, ",")
		point := [3]int{}
		for i, numberString := range numberStrings {
			number, _ := strconv.Atoi(numberString)
			point[i] = number
		}
		points = append(points, point)
	}
	grid := helper.GetGrid(points, '.')
	return grid
}

func getDistances(grid helper.Grid[[3]int]) []Distance {
	distances := []Distance{}
	for _, p1 := range grid.Points {
		for _, p2 := range grid.Points {
			if p1.Position != p2.Position {
				dis := helper.EuclidianDistance(p1, p2)
				distance := Distance{p1, p2, dis}
				distances = append(distances, distance)
			}
		}
	}
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].length < distances[j].length
	})
	newDistances := []Distance{}
	for i := 0; i < len(distances); i++ {
		if i%2 == 0 {
			newDistances = append(newDistances, distances[i])
		}
	}
	distances = newDistances
	return distances
}
