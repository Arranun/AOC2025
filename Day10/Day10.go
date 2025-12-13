package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type machine struct {
	indicators    []bool
	indicatorGoal []bool
	buttons       [][]int
	voltages      []int
}

func main() {
	//args := os.Args[1:]
	file, err := os.ReadFile("part1")
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
	sum := 0
	for _, m := range machines {
		sum += m.getAmountOfBottonPresses()
	}
	fmt.Println(sum)
}

func (m machine) getAmountOfBottonPresses() int {
	relevantIndicators := make(map[int]bool)
	for i, indicator := range m.indicatorGoal {
		if indicator {
			relevantIndicators[i] = true
		}
	}
	relevantButtons := []int{}
	for i, _ := range m.buttons {
		relevantButtons = append(relevantButtons, i)
	}
	shortestRoute := len(relevantButtons) + 1
	activeRoutes := map[string][]int{}
	visited := map[string]bool{}
	finishedRoutes := map[string][]int{}
	for relevantButton := range relevantButtons {
		activeRoutes[strconv.Itoa(relevantButton)] = []int{relevantButton}
	}
	for len(activeRoutes) > 0 {
		routeString := ""
		routeCombination := []int{}
		for str, combi := range activeRoutes {
			routeString = str
			routeCombination = combi
			break
		}
		visited[routeString] = true
		delete(activeRoutes, routeString)
		if m.TestButtonCombination(routeCombination) {
			finishedRoutes[routeString] = routeCombination
			if len(routeCombination) < shortestRoute {
				shortestRoute = len(routeCombination)
			}
			continue
		}
		if len(routeCombination) == len(relevantButtons) {
			continue
		}
		currRelevantButton := []int{}
		for _, relevantButton := range relevantButtons {
			if !contains(routeCombination, relevantButton) {
				currRelevantButton = append(currRelevantButton, relevantButton)
			}
		}
		for _, button := range currRelevantButton {
			newRouteCombination := make([]int, len(routeCombination))
			copy(newRouteCombination, routeCombination)
			newRouteCombination = append(newRouteCombination, button)
			sort.Ints(newRouteCombination)
			newRouteString := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(newRouteCombination)), ","), "[]")
			if !visited[newRouteString] {
				activeRoutes[newRouteString] = newRouteCombination
			}
		}
	}
	return shortestRoute
}

func contains(routeCombination []int, relevantButton int) bool {
	for _, button := range routeCombination {
		if relevantButton == button {
			return true
		}
	}
	return false
}

func (m machine) TestButtonCombination(buttons []int) bool {
	indicator := make([]bool, len(m.indicatorGoal))
	for _, buttonIndex := range buttons {
		button := m.buttons[buttonIndex]
		for _, changedIndicator := range button {
			indicator[changedIndicator] = !indicator[changedIndicator]
		}
	}
	for i := 0; i < len(m.indicatorGoal); i++ {
		if indicator[i] != m.indicatorGoal[i] {
			return false
		}
	}
	return true
}
