package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func advent061(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner
	const simulationDays int = 80
	var fish int
	var fishes []int
	var newCount int = 0

	fd, err = os.Open(filename)
	check(err)

	scanner = bufio.NewScanner(fd)
	scanner.Scan()
	line := scanner.Text()
	lineParts := strings.Split(line, ",")
	for _, s := range lineParts {
		fish, err = strconv.Atoi(s)
		check(err)
		fishes = append(fishes, fish)
	}
	fd.Close()

	fmt.Println("Starting fish population is", len(fishes))

	for i := 1; i <= simulationDays; i++ {
		newCount = 0
		for j := 0; j < len(fishes); j++ {
			fishes[j]--
			if fishes[j] < 0 {
				fishes[j] = 6
				newCount++
			}
		}
		fmt.Println("On day", i, ", there were", len(fishes), "fishes and", newCount, "were born")
		for j := 1; j <= newCount; j++ {
			fishes = append(fishes, 8)
		}
	}

	fmt.Println("Final fish population is", len(fishes))
}

func advent062(filename string) {
	var getNewFish func(days int, calculations []int) int
	getNewFish = func(days int, calculations []int) int {
		const bornCycle int = 7
		const maturityTime int = 2
		var newDirectFishes int = 0
		var newTotal int = 0

		if days < 0 {
			return 0
		}
		if calculations[days] > -1 {
			return calculations[days]
		}

		newDirectFishes = max((days+2*bornCycle-1)/bornCycle-1, 0)
		for i := 0; i < newDirectFishes; i++ {
			newTotal = newTotal + getNewFish(days-i*bornCycle-bornCycle-maturityTime, calculations)
		}
		newTotal = newTotal + newDirectFishes

		calculations[days] = newTotal
		return newTotal
	}

	var fd *os.File
	var err error
	var scanner *bufio.Scanner
	const simulationDays int = 256
	var fish int
	var fishes []int
	var totalCount int = 0
	var newFishes int
	var calculations []int

	fd, err = os.Open(filename)
	check(err)

	scanner = bufio.NewScanner(fd)
	scanner.Scan()
	line := scanner.Text()
	lineParts := strings.Split(line, ",")
	for _, s := range lineParts {
		fish, err = strconv.Atoi(s)
		check(err)
		fishes = append(fishes, fish)
	}
	fd.Close()

	fmt.Println("Starting fish population is", len(fishes))

	for i := 0; i < simulationDays; i++ {
		calculations = append(calculations, -1)
	}
	totalCount = len(fishes)
	for i := 0; i < len(fishes); i++ {
		newFishes = getNewFish(simulationDays-fishes[i], calculations)
		totalCount = totalCount + newFishes
		fmt.Println("Fish", i, ", generates", newFishes, "new fishes")
	}

	fmt.Println("Final fish population is", totalCount)
}
