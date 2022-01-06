package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func advent071(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner
	var crab int
	var crabs []int
	var fuel int
	var position int

	fd, err = os.Open(filename)
	check(err)

	scanner = bufio.NewScanner(fd)
	scanner.Scan()
	lineParts := strings.Split(scanner.Text(), ",")
	for _, s := range lineParts {
		crab, err = strconv.Atoi(s)
		check(err)
		crabs = append(crabs, crab)
	}
	fd.Close()

	fmt.Println("Number of Crabs is", len(crabs))

	sort.Ints(crabs)

	if len(crabs)%2 == 0 {
		position = (crabs[len(crabs)/2-1] + crabs[len(crabs)/2]) / 2
	} else {
		position = crabs[(len(crabs)-1)/2]
	}
	fmt.Println("Best position is", position)

	fuel = 0
	position = crabs[len(crabs)/2]
	for i := 0; i < len(crabs); i++ {
		if crabs[i] > position {
			fuel = fuel + crabs[i] - position
		} else {
			fuel = fuel - crabs[i] + position
		}
	}

	fmt.Println("The amount of fuel needed for the best position is", fuel)
}

func advent072(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner
	var crab int
	var crabs []int
	var fuel int
	var distance int
	var bestFuel int
	var position int
	var max int
	var min int

	fd, err = os.Open(filename)
	check(err)

	scanner = bufio.NewScanner(fd)
	scanner.Scan()
	lineParts := strings.Split(scanner.Text(), ",")
	for _, s := range lineParts {
		crab, err = strconv.Atoi(s)
		check(err)
		crabs = append(crabs, crab)
	}
	fd.Close()

	fmt.Println("Number of Crabs is", len(crabs))

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
	fmt.Println("Best position is", position)
	fmt.Println("The amount of fuel needed for the best position is", bestFuel)
}
