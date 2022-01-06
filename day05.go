package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func advent051(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner
	const rowCount int = 1000
	const colCount int = 1000
	var ventsCount [rowCount][colCount]int
	var x1, y1, x2, y2 int
	var coordinates []string
	var dangerCount int = 0

	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			ventsCount[i][j] = 0
		}
	}

	fd, err = os.Open(filename)
	check(err)

	scanner = bufio.NewScanner(fd)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, " -> ", ",", -1)
		coordinates = strings.Split(line, ",")
		x1, err = strconv.Atoi(coordinates[0])
		check(err)
		y1, err = strconv.Atoi(coordinates[1])
		check(err)
		x2, err = strconv.Atoi(coordinates[2])
		check(err)
		y2, err = strconv.Atoi(coordinates[3])
		check(err)
		fmt.Printf("Vent (%d,%d)->(%d,%d)", x1, y1, x2, y2)
		if x1 == x2 {
			fmt.Printf(" horizontal of %d steps\n", max(y1, y2)-min(y1, y2))
			for i := min(y1, y2); i <= max(y1, y2); i++ {
				ventsCount[x1][i]++
				if ventsCount[x1][i] == 2 {
					dangerCount++
				}
			}
		} else if y1 == y2 {
			fmt.Printf(" vertical of %d steps\n", max(x1, x2)-min(x1, x2))
			for i := min(x1, x2); i <= max(x1, x2); i++ {
				ventsCount[i][y1]++
				if ventsCount[i][y1] == 2 {
					dangerCount++
				}
			}
		} else {
			fmt.Printf(" diagonal not to be considered\n")
		}
	}

	fd.Close()

	fmt.Println(dangerCount, "Danger Fields detected")
}

func advent052(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner
	const rowCount int = 1000
	const colCount int = 1000
	var ventsCount [rowCount][colCount]int
	var x1, y1, x2, y2 int
	var coordinates []string
	var dangerCount int = 0
	var stepX, stepY int

	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			ventsCount[i][j] = 0
		}
	}

	fd, err = os.Open(filename)
	check(err)

	scanner = bufio.NewScanner(fd)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, " -> ", ",", -1)
		coordinates = strings.Split(line, ",")
		x1, err = strconv.Atoi(coordinates[0])
		check(err)
		y1, err = strconv.Atoi(coordinates[1])
		check(err)
		x2, err = strconv.Atoi(coordinates[2])
		check(err)
		y2, err = strconv.Atoi(coordinates[3])
		check(err)
		fmt.Printf("Vent (%d,%d)->(%d,%d)", x1, y1, x2, y2)
		if x1 == x2 {
			fmt.Printf(" horizontal of %d steps\n", absolute(y2-y1))
		} else if y1 == y2 {
			fmt.Printf(" vertical of %d steps\n", absolute(x2-x1))
		} else {
			fmt.Printf(" straight diagonal of %d steps\n", absolute(y2-y1))
		}

		if x2 == x1 {
			stepX = 0
		} else if x2 > x1 {
			stepX = 1
		} else {
			stepX = -1
		}
		if y2 == y1 {
			stepY = 0
		} else if y2 > y1 {
			stepY = 1
		} else {
			stepY = -1
		}

		for i := 0; i <= max(absolute(x2-x1), absolute(y2-y1)); i++ {
			ventsCount[x1+i*stepX][y1+i*stepY]++
			if ventsCount[x1+i*stepX][y1+i*stepY] == 2 {
				dangerCount++
			}
		}
	}

	fd.Close()

	fmt.Println(dangerCount, "Danger Fields detected")
}
