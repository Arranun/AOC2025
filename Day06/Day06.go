package main

import (
	"fmt"
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
	part2 := false
	if args[1] == "2" {
		part2 = true
	}
	fmt.Println(part2)
	lines := strings.Split(string(file), "\r\n")
	rows := [][]string{}
	cutIndexes := []int{}
	for i, r := range lines[len(lines)-1] {
		if r != ' ' {
			cutIndexes = append(cutIndexes, i)
		}
	}
	for _, line := range lines {
		row := []string{}
		for i := 0; i < len(cutIndexes)-1; i++ {
			row = append(row, line[cutIndexes[i]:cutIndexes[i+1]])
		}
		row = append(row, line[cutIndexes[len(cutIndexes)-1]:]+" ")
		rows = append(rows, row)
	}
	sum := 0
	for i := 0; i < len(rows[0]); i++ {
		rowSum := 0
		op := strings.TrimSpace(rows[len(rows)-1][i])
		numbers := []int{}
		if !part2 {
			numbers = getNumbersPart1(rows, i, numbers)
		} else {
			numbers = getNumbersPart2(rows, i, numbers)
		}
		if op == "*" {
			rowSum = 1
		}
		for _, number := range numbers {
			if op == "*" {
				rowSum *= number
			} else {
				rowSum += number
			}
		}
		sum += rowSum
	}
	fmt.Println(sum)
}

func getNumbersPart1(rows [][]string, i int, numbers []int) []int {
	for j := 0; j < len(rows)-1; j++ {
		number, _ := strconv.Atoi(strings.TrimSpace(rows[j][i]))
		numbers = append(numbers, number)
	}
	return numbers
}

func getNumbersPart2(rows [][]string, i int, numbers []int) []int {
	numbersStrings := []string{}
	for k := len(rows) - 2; k > -1; k-- {
		numbersStrings = append(numbersStrings, rows[k][i])
	}
	for h := len(numbersStrings[0]) - 2; h > -1; h-- {
		number := 0
		digitIndex := 0
		for _, s := range numbersStrings {
			if s[h] != ' ' {
				n := int(s[h] - '0')
				for k := 0; k < digitIndex; k++ {
					n *= 10
				}
				number += n
				digitIndex++
			}
		}
		numbers = append(numbers, number)
	}
	return numbers
}
