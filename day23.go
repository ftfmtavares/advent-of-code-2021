package main

import (
	"fmt"
	"strings"
)

func advent231(input []string) []string {
	output := []string{}
	var b [15]string

	getEnergy := func(energyType string) int {
		switch energyType {
		case "A":
			return 1
		case "B":
			return 10
		case "C":
			return 100
		case "D":
			return 1000
		default:
			return 0
		}
	}

	validMove := func(b [15]string, start, end int) bool {
		if b[start] == "" || b[end] != "" {
			return false
		}

		if start >= 0 && start <= 6 {
			if end < 7 || end > 14 {
				return false
			} else if b[start] == "A" && end != 7 && end != 8 {
				return false
			} else if b[start] == "B" && end != 9 && end != 10 {
				return false
			} else if b[start] == "C" && end != 11 && end != 12 {
				return false
			} else if b[start] == "D" && end != 13 && end != 14 {
				return false
			} else if end == 7 && b[8] != "A" {
				return false
			} else if end == 9 && b[10] != "B" {
				return false
			} else if end == 11 && b[12] != "C" {
				return false
			} else if end == 13 && b[14] != "D" {
				return false
			} else if end == 8 && b[7] != "" {
				return false
			} else if end == 10 && b[9] != "" {
				return false
			} else if end == 12 && b[11] != "" {
				return false
			} else if end == 14 && b[13] != "" {
				return false
			}
		} else {
			if end < 0 || end > 6 {
				return false
			} else if b[start] == "A" && start == 8 {
				return false
			} else if b[start] == "B" && start == 10 {
				return false
			} else if b[start] == "C" && start == 12 {
				return false
			} else if b[start] == "D" && start == 14 {
				return false
			} else if b[start] == "A" && start == 7 && b[8] == "A" {
				return false
			} else if b[start] == "B" && start == 9 && b[10] == "B" {
				return false
			} else if b[start] == "C" && start == 11 && b[12] == "C" {
				return false
			} else if b[start] == "D" && start == 13 && b[14] == "D" {
				return false
			} else if start == 8 && b[7] != "" {
				return false
			} else if start == 10 && b[9] != "" {
				return false
			} else if start == 12 && b[11] != "" {
				return false
			} else if start == 14 && b[13] != "" {
				return false
			}
		}

		if min(start, end) == 0 {
			if b[1] != "" {
				return false
			}
			if max(start, end) > 8 && b[2] != "" {
				return false
			}
			if max(start, end) > 10 && b[3] != "" {
				return false
			}
			if max(start, end) > 12 && b[4] != "" {
				return false
			}
		} else if min(start, end) == 1 {
			if max(start, end) > 8 && b[2] != "" {
				return false
			}
			if max(start, end) > 10 && b[3] != "" {
				return false
			}
			if max(start, end) > 12 && b[4] != "" {
				return false
			}
		} else if min(start, end) == 2 {
			if max(start, end) > 10 && b[3] != "" {
				return false
			}
			if max(start, end) > 12 && b[4] != "" {
				return false
			}
		} else if min(start, end) == 3 {
			if max(start, end) < 9 && b[2] != "" {
				return false
			}
			if max(start, end) > 12 && b[4] != "" {
				return false
			}
		} else if min(start, end) == 4 {
			if max(start, end) < 11 && b[3] != "" {
				return false
			}
			if max(start, end) < 9 && b[2] != "" {
				return false
			}
		} else if min(start, end) == 5 {
			if max(start, end) < 13 && b[4] != "" {
				return false
			}
			if max(start, end) < 11 && b[3] != "" {
				return false
			}
			if max(start, end) < 9 && b[2] != "" {
				return false
			}
		} else if min(start, end) == 6 {
			if b[5] != "" {
				return false
			}
			if max(start, end) < 13 && b[4] != "" {
				return false
			}
			if max(start, end) < 11 && b[3] != "" {
				return false
			}
			if max(start, end) < 9 && b[2] != "" {
				return false
			}
		}

		return true
	}

	getSteps := func(start, end int) int {
		steps := 0
		if min(start, end) == 0 {
			steps = steps + 2
			if max(start, end) > 8 {
				steps = steps + 2
			}
			if max(start, end) > 10 {
				steps = steps + 2
			}
			if max(start, end) > 12 {
				steps = steps + 2
			}
			steps = steps + 1
			if max(start, end) == 8 || max(start, end) == 10 || max(start, end) == 12 || max(start, end) == 14 {
				steps = steps + 1
			}
		} else if min(start, end) == 1 {
			steps = steps + 1
			if max(start, end) > 8 {
				steps = steps + 2
			}
			if max(start, end) > 10 {
				steps = steps + 2
			}
			if max(start, end) > 12 {
				steps = steps + 2
			}
			steps = steps + 1
			if max(start, end) == 8 || max(start, end) == 10 || max(start, end) == 12 || max(start, end) == 14 {
				steps = steps + 1
			}
		} else if min(start, end) == 2 {
			steps = steps + 1
			if max(start, end) > 10 {
				steps = steps + 2
			}
			if max(start, end) > 12 {
				steps = steps + 2
			}
			steps = steps + 1
			if max(start, end) == 8 || max(start, end) == 10 || max(start, end) == 12 || max(start, end) == 14 {
				steps = steps + 1
			}
		} else if min(start, end) == 3 {
			steps = steps + 1
			if max(start, end) < 9 {
				steps = steps + 2
			}
			if max(start, end) > 12 {
				steps = steps + 2
			}
			steps = steps + 1
			if max(start, end) == 8 || max(start, end) == 10 || max(start, end) == 12 || max(start, end) == 14 {
				steps = steps + 1
			}
		} else if min(start, end) == 4 {
			steps = steps + 1
			if max(start, end) < 11 {
				steps = steps + 2
			}
			if max(start, end) < 9 {
				steps = steps + 2
			}
			steps = steps + 1
			if max(start, end) == 8 || max(start, end) == 10 || max(start, end) == 12 || max(start, end) == 14 {
				steps = steps + 1
			}
		} else if min(start, end) == 5 {
			steps = steps + 1
			if max(start, end) < 13 {
				steps = steps + 2
			}
			if max(start, end) < 11 {
				steps = steps + 2
			}
			if max(start, end) < 9 {
				steps = steps + 2
			}
			steps = steps + 1
			if max(start, end) == 8 || max(start, end) == 10 || max(start, end) == 12 || max(start, end) == 14 {
				steps = steps + 1
			}
		} else if min(start, end) == 6 {
			steps = steps + 2
			if max(start, end) < 13 {
				steps = steps + 2
			}
			if max(start, end) < 11 {
				steps = steps + 2
			}
			if max(start, end) < 9 {
				steps = steps + 2
			}
			steps = steps + 1
			if max(start, end) == 8 || max(start, end) == 10 || max(start, end) == 12 || max(start, end) == 14 {
				steps = steps + 1
			}
		}
		return steps
	}

	var nextMove func(b [15]string, currentEnergy int) int
	nextMove = func(b [15]string, currentEnergy int) int {
		if b[7] == "A" && b[8] == "A" && b[9] == "B" && b[10] == "B" && b[11] == "C" && b[12] == "C" && b[13] == "D" && b[14] == "D" {
			return currentEnergy
		}

		for start := 0; start <= 6; start++ {
			for end := 7; end <= 14; end++ {
				if validMove(b, start, end) {
					b[end] = b[start]
					b[start] = ""
					newEnergy := currentEnergy + getSteps(start, end)*getEnergy(b[end])
					newEnergy = nextMove(b, newEnergy)
					b[start] = b[end]
					b[end] = ""
					return newEnergy
				}
			}
		}

		bestMove := -1
		for start := 7; start <= 14; start++ {
			for end := 0; end <= 6; end++ {
				if validMove(b, start, end) {
					b[end] = b[start]
					b[start] = ""
					newEnergy := currentEnergy + getSteps(start, end)*getEnergy(b[end])
					newEnergy = nextMove(b, newEnergy)
					if newEnergy != -1 && (newEnergy < bestMove || bestMove == -1) {
						bestMove = newEnergy
					}
					b[start] = b[end]
					b[end] = ""
				}
			}
		}

		return bestMove
	}

	b = [15]string{}
	for i := 0; i < 7; i++ {
		b[i] = ""
	}

	line := input[2]
	line = strings.Replace(line, "#", "", -1)
	b[7] = line[:1]
	b[9] = line[1:2]
	b[11] = line[2:3]
	b[13] = line[3:]

	line = input[3]
	line = strings.Replace(line, "#", "", -1)
	line = strings.Replace(line, " ", "", -1)
	b[8] = line[:1]
	b[10] = line[1:2]
	b[12] = line[2:3]
	b[14] = line[3:]

	output = append(output, fmt.Sprintln("Initial Positions are", b))

	leastEnergy := nextMove(b, 0)
	output = append(output, fmt.Sprintln("Best solution uses", leastEnergy, "energy"))

	return output
}

