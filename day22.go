package main

import (
	"fmt"
	"strconv"
	"strings"
)

func advent221(input []string) []string {
	var err error
	output := []string{}
	const reactorSize = 101
	const offset = 50
	var reactor [reactorSize][reactorSize][reactorSize]bool
	for x := 0; x < reactorSize; x++ {
		for y := 0; y < reactorSize; y++ {
			for z := 0; z < reactorSize; z++ {
				reactor[x][y][z] = false
			}
		}
	}

	for _, inputLine := range input {
		line := inputLine
		lit := false
		if line[:3] == "on " {
			lit = true
			line = strings.Replace(line, "on ", "", -1)
		} else if line[:4] == "off " {
			lit = false
			line = strings.Replace(line, "off ", "", -1)
		}

		startX := 0
		endX := 0
		startY := 0
		endY := 0
		startZ := 0
		endZ := 0
		parts := strings.Split(line, ",")
		for _, part := range parts {
			if part[0] == 'x' {
				part = strings.Replace(part, "x=", "", -1)
				rangeIndex := strings.Split(part, "..")
				startX, err = strconv.Atoi(rangeIndex[0])
				check(err)
				endX, err = strconv.Atoi(rangeIndex[1])
				check(err)
			} else if part[0] == 'y' {
				part = strings.Replace(part, "y=", "", -1)
				rangeIndex := strings.Split(part, "..")
				startY, err = strconv.Atoi(rangeIndex[0])
				check(err)
				endY, err = strconv.Atoi(rangeIndex[1])
				check(err)
			} else if part[0] == 'z' {
				part = strings.Replace(part, "z=", "", -1)
				rangeIndex := strings.Split(part, "..")
				startZ, err = strconv.Atoi(rangeIndex[0])
				check(err)
				endZ, err = strconv.Atoi(rangeIndex[1])
				check(err)
			}
		}

		for x := max(startX+offset, 0); x <= min(endX+offset, reactorSize-1); x++ {
			for y := max(startY+offset, 0); y <= min(endY+offset, reactorSize-1); y++ {
				for z := max(startZ+offset, 0); z <= min(endZ+offset, reactorSize-1); z++ {
					reactor[x][y][z] = lit
				}
			}
		}
	}

	litCount := 0
	for x := 0; x < reactorSize; x++ {
		for y := 0; y < reactorSize; y++ {
			for z := 0; z < reactorSize; z++ {
				if reactor[x][y][z] {
					litCount++
				}
			}
		}
	}

	output = append(output, fmt.Sprintln("There are", litCount, "On"))

	return output
}

func advent222(input []string) []string {
	var err error
	output := []string{}
	lits := []bool{}
	startXs := []int{}
	endXs := []int{}
	startYs := []int{}
	endYs := []int{}
	startZs := []int{}
	endZs := []int{}
	litCount := 0

	for _, inputLine := range input {
		line := inputLine
		output = append(output, fmt.Sprintln(line))

		lit := false
		if line[:3] == "on " {
			lit = true
			line = strings.Replace(line, "on ", "", -1)
		} else if line[:4] == "off " {
			lit = false
			line = strings.Replace(line, "off ", "", -1)
		}

		startX := 0
		endX := 0
		startY := 0
		endY := 0
		startZ := 0
		endZ := 0
		parts := strings.Split(line, ",")
		for _, part := range parts {
			if part[0] == 'x' {
				part = strings.Replace(part, "x=", "", -1)
				rangeIndex := strings.Split(part, "..")
				startX, err = strconv.Atoi(rangeIndex[0])
				check(err)
				endX, err = strconv.Atoi(rangeIndex[1])
				check(err)
			} else if part[0] == 'y' {
				part = strings.Replace(part, "y=", "", -1)
				rangeIndex := strings.Split(part, "..")
				startY, err = strconv.Atoi(rangeIndex[0])
				check(err)
				endY, err = strconv.Atoi(rangeIndex[1])
				check(err)
			} else if part[0] == 'z' {
				part = strings.Replace(part, "z=", "", -1)
				rangeIndex := strings.Split(part, "..")
				startZ, err = strconv.Atoi(rangeIndex[0])
				check(err)
				endZ, err = strconv.Atoi(rangeIndex[1])
				check(err)
			}
		}

		for i := 0; i < len(lits); i++ {
			overlapStartX := max(startX, startXs[i])
			overlapEndX := min(endX, endXs[i])
			overlapStartY := max(startY, startYs[i])
			overlapEndY := min(endY, endYs[i])
			overlapStartZ := max(startZ, startZs[i])
			overlapEndZ := min(endZ, endZs[i])
			if overlapEndX >= overlapStartX && overlapEndY >= overlapStartY && overlapEndZ >= overlapStartZ {
				lits = append(lits[:i+1], lits[i:]...)
				startXs = append(startXs[:i+1], startXs[i:]...)
				startXs[i+1] = overlapStartX
				endXs = append(endXs[:i+1], endXs[i:]...)
				endXs[i+1] = overlapEndX
				startYs = append(startYs[:i+1], startYs[i:]...)
				startYs[i+1] = overlapStartY
				endYs = append(endYs[:i+1], endYs[i:]...)
				endYs[i+1] = overlapEndY
				startZs = append(startZs[:i+1], startZs[i:]...)
				startZs[i+1] = overlapStartZ
				endZs = append(endZs[:i+1], endZs[i:]...)
				endZs[i+1] = overlapEndZ
				if lits[i] {
					lits[i+1] = false
					litCount = litCount - (overlapEndX-overlapStartX+1)*(overlapEndY-overlapStartY+1)*(overlapEndZ-overlapStartZ+1)
					output = append(output, fmt.Sprintln("Removing an overlap Cube of", (overlapEndX-overlapStartX+1), "x", (overlapEndY-overlapStartY+1), "x", (overlapEndZ-overlapStartZ+1), "with", (overlapEndX-overlapStartX+1)*(overlapEndY-overlapStartY+1)*(overlapEndZ-overlapStartZ+1), "single cubes"))
					output = append(output, fmt.Sprintln("Current count is", litCount))
				} else if !lits[i] {
					lits[i+1] = true
					litCount = litCount + (overlapEndX-overlapStartX+1)*(overlapEndY-overlapStartY+1)*(overlapEndZ-overlapStartZ+1)
					output = append(output, fmt.Sprintln("Adding an overlap Cube of", (overlapEndX-overlapStartX+1), "x", (overlapEndY-overlapStartY+1), "x", (overlapEndZ-overlapStartZ+1), "with", (overlapEndX-overlapStartX+1)*(overlapEndY-overlapStartY+1)*(overlapEndZ-overlapStartZ+1), "single cubes"))
					output = append(output, fmt.Sprintln("Current count is", litCount))
				}
				i++
			}
		}
		if lit {
			lits = append(lits, lit)
			startXs = append(startXs, startX)
			endXs = append(endXs, endX)
			startYs = append(startYs, startY)
			endYs = append(endYs, endY)
			startZs = append(startZs, startZ)
			endZs = append(endZs, endZ)
			litCount = litCount + (endX-startX+1)*(endY-startY+1)*(endZ-startZ+1)
			output = append(output, fmt.Sprintln("Adding a Cube of", (endX-startX+1), "x", (endY-startY+1), "x", (endZ-startZ+1), "with", (endX-startX+1)*(endY-startY+1)*(endZ-startZ+1), "single cubes"))
			output = append(output, fmt.Sprintln("Current count is", litCount))
		}
	}

	output = append(output, fmt.Sprintln("There are", litCount, "On"))

	return output
}
