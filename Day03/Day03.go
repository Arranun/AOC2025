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
	lines := strings.Split(string(file), "\r\n")
	sum := 0
	remaningN, _ := strconv.Atoi(args[1])
	for _, line := range lines {
		number := getNumber(line, remaningN)
		sum += number
	}
	fmt.Println(sum)
}

func getNumber(line string, remaningN int) int {
	digits := []int{}
	index := 0
	n := 0
	for remaningN >= 0 {
		index = findDigitN(&digits, n, line, index, remaningN)
		n++
		remaningN--
	}
	number := 0
	for i := len(digits) - 1; i > -1; i-- {
		multp := 1
		for j := 0; j < len(digits)-1-i; j++ {
			multp *= 10
		}
		number += digits[i] * multp
	}
	return number
}

func findDigitN(digits *[]int, n int, line string, index int, remaningN int) int {
	*digits = append(*digits, 0)
	(*digits)[n] = 0
	for i := index; i < len(line)-remaningN; i++ {
		pos := int(line[i] - '0')
		if pos > (*digits)[n] {
			(*digits)[n] = pos
			index = i
		}
	}
	return index + 1
}
