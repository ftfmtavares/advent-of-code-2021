package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func advent021(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner
	var commands []string
	var pos int = 0
	var depth int = 0
	var readValue int

	fd, err = os.Open(filename)
	check(err)

	scanner = bufio.NewScanner(fd)

	for scanner.Scan() {
		commands = strings.Fields(scanner.Text())
		if len(commands) == 2 {
			switch commands[0] {
			case "forward":
				readValue, err = strconv.Atoi(commands[1])
				check(err)
				pos = pos + readValue
				fmt.Printf("Submarine moved forward %d steps to %d position\n", readValue, pos)
			case "down":
				readValue, err = strconv.Atoi(commands[1])
				check(err)
				depth = depth + readValue
				fmt.Printf("Submarine lowered %d steps to %d depth\n", readValue, depth)
			case "up":
				readValue, err = strconv.Atoi(commands[1])
				check(err)
				depth = depth - readValue
				if depth < 0 {
					fmt.Println("WARNING - SUBMARINE EMERGED")
				}
				fmt.Printf("Submarine climbed %d steps to %d depth\n", readValue, depth)
			}
		}
	}

	fmt.Printf("Submarine final pos=%d and depth=%d giving an result of %d\n", pos, depth, pos*depth)

	fd.Close()
}

func advent022(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner
	var commands []string
	var pos int = 0
	var depth int = 0
	var aim int = 0
	var readValue int

	fd, err = os.Open(filename)
	check(err)

	scanner = bufio.NewScanner(fd)

	for scanner.Scan() {
		commands = strings.Fields(scanner.Text())
		if len(commands) == 2 {
			switch commands[0] {
			case "forward":
				readValue, err = strconv.Atoi(commands[1])
				check(err)
				pos = pos + readValue
				depth = depth + aim*readValue
				if depth < 0 {
					fmt.Println("WARNING - SUBMARINE EMERGED")
				}
				if aim > 0 {
					fmt.Printf("Submarine moved forward %d steps to %d position and lowered %d steps to %d depth\n", readValue, pos, aim*readValue, depth)
				} else if aim < 0 {
					fmt.Printf("Submarine moved forward %d steps to %d position and climbed %d steps to %d depth\n", readValue, pos, -aim*readValue, depth)
				} else {
					fmt.Printf("Submarine moved forward %d steps to %d position\n", readValue, pos)
				}

			case "down":
				readValue, err = strconv.Atoi(commands[1])
				check(err)
				aim = aim + readValue
				fmt.Printf("Submarine aimed %d degrees down to %d\n", readValue, aim)
			case "up":
				readValue, err = strconv.Atoi(commands[1])
				check(err)
				aim = aim - readValue
				fmt.Printf("Submarine aimed %d degrees up to %d\n", readValue, aim)
			}
		}
	}

	fmt.Printf("Submarine final pos=%d and depth=%d giving an result of %d\n", pos, depth, pos*depth)

	fd.Close()
}
