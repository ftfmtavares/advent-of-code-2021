package main

import (
	"fmt"
	"strconv"
	"strings"
)

func advent021(input []string) []string {
	var err error
	output := []string{}
	var commands []string
	var pos int = 0
	var depth int = 0
	var readValue int

	for _, inputLine := range input {
		commands = strings.Fields(inputLine)
		if len(commands) == 2 {
			switch commands[0] {
			case "forward":
				readValue, err = strconv.Atoi(commands[1])
				check(err)
				pos = pos + readValue
				output = append(output, fmt.Sprintf("Submarine moved forward %d steps to %d position\n", readValue, pos))
			case "down":
				readValue, err = strconv.Atoi(commands[1])
				check(err)
				depth = depth + readValue
				output = append(output, fmt.Sprintf("Submarine lowered %d steps to %d depth\n", readValue, depth))
			case "up":
				readValue, err = strconv.Atoi(commands[1])
				check(err)
				depth = depth - readValue
				if depth < 0 {
					output = append(output, fmt.Sprintln("WARNING - SUBMARINE EMERGED"))
				}
				output = append(output, fmt.Sprintf("Submarine climbed %d steps to %d depth\n", readValue, depth))
			}
		}
	}

	output = append(output, fmt.Sprintf("Submarine final pos=%d and depth=%d giving an result of %d\n", pos, depth, pos*depth))

	return output
}

func advent022(input []string) []string {
	var err error
	output := []string{}
	var commands []string
	var pos int = 0
	var depth int = 0
	var aim int = 0
	var readValue int

	for _, inputLine := range input {
		commands = strings.Fields(inputLine)
		if len(commands) == 2 {
			switch commands[0] {
			case "forward":
				readValue, err = strconv.Atoi(commands[1])
				check(err)
				pos = pos + readValue
				depth = depth + aim*readValue
				if depth < 0 {
					output = append(output, fmt.Sprintln("WARNING - SUBMARINE EMERGED"))
				}
				if aim > 0 {
					output = append(output, fmt.Sprintf("Submarine moved forward %d steps to %d position and lowered %d steps to %d depth\n", readValue, pos, aim*readValue, depth))
				} else if aim < 0 {
					output = append(output, fmt.Sprintf("Submarine moved forward %d steps to %d position and climbed %d steps to %d depth\n", readValue, pos, -aim*readValue, depth))
				} else {
					output = append(output, fmt.Sprintf("Submarine moved forward %d steps to %d position\n", readValue, pos))
				}

			case "down":
				readValue, err = strconv.Atoi(commands[1])
				check(err)
				aim = aim + readValue
				output = append(output, fmt.Sprintf("Submarine aimed %d degrees down to %d\n", readValue, aim))
			case "up":
				readValue, err = strconv.Atoi(commands[1])
				check(err)
				aim = aim - readValue
				output = append(output, fmt.Sprintf("Submarine aimed %d degrees up to %d\n", readValue, aim))
			}
		}
	}

	output = append(output, fmt.Sprintf("Submarine final pos=%d and depth=%d giving an result of %d\n", pos, depth, pos*depth))

	return output
}
