package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type machine struct {
	registers     map[string]int
	indicators    []bool
	indicatorGoal []bool
	buttons       [][]int
	voltages      []int
}

func main() {
	//args := os.Args[1:]
	file, err := os.ReadFile("testInput")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\r\n")
	var machines []machine
	for _, line := range lines {
		endIndictors := strings.Index(line, "]")
		m := machine{}
		for i := 1; i < endIndictors; i++ {
			m.indicators = append(m.indicators, false)
			m.indicatorGoal = append(m.indicatorGoal, line[i] == '#')
		}
		beginVoltages := strings.Index(line, "{")
		stringsButtongs := strings.Split(line[endIndictors+2:beginVoltages-1], " ")
		for _, stringButton := range stringsButtongs {
			button := []int{}
			stringsButton := strings.Split(stringButton[1:len(stringButton)-1], ",")
			for _, s := range stringsButton {
				number, _ := strconv.Atoi(s)
				button = append(button, number)
			}
			m.buttons = append(m.buttons, button)
		}
		stringsVoltage := strings.Split(line[beginVoltages+1:len(line)-1], ",")
		for _, s := range stringsVoltage {
			number, _ := strconv.Atoi(s)
			m.voltages = append(m.voltages, number)
		}
		machines = append(machines, m)
	}
	for _, m := range machines {
		fmt.Println(m)
	}
}
