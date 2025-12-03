package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	file, err := os.ReadFile(args[0])
	if err != nil {
		panic(err)
	}
	ranges := strings.Split(string(file), ",")
	part1 := getRepeatIDs(ranges, false)
	fmt.Println(part1)
}

func getRepeatIDs(ranges []string, part1 bool) int {
	sum := 0
	for _, r := range ranges {
		startEnd := strings.Split(r, "-")
		length := 0
		if len(startEnd[0])%2 == 0 {
			length = len(startEnd[0]) / 2
		} else {
			length = len(startEnd[1]) / 2
		}
		start, _ := strconv.Atoi(startEnd[0])
		end, _ := strconv.Atoi(startEnd[1])
		hits := map[int]bool{}
		for i := length; i > 0; i-- {
			getHits(&hits, i, len(startEnd[1]), start, end, part1)
		}
		for k := range hits {
			sum += k
		}
		fmt.Println(startEnd[0], startEnd[1])
		fmt.Println(hits)
	}
	return sum
}

func getHits(hits *map[int]bool, length int, maxLength int, start int, end int, part1 bool) {
	n := math.Pow(10, float64(length))
	for i := n / 10; i < n; i++ {
		partNumber := strconv.Itoa(int(i))
		stringNumber := ""
		repeats := 2
		for len(stringNumber) < maxLength {
			stringNumber = strings.Repeat(partNumber, repeats)
			numberNumber, _ := strconv.Atoi(stringNumber)
			if numberNumber >= start && numberNumber <= end {
				(*hits)[numberNumber] = true
			}
			if part1 {
				break
			}
			repeats++
		}

	}
}
