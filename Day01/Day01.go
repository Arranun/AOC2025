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
	pos := 50
	cnt := 0
	part1 := 0
	part2 := 0
	for _, line := range lines {
		fmt.Println(line)
		pos, cnt = rotate(pos, line)
		fmt.Println(pos, cnt)
		if pos == 0 {
			part1++
		}
		part2 += cnt
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func rotate(pos int, command string) (int, int) {
	cnt := 0
	start := pos
	dir, steps := parseCommand(command)
	n := pos + steps*dir
	if n == 0 {
		cnt++
	}
	crossings := n / 100
	if crossings < 0 {
		crossings = -crossings
	}
	if n < 0 && start != 0 {
		crossings++
	}
	n = n % 100
	if n < 0 {
		n = 100 + n
	}
	cnt += crossings
	return n, cnt
}

func parseCommand(command string) (int, int) {
	dir := 1
	if command[0] == 'L' {
		dir = -1
	}
	steps, _ := strconv.Atoi(command[1:])
	return dir, steps
}
