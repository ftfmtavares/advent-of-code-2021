package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func advent071(input []string) []string {
	var err error
	output := []string{}
	var crab int
	var crabs []int
	var fuel int
	var position int

	lineParts := strings.Split(input[0], ",")
	for _, s := range lineParts {
		crab, err = strconv.Atoi(s)
		check(err)
		crabs = append(crabs, crab)
	}

	output = append(output, fmt.Sprintln("Number of Crabs is", len(crabs)))

	sort.Ints(crabs)

	if len(crabs)%2 == 0 {
		position = (crabs[len(crabs)/2-1] + crabs[len(crabs)/2]) / 2
	} else {
		position = crabs[(len(crabs)-1)/2]
	}
	output = append(output, fmt.Sprintln("Best position is", position))

	fuel = 0
	position = crabs[len(crabs)/2]
	for i := 0; i < len(crabs); i++ {
		if crabs[i] > position {
			fuel = fuel + crabs[i] - position
		} else {
			fuel = fuel - crabs[i] + position
		}
	}

	output = append(output, fmt.Sprintln("The amount of fuel needed for the best position is", fuel))

	return output
}

func advent072(input []string) []string {
	var err error
	output := []string{}
	var crab int
	var crabs []int
	var fuel int
	var distance int
	var bestFuel int
	var position int
	var max int
	var min int

	lineParts := strings.Split(input[0], ",")
	for _, s := range lineParts {
		crab, err = strconv.Atoi(s)
		check(err)
		crabs = append(crabs, crab)
	}

	output = append(output, fmt.Sprintln("Number of Crabs is", len(crabs)))

	max = crabs[0]
	min = crabs[0]
	for i := 1; i < len(crabs); i++ {
		if crabs[i] > max {
			max = crabs[i]
		}
		if crabs[i] < min {
			min = crabs[i]
		}
	}
	position = min
	bestFuel = -1
	for i := min; i <= max; i++ {
		fuel = 0
		for j := 0; j < len(crabs); j++ {
			if crabs[j] > i {
				distance = crabs[j] - i
			} else {
				distance = i - crabs[j]
			}
			for k := distance - 1; k > 0; k-- {
				distance = distance + k
			}
			fuel = fuel + distance
		}
		if fuel < bestFuel || bestFuel == -1 {
			position = i
			bestFuel = fuel
		}
	}
	output = append(output, fmt.Sprintln("Best position is", position))
	output = append(output, fmt.Sprintln("The amount of fuel needed for the best position is", bestFuel))

	return output
}
