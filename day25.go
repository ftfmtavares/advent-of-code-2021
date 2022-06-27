package main

import (
	"fmt"
)

func advent251(input []string) []string {
	output := []string{}
	seaCucumbers := [][]string{}
	for _, inputLine := range input {
		seaCucumbers = append(seaCucumbers, []string{})
		for _, c := range inputLine {
			seaCucumbers[len(seaCucumbers)-1] = append(seaCucumbers[len(seaCucumbers)-1], string(c))
		}
	}

	output = append(output, fmt.Sprintln("Initial Positions"))
	for _, line := range seaCucumbers {
		output = append(output, fmt.Sprintln(line))
	}

	stop := false
	stepCount := 0
	for !stop {
		stop = true
		for y := 0; y < len(seaCucumbers); y++ {
			beginX := seaCucumbers[y][0]
			endX := seaCucumbers[y][len(seaCucumbers[y])-1]
			for x := 0; x < len(seaCucumbers[y])-1; x++ {
				if seaCucumbers[y][x] == ">" && seaCucumbers[y][x+1] == "." {
					seaCucumbers[y][x+1] = ">"
					seaCucumbers[y][x] = "."
					x++
					stop = false
				}
			}
			if endX == ">" && beginX == "." {
				seaCucumbers[y][0] = ">"
				seaCucumbers[y][len(seaCucumbers[y])-1] = "."
				stop = false
			}
		}
		for x := 0; x < len(seaCucumbers[0]); x++ {
			beginY := seaCucumbers[0][x]
			endY := seaCucumbers[len(seaCucumbers)-1][x]
			for y := 0; y < len(seaCucumbers)-1; y++ {
				if seaCucumbers[y][x] == "v" && seaCucumbers[y+1][x] == "." {
					seaCucumbers[y+1][x] = "v"
					seaCucumbers[y][x] = "."
					y++
					stop = false
				}
			}
			if endY == "v" && beginY == "." {
				seaCucumbers[0][x] = "v"
				seaCucumbers[len(seaCucumbers)-1][x] = "."
				stop = false
			}
		}
		stepCount++
		output = append(output, fmt.Sprintln("After step", stepCount))
		for _, line := range seaCucumbers {
			output = append(output, fmt.Sprintln(line))
		}
	}

	output = append(output, fmt.Sprintln("The Sea Cucumbers stopped after", stepCount, "steps"))

	return output
}
