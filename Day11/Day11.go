package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//args := os.Args[1:]
	file, err := os.ReadFile("part2TestInput")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\r\n")
	connections := map[string]map[string]bool{}
	for _, line := range lines {
		leftRight := strings.Split(line, ": ")
		connections[leftRight[0]] = make(map[string]bool)
		for _, connection := range strings.Split(leftRight[1], " ") {
			connections[leftRight[0]][connection] = true
		}
	}
	part1(connections, "you", "out")
	connections = map[string]map[string]bool{}
	possibilities := map[int]bool{}
	for i := 0; i < 200; i++ {
		for _, line := range lines {
			leftRight := strings.Split(line, ": ")
			connections[leftRight[0]] = make(map[string]bool)
			for _, connection := range strings.Split(leftRight[1], " ") {
				connections[leftRight[0]][connection] = true
			}
		}
		result := part1(connections, "svr", "out")
		possibilities[result] = true
	}
	for p, _ := range possibilities {
		fmt.Println(p)
	}

}

func part1(connections map[string]map[string]bool, start string, goal string) int {
	amount := make(map[string][4]int)
	finished := make(map[string]bool)
	for connection, to := range connections {
		for toConnection := range to {
			if toConnection == goal {
				delete(connections[connection], toConnection)
				amount[connection] = [4]int{1, 0, 0}
				if len(connections[connection]) == 0 {
					finished[connection] = true
					delete(connections, connection)
				}
			}
		}
	}
	for len(finished) > 0 {
		finishConnection := ""
		for f, _ := range finished {
			finishConnection = f
			break
		}
		for f, _ := range finished {
			if amount[f][3] > amount[finishConnection][3] {
				finishConnection = f
			}
		}
		delete(finished, finishConnection)

		for connection, to := range connections {
			if to[finishConnection] {
				newAmount := [4]int{amount[connection][0] + amount[finishConnection][0], amount[connection][1] + amount[finishConnection][1], amount[connection][2] + amount[finishConnection][2], amount[connection][3] + amount[finishConnection][3]}
				if finishConnection == "dac" {
					newAmount[1] += newAmount[0]
					newAmount[3] += newAmount[2]
					newAmount[0] = 0
					newAmount[2] = 0
				}
				if finishConnection == "fft" {
					newAmount[2] += newAmount[0]
					newAmount[3] += newAmount[1]
					newAmount[0] = 0
					newAmount[1] = 0
				}
				amount[connection] = newAmount
				delete(to, finishConnection)
				if len(to) == 0 {
					finished[connection] = true
					delete(connections, connection)
				}
			}
		}
	}
	fmt.Println(amount[start])
	return amount[start][3]
}
