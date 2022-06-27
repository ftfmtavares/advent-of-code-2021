package main

import (
	"fmt"
)

func advent091(input []string) []string {
	output := []string{}
	var riskSum int = 0

	lines := []string{}
	lines = append(lines, input...)

	var lowerPoint bool
	var lowerPointLevel int
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			lowerPoint = true
			if i > 0 && lines[i][j] >= lines[i-1][j] {
				lowerPoint = false
			}
			if i < len(lines)-1 && lines[i][j] >= lines[i+1][j] {
				lowerPoint = false
			}
			if j > 0 && lines[i][j] >= lines[i][j-1] {
				lowerPoint = false
			}
			if j < len(lines[i])-1 && lines[i][j] >= lines[i][j+1] {
				lowerPoint = false
			}
			if lowerPoint {
				lowerPointLevel = int(lines[i][j] - '0')
				riskSum = riskSum + lowerPointLevel + 1
			}
		}
	}

	output = append(output, fmt.Sprintln("Sum of all risk levels is", riskSum))

	return output
}

func advent092(input []string) []string {
	var getBasinSize func(lines [][]int, x int, y int) int
	getBasinSize = func(lines [][]int, x int, y int) int {
		res := 1
		lines[x][y] = -1
		if x > 0 && lines[x-1][y] < 9 && lines[x-1][y] >= 0 {
			res = res + getBasinSize(lines, x-1, y)
		}
		if x < len(lines)-1 && lines[x+1][y] < 9 && lines[x+1][y] >= 0 {
			res = res + getBasinSize(lines, x+1, y)
		}
		if y > 0 && lines[x][y-1] < 9 && lines[x][y-1] >= 0 {
			res = res + getBasinSize(lines, x, y-1)
		}
		if y < len(lines[x])-1 && lines[x][y+1] < 9 && lines[x][y+1] >= 0 {
			res = res + getBasinSize(lines, x, y+1)
		}
		return res
	}

	output := []string{}
	var productTopBasins int = 0

	lines := [][]int{}
	for _, inputLine := range input {
		line := inputLine
		lines = append(lines, []int{})
		for i := 0; i < len(line); i++ {
			lines[len(lines)-1] = append(lines[len(lines)-1], int(line[i]-'0'))
		}
	}

	lowerPoint := true
	topBasins := [3]int{0, 0, 0}
	basinSize := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			lowerPoint = true
			if i > 0 && lines[i][j] >= lines[i-1][j] {
				lowerPoint = false
			}
			if i < len(lines)-1 && lines[i][j] >= lines[i+1][j] {
				lowerPoint = false
			}
			if j > 0 && lines[i][j] >= lines[i][j-1] {
				lowerPoint = false
			}
			if j < len(lines[i])-1 && lines[i][j] >= lines[i][j+1] {
				lowerPoint = false
			}
			if lowerPoint {
				basinSize = getBasinSize(lines, i, j)
				if basinSize > topBasins[0] {
					topBasins[2] = topBasins[1]
					topBasins[1] = topBasins[0]
					topBasins[0] = basinSize
				} else if basinSize > topBasins[1] {
					topBasins[2] = topBasins[1]
					topBasins[1] = basinSize
				} else if basinSize > topBasins[2] {
					topBasins[2] = basinSize
				}
			}
		}
	}
	productTopBasins = topBasins[0] * topBasins[1] * topBasins[2]
	output = append(output, fmt.Sprintln("Multiplication of the 3 top Basins is", productTopBasins))

	return output
}