func advent232(input []string) []string {
	output := []string{}
	var b [23]string

	getEnergy := func(energyType string) int {
		switch energyType {
		case "A":
			return 1
		case "B":
			return 10
		case "C":
			return 100
		case "D":
			return 1000
		default:
			return 0
		}
	}

	validMove := func(b [23]string, start, end int) bool {
		if b[start] == "" || b[end] != "" {
			return false
		}

		if start >= 0 && start <= 6 {
			if end < 7 || end > 22 {
				return false
			} else if b[start] == "A" && (end < 7 || end > 10) {
				return false
			} else if b[start] == "B" && (end < 11 || end > 14) {
				return false
			} else if b[start] == "C" && (end < 15 || end > 18) {
				return false
			} else if b[start] == "D" && (end < 19 || end > 22) {
				return false
			} else if end == 7 && (b[8] != "A" || b[9] != "A" || b[10] != "A") {
				return false
			} else if end == 8 && (b[9] != "A" || b[10] != "A") {
				return false
			} else if end == 9 && b[10] != "A" {
				return false
			} else if end == 11 && (b[12] != "B" || b[13] != "B" || b[14] != "B") {
				return false
			} else if end == 12 && (b[13] != "B" || b[14] != "B") {
				return false
			} else if end == 13 && b[14] != "B" {
				return false
			} else if end == 15 && (b[16] != "C" || b[17] != "C" || b[18] != "C") {
				return false
			} else if end == 16 && (b[17] != "C" || b[18] != "C") {
				return false
			} else if end == 17 && b[18] != "C" {
				return false
			} else if end == 19 && (b[20] != "D" || b[21] != "D" || b[22] != "D") {
				return false
			} else if end == 20 && (b[21] != "D" || b[22] != "D") {
				return false
			} else if end == 21 && b[22] != "D" {
				return false
			}
		} else {
			if end < 0 || end > 6 {
				return false
			} else if b[start] == "A" && start == 10 {
				return false
			} else if b[start] == "B" && start == 14 {
				return false
			} else if b[start] == "C" && start == 18 {
				return false
			} else if b[start] == "D" && start == 22 {
				return false
			} else if b[start] == "A" && start == 9 && b[10] == "A" {
				return false
			} else if b[start] == "B" && start == 13 && b[14] == "B" {
				return false
			} else if b[start] == "C" && start == 17 && b[18] == "C" {
				return false
			} else if b[start] == "D" && start == 21 && b[22] == "D" {
				return false
			} else if b[start] == "A" && start == 8 && b[10] == "A" && b[9] == "A" {
				return false
			} else if b[start] == "B" && start == 12 && b[14] == "B" && b[13] == "B" {
				return false
			} else if b[start] == "C" && start == 16 && b[18] == "C" && b[17] == "C" {
				return false
			} else if b[start] == "D" && start == 20 && b[22] == "D" && b[21] == "D" {
				return false
			} else if b[start] == "A" && start == 7 && b[10] == "A" && b[9] == "A" && b[8] == "A" {
				return false
			} else if b[start] == "B" && start == 11 && b[14] == "B" && b[13] == "B" && b[12] == "B" {
				return false
			} else if b[start] == "C" && start == 15 && b[18] == "C" && b[17] == "C" && b[16] == "C" {
				return false
			} else if b[start] == "D" && start == 19 && b[22] == "D" && b[21] == "D" && b[20] == "D" {
				return false
			} else if start == 8 && b[7] != "" {
				return false
			} else if start == 9 && b[8] != "" {
				return false
			} else if start == 10 && b[9] != "" {
				return false
			} else if start == 12 && b[11] != "" {
				return false
			} else if start == 13 && b[12] != "" {
				return false
			} else if start == 14 && b[13] != "" {
				return false
			} else if start == 16 && b[15] != "" {
				return false
			} else if start == 17 && b[16] != "" {
				return false
			} else if start == 18 && b[17] != "" {
				return false
			} else if start == 20 && b[19] != "" {
				return false
			} else if start == 21 && b[20] != "" {
				return false
			} else if start == 22 && b[21] != "" {
				return false
			}
		}

		if min(start, end) == 0 {
			if b[1] != "" {
				return false
			}
			if max(start, end) > 10 && b[2] != "" {
				return false
			}
			if max(start, end) > 14 && b[3] != "" {
				return false
			}
			if max(start, end) > 18 && b[4] != "" {
				return false
			}
		} else if min(start, end) == 1 {
			if max(start, end) > 10 && b[2] != "" {
				return false
			}
			if max(start, end) > 14 && b[3] != "" {
				return false
			}
			if max(start, end) > 18 && b[4] != "" {
				return false
			}
		} else if min(start, end) == 2 {
			if max(start, end) > 14 && b[3] != "" {
				return false
			}
			if max(start, end) > 18 && b[4] != "" {
				return false
			}
		} else if min(start, end) == 3 {
			if max(start, end) < 11 && b[2] != "" {
				return false
			}
			if max(start, end) > 18 && b[4] != "" {
				return false
			}
		} else if min(start, end) == 4 {
			if max(start, end) < 15 && b[3] != "" {
				return false
			}
			if max(start, end) < 11 && b[2] != "" {
				return false
			}
		} else if min(start, end) == 5 {
			if max(start, end) < 19 && b[4] != "" {
				return false
			}
			if max(start, end) < 15 && b[3] != "" {
				return false
			}
			if max(start, end) < 11 && b[2] != "" {
				return false
			}
		} else if min(start, end) == 6 {
			if b[5] != "" {
				return false
			}
			if max(start, end) < 19 && b[4] != "" {
				return false
			}
			if max(start, end) < 15 && b[3] != "" {
				return false
			}
			if max(start, end) < 11 && b[2] != "" {
				return false
			}
		}

		return true
	}

	getSteps := func(start, end int) int {
		steps := 0
		if min(start, end) == 0 {
			steps = steps + 2
			if max(start, end) > 10 {
				steps = steps + 2
			}
			if max(start, end) > 14 {
				steps = steps + 2
			}
			if max(start, end) > 18 {
				steps = steps + 2
			}
		} else if min(start, end) == 1 {
			steps = steps + 1
			if max(start, end) > 10 {
				steps = steps + 2
			}
			if max(start, end) > 14 {
				steps = steps + 2
			}
			if max(start, end) > 18 {
				steps = steps + 2
			}
		} else if min(start, end) == 2 {
			steps = steps + 1
			if max(start, end) > 14 {
				steps = steps + 2
			}
			if max(start, end) > 18 {
				steps = steps + 2
			}
		} else if min(start, end) == 3 {
			steps = steps + 1
			if max(start, end) < 11 {
				steps = steps + 2
			}
			if max(start, end) > 18 {
				steps = steps + 2
			}
		} else if min(start, end) == 4 {
			steps = steps + 1
			if max(start, end) < 15 {
				steps = steps + 2
			}
			if max(start, end) < 11 {
				steps = steps + 2
			}
		} else if min(start, end) == 5 {
			steps = steps + 1
			if max(start, end) < 19 {
				steps = steps + 2
			}
			if max(start, end) < 15 {
				steps = steps + 2
			}
			if max(start, end) < 11 {
				steps = steps + 2
			}
		} else if min(start, end) == 6 {
			steps = steps + 2
			if max(start, end) < 19 {
				steps = steps + 2
			}
			if max(start, end) < 15 {
				steps = steps + 2
			}
			if max(start, end) < 11 {
				steps = steps + 2
			}
		}
		steps = steps + 1
		if max(start, end) == 8 || max(start, end) == 12 || max(start, end) == 16 || max(start, end) == 20 {
			steps = steps + 1
		} else if max(start, end) == 9 || max(start, end) == 13 || max(start, end) == 17 || max(start, end) == 21 {
			steps = steps + 2
		} else if max(start, end) == 10 || max(start, end) == 14 || max(start, end) == 18 || max(start, end) == 22 {
			steps = steps + 3
		}
		return steps
	}

	var nextMove func(b [23]string, currentEnergy int, level int) int
	nextMove = func(b [23]string, currentEnergy int, level int) int {
		if b[7] == "A" && b[8] == "A" && b[9] == "A" && b[10] == "A" && b[11] == "B" && b[12] == "B" && b[13] == "B" && b[14] == "B" && b[15] == "C" && b[16] == "C" && b[17] == "C" && b[18] == "C" && b[19] == "D" && b[20] == "D" && b[21] == "D" && b[22] == "D" {
			return currentEnergy
		}

		for start := 0; start <= 6; start++ {
			for end := 7; end <= 22; end++ {
				if validMove(b, start, end) {
					b[end] = b[start]
					b[start] = ""
					newEnergy := currentEnergy + getSteps(start, end)*getEnergy(b[end])
					newEnergy = nextMove(b, newEnergy, level+1)
					b[start] = b[end]
					b[end] = ""
					return newEnergy
				}
			}
		}

		bestMove := -1
		for start := 7; start <= 22; start++ {
			for end := 0; end <= 6; end++ {
				if validMove(b, start, end) {
					b[end] = b[start]
					b[start] = ""
					newEnergy := currentEnergy + getSteps(start, end)*getEnergy(b[end])
					newEnergy = nextMove(b, newEnergy, level+1)
					if newEnergy != -1 && (newEnergy < bestMove || bestMove == -1) {
						bestMove = newEnergy
					}
					b[start] = b[end]
					b[end] = ""
				}
			}
		}

		return bestMove
	}

	b = [23]string{}
	for i := 0; i < 7; i++ {
		b[i] = ""
	}

	line := input[2]
	line = strings.Replace(line, "#", "", -1)
	b[7] = line[:1]
	b[11] = line[1:2]
	b[15] = line[2:3]
	b[19] = line[3:]

	line = "  #D#C#B#A#"
	line = strings.Replace(line, "#", "", -1)
	line = strings.Replace(line, " ", "", -1)
	b[8] = line[:1]
	b[12] = line[1:2]
	b[16] = line[2:3]
	b[20] = line[3:]

	line = "  #D#B#A#C#"
	line = strings.Replace(line, "#", "", -1)
	line = strings.Replace(line, " ", "", -1)
	b[9] = line[:1]
	b[13] = line[1:2]
	b[17] = line[2:3]
	b[21] = line[3:]

	line = input[3]
	line = strings.Replace(line, "#", "", -1)
	line = strings.Replace(line, " ", "", -1)
	b[10] = line[:1]
	b[14] = line[1:2]
	b[18] = line[2:3]
	b[22] = line[3:]

	output = append(output, fmt.Sprintln("Initial Positions are", b))

	leastEnergy := nextMove(b, 0, 1)
	output = append(output, fmt.Sprintln("Best solution uses", leastEnergy, "energy"))

	return output
}
