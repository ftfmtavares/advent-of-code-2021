package main

import (
	"fmt"
	"strconv"
	"strings"
)

func advent131(input []string) []string {
	var err error
	output := []string{}
	dots := [][]rune{}
	folds := []string{}
	for _, inputLine := range input {
		if strings.HasPrefix(inputLine, "fold along ") {
			folds = append(folds, strings.Replace(inputLine, "fold along ", "", -1))
		} else if inputLine != "" {
			newDot := strings.Split(inputLine, ",")
			x, err := strconv.Atoi(newDot[0])
			check(err)
			y, err := strconv.Atoi(newDot[1])
			check(err)
			for y >= len(dots) {
				dots = append(dots, []rune{})
			}
			for i := 0; i < len(dots); i++ {
				for x >= len(dots[i]) {
					dots[i] = append(dots[i], '.')
				}
			}
			dots[y][x] = '#'
		}
	}

	for i := 0; i < len(dots); i++ {
		output = append(output, fmt.Sprintln(string(dots[i])))
	}

	output = append(output, fmt.Sprintln(folds[0]))
	newFold := strings.Split(folds[0], "=")
	pos, err := strconv.Atoi(newFold[1])
	check(err)
	if newFold[0] == "y" {
		for j := pos + 1; j < len(dots); j++ {
			for k := 0; k < len(dots[j]); k++ {
				if dots[j][k] == '#' {
					dots[2*pos-j][k] = '#'
				}
			}
		}
		dots = dots[:pos]
	} else {
		for j := 0; j < len(dots); j++ {
			for k := pos + 1; k < len(dots[j]); k++ {
				if dots[j][k] == '#' {
					dots[j][2*pos-k] = '#'
				}
			}
			dots[j] = dots[j][:pos]
		}
	}
	for j := 0; j < len(dots); j++ {
		output = append(output, fmt.Sprintln(string(dots[j])))
	}

	visibleDots := 0
	for i := 0; i < len(dots); i++ {
		for j := 0; j < len(dots[i]); j++ {
			if dots[i][j] == '#' {
				visibleDots++
			}
		}
	}
	output = append(output, fmt.Sprintln("Visible dots are", visibleDots))

	return output
}

func advent132(input []string) []string {
	output := []string{}
	dots := [][]rune{}
	folds := []string{}
	for _, inputLine := range input {
		if strings.HasPrefix(inputLine, "fold along ") {
			folds = append(folds, strings.Replace(inputLine, "fold along ", "", -1))
		} else if inputLine != "" {
			newDot := strings.Split(inputLine, ",")
			x, err := strconv.Atoi(newDot[0])
			check(err)
			y, err := strconv.Atoi(newDot[1])
			check(err)
			for y >= len(dots) {
				dots = append(dots, []rune{})
			}
			for i := 0; i < len(dots); i++ {
				for x >= len(dots[i]) {
					dots[i] = append(dots[i], '.')
				}
			}
			dots[y][x] = '#'
		}
	}

	for i := 0; i < len(dots); i++ {
		output = append(output, fmt.Sprintln(string(dots[i])))
	}

	for i := 0; i < len(folds); i++ {
		output = append(output, fmt.Sprintln(folds[i]))
		newFold := strings.Split(folds[i], "=")
		pos, err := strconv.Atoi(newFold[1])
		check(err)
		if newFold[0] == "y" {
			for j := pos + 1; j < len(dots); j++ {
				for k := 0; k < len(dots[j]); k++ {
					if dots[j][k] == '#' {
						dots[2*pos-j][k] = '#'
					}
				}
			}
			dots = dots[:pos]
		} else {
			for j := 0; j < len(dots); j++ {
				for k := pos + 1; k < len(dots[j]); k++ {
					if dots[j][k] == '#' {
						dots[j][2*pos-k] = '#'
					}
				}
				dots[j] = dots[j][:pos]
			}
		}
		for j := 0; j < len(dots); j++ {
			output = append(output, fmt.Sprintln(string(dots[j])))
		}
	}

	visibleDots := 0
	for i := 0; i < len(dots); i++ {
		for j := 0; j < len(dots[i]); j++ {
			if dots[i][j] == '#' {
				visibleDots++
			}
		}
	}
	output = append(output, fmt.Sprintln("Visible dots are", visibleDots))

	return output
}
