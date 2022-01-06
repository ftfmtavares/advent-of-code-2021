package main

import (
	"bufio"
	"fmt"
	"os"
)

func advent091(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner
	var riskSum int = 0

	fd, err = os.Open(filename)
	check(err)

	scanner = bufio.NewScanner(fd)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fd.Close()

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

	fmt.Println("Sum of all risk levels is", riskSum)
}

func advent092(filename string) {
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

	var fd *os.File
	var err error
	var scanner *bufio.Scanner
	var productTopBasins int = 0

	fd, err = os.Open(filename)
	check(err)

	scanner = bufio.NewScanner(fd)

	lines := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, []int{})
		for i := 0; i < len(line); i++ {
			lines[len(lines)-1] = append(lines[len(lines)-1], int(line[i]-'0'))
		}
	}

	fd.Close()

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
	fmt.Println("Multiplication of the 3 top Basins is", productTopBasins)
}
