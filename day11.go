package main

import (
	"fmt"
)

func advent111(input []string) []string {
	output := []string{}
	var totalFlashes int = 0
	const steps int = 100

	lines := [][]int{}
	for _, inputLine := range input {
		line := inputLine
		lines = append(lines, []int{})
		for i := 0; i < len(line); i++ {
			lines[len(lines)-1] = append(lines[len(lines)-1], int(line[i]-'0'))
		}
	}

	for step := 0; step < steps; step++ {
		for i := 0; i < len(lines); i++ {
			for j := 0; j < len(lines[i]); j++ {
				lines[i][j]++
			}
		}

		newFlashes := true
		for newFlashes {
			newFlashes = false
			for i := 0; i < len(lines); i++ {
				for j := 0; j < len(lines[i]); j++ {
					if lines[i][j] == 10 {
						totalFlashes++
						lines[i][j] = 11
						newFlashes = true

						if i > 0 && j > 0 && lines[i-1][j-1] < 10 {
							lines[i-1][j-1]++
						}
						if i > 0 && lines[i-1][j] < 10 {
							lines[i-1][j]++
						}
						if i > 0 && j < len(lines[i])-1 && lines[i-1][j+1] < 10 {
							lines[i-1][j+1]++
						}
						if j > 0 && lines[i][j-1] < 10 {
							lines[i][j-1]++
						}
						if j < len(lines[i])-1 && lines[i][j+1] < 10 {
							lines[i][j+1]++
						}
						if i < len(lines)-1 && j > 0 && lines[i+1][j-1] < 10 {
							lines[i+1][j-1]++
						}
						if i < len(lines)-1 && lines[i+1][j] < 10 {
							lines[i+1][j]++
						}
						if i < len(lines)-1 && j < len(lines[i])-1 && lines[i+1][j+1] < 10 {
							lines[i+1][j+1]++
						}
					}
				}
			}
		}

		for i := 0; i < len(lines); i++ {
			for j := 0; j < len(lines[i]); j++ {
				if lines[i][j] == 11 {
					lines[i][j] = 0
				}
			}
		}
	}

	output = append(output, fmt.Sprintln("Total number of Flashes is", totalFlashes))

	return output
}

func advent112(input []string) []string {
	output := []string{}
	var totalFlashes int = 0
	var step int = 0

	lines := [][]int{}
	for _, inputLine := range input {
		line := inputLine
		lines = append(lines, []int{})
		for i := 0; i < len(line); i++ {
			lines[len(lines)-1] = append(lines[len(lines)-1], int(line[i]-'0'))
		}
	}

	for totalFlashes < 100 {
		totalFlashes = 0
		step++
		for i := 0; i < len(lines); i++ {
			for j := 0; j < len(lines[i]); j++ {
				lines[i][j]++
			}
		}

		newFlashes := true
		for newFlashes {
			newFlashes = false
			for i := 0; i < len(lines); i++ {
				for j := 0; j < len(lines[i]); j++ {
					if lines[i][j] == 10 {
						totalFlashes++
						lines[i][j] = 11
						newFlashes = true

						if i > 0 && j > 0 && lines[i-1][j-1] < 10 {
							lines[i-1][j-1]++
						}
						if i > 0 && lines[i-1][j] < 10 {
							lines[i-1][j]++
						}
						if i > 0 && j < len(lines[i])-1 && lines[i-1][j+1] < 10 {
							lines[i-1][j+1]++
						}
						if j > 0 && lines[i][j-1] < 10 {
							lines[i][j-1]++
						}
						if j < len(lines[i])-1 && lines[i][j+1] < 10 {
							lines[i][j+1]++
						}
						if i < len(lines)-1 && j > 0 && lines[i+1][j-1] < 10 {
							lines[i+1][j-1]++
						}
						if i < len(lines)-1 && lines[i+1][j] < 10 {
							lines[i+1][j]++
						}
						if i < len(lines)-1 && j < len(lines[i])-1 && lines[i+1][j+1] < 10 {
							lines[i+1][j+1]++
						}
					}
				}
			}
		}

		for i := 0; i < len(lines); i++ {
			for j := 0; j < len(lines[i]); j++ {
				if lines[i][j] == 11 {
					lines[i][j] = 0
				}
			}
		}
	}

	output = append(output, fmt.Sprintln("All octuposes flash at the same time on step", step))

	return output
}
