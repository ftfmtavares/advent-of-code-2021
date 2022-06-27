package main

import (
	"fmt"
	"strconv"
	"strings"
)

func advent051(input []string) []string {
	var err error
	output := []string{}
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

	for _, inputLine := range input {
		line := inputLine
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
		output = append(output, fmt.Sprintf("Vent (%d,%d)->(%d,%d)", x1, y1, x2, y2))
		if x1 == x2 {
			output = append(output, fmt.Sprintf(" horizontal of %d steps\n", max(y1, y2)-min(y1, y2)))
			for i := min(y1, y2); i <= max(y1, y2); i++ {
				ventsCount[x1][i]++
				if ventsCount[x1][i] == 2 {
					dangerCount++
				}
			}
		} else if y1 == y2 {
			output = append(output, fmt.Sprintf(" vertical of %d steps\n", max(x1, x2)-min(x1, x2)))
			for i := min(x1, x2); i <= max(x1, x2); i++ {
				ventsCount[i][y1]++
				if ventsCount[i][y1] == 2 {
					dangerCount++
				}
			}
		} else {
			output = append(output, fmt.Sprintln(" diagonal not to be considered"))
		}
	}

	output = append(output, fmt.Sprintln(dangerCount, "Danger Fields detected"))

	return output
}

func advent052(input []string) []string {
	var err error
	output := []string{}
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

	for _, inputLine := range input {
		line := inputLine
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
		output = append(output, fmt.Sprintf("Vent (%d,%d)->(%d,%d)", x1, y1, x2, y2))
		if x1 == x2 {
			output = append(output, fmt.Sprintf(" horizontal of %d steps\n", absolute(y2-y1)))
		} else if y1 == y2 {
			output = append(output, fmt.Sprintf(" vertical of %d steps\n", absolute(x2-x1)))
		} else {
			output = append(output, fmt.Sprintf(" straight diagonal of %d steps\n", absolute(y2-y1)))
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

	output = append(output, fmt.Sprintln(dangerCount, "Danger Fields detected"))

	return output
}
