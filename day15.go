package main

import (
	"fmt"
)

func advent151(input []string) []string {
	output := []string{}
	cavernRisk := [][]int{}
	cavernBestPath := [][]int{}
	bestPath := -1

	var getBestPath func(x, y, path int) int
	getBestPath = func(x, y, path int) int {
		path = path + cavernRisk[y][x]
		if path > bestPath && bestPath != -1 {
			return -1
		}
		if path < cavernBestPath[y][x] || cavernBestPath[y][x] == -1 {
			cavernBestPath[y][x] = path
		}
		if y == len(cavernRisk)-1 && x == len(cavernRisk[len(cavernRisk)-1])-1 {
			if path < bestPath || bestPath == -1 {
				bestPath = path
				output = append(output, fmt.Sprintln("A new best path has been found with risk Level of", bestPath))
			}
			return path
		}

		downPath := -1
		rightPath := -1
		upPath := -1
		leftPath := -1
		if y < len(cavernRisk)-1 && (path+cavernRisk[y+1][x] < cavernBestPath[y+1][x] || cavernBestPath[y+1][x] == -1) {
			downPath = getBestPath(x, y+1, path)
		}
		if x < len(cavernRisk[y])-1 && (path+cavernRisk[y][x+1] < cavernBestPath[y][x+1] || cavernBestPath[y][x+1] == -1) {
			rightPath = getBestPath(x+1, y, path)
		}
		if y > 1 && (path+cavernRisk[y-1][x] < cavernBestPath[y-1][x] || cavernBestPath[y-1][x] == -1) {
			upPath = getBestPath(x, y-1, path)
		}
		if x > 1 && (path+cavernRisk[y][x-1] < cavernBestPath[y][x-1] || cavernBestPath[y][x-1] == -1) {
			leftPath = getBestPath(x-1, y, path)
		}
		res := downPath
		if res < rightPath && rightPath != -1 {
			res = rightPath
		}
		if res < upPath && upPath != -1 {
			res = upPath
		}
		if res < leftPath && leftPath != -1 {
			res = leftPath
		}

		return res
	}

	for _, inputLine := range input {
		cavernRisk = append(cavernRisk, []int{})
		cavernBestPath = append(cavernBestPath, []int{})
		line := inputLine
		for _, v := range line {
			cavernRisk[len(cavernRisk)-1] = append(cavernRisk[len(cavernRisk)-1], int(v-'0'))
			cavernBestPath[len(cavernBestPath)-1] = append(cavernBestPath[len(cavernBestPath)-1], -1)
		}
	}

	cavernBestPath[0][0] = 0
	getBestPath(1, 0, 0)
	getBestPath(0, 1, 0)

	output = append(output, fmt.Sprintln("Best Path Risk Level is", bestPath))

	return output
}

func advent152(input []string) []string {
	output := []string{}
	cavernRisk := [][]int{}
	cavernBestPath := [][]int{}
	bestPath := -1

	var getBestPath func(x, y, path int) int
	getBestPath = func(x, y, path int) int {
		path = path + cavernRisk[y][x]
		if path > bestPath && bestPath != -1 {
			return -1
		}
		if path < cavernBestPath[y][x] || cavernBestPath[y][x] == -1 {
			cavernBestPath[y][x] = path
		}
		if y == len(cavernRisk)-1 && x == len(cavernRisk[len(cavernRisk)-1])-1 {
			if path < bestPath || bestPath == -1 {
				bestPath = path
				output = append(output, fmt.Sprintln("A new best path has been found with risk Level of", bestPath))
			}
			return path
		}

		downPath := -1
		rightPath := -1
		upPath := -1
		leftPath := -1
		if y < len(cavernRisk)-1 && (path+cavernRisk[y+1][x] < cavernBestPath[y+1][x] || cavernBestPath[y+1][x] == -1) {
			downPath = getBestPath(x, y+1, path)
		}
		if x < len(cavernRisk[y])-1 && (path+cavernRisk[y][x+1] < cavernBestPath[y][x+1] || cavernBestPath[y][x+1] == -1) {
			rightPath = getBestPath(x+1, y, path)
		}
		if y > 1 && (path+cavernRisk[y-1][x] < cavernBestPath[y-1][x] || cavernBestPath[y-1][x] == -1) {
			upPath = getBestPath(x, y-1, path)
		}
		if x > 1 && (path+cavernRisk[y][x-1] < cavernBestPath[y][x-1] || cavernBestPath[y][x-1] == -1) {
			leftPath = getBestPath(x-1, y, path)
		}
		res := downPath
		if res < rightPath && rightPath != -1 {
			res = rightPath
		}
		if res < upPath && upPath != -1 {
			res = upPath
		}
		if res < leftPath && leftPath != -1 {
			res = leftPath
		}

		return res
	}

	for _, inputLine := range input {
		cavernRisk = append(cavernRisk, []int{})
		cavernBestPath = append(cavernBestPath, []int{})
		line := inputLine
		for _, v := range line {
			cavernRisk[len(cavernRisk)-1] = append(cavernRisk[len(cavernRisk)-1], int(v-'0'))
			cavernBestPath[len(cavernBestPath)-1] = append(cavernBestPath[len(cavernBestPath)-1], -1)
		}
	}

	lenY := len(cavernRisk)
	lenX := len(cavernRisk[0])
	for i := 0; i < 4; i++ {
		for y := 0; y < lenY; y++ {
			for x := 0 + i*lenX; x < (i+1)*lenX; x++ {
				cavernBestPath[y] = append(cavernBestPath[y], -1)
				if cavernRisk[y][x] == 9 {
					cavernRisk[y] = append(cavernRisk[y], 1)
				} else {
					cavernRisk[y] = append(cavernRisk[y], cavernRisk[y][x]+1)
				}
			}
		}
	}
	for i := 0; i < 4; i++ {
		for y := 0 + i*lenY; y < (i+1)*lenY; y++ {
			cavernRisk = append(cavernRisk, []int{})
			cavernBestPath = append(cavernBestPath, []int{})
			for x := 0; x < len(cavernRisk[y]); x++ {
				cavernBestPath[len(cavernRisk)-1] = append(cavernBestPath[len(cavernRisk)-1], -1)
				if cavernRisk[y][x] == 9 {
					cavernRisk[len(cavernRisk)-1] = append(cavernRisk[len(cavernRisk)-1], 1)
				} else {
					cavernRisk[len(cavernRisk)-1] = append(cavernRisk[len(cavernRisk)-1], cavernRisk[y][x]+1)
				}
			}
		}
	}

	cavernBestPath[0][0] = 0
	getBestPath(1, 0, 0)
	getBestPath(0, 1, 0)

	output = append(output, fmt.Sprintln("Best Path Risk Level is", bestPath))

	return output
}
