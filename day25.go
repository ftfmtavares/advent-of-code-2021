package main

import (
	"bufio"
	"fmt"
	"os"
)

func advent251(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner

	fd, err = os.Open(filename)
	check(err)
	scanner = bufio.NewScanner(fd)

	seaCucumbers := [][]string{}
	for scanner.Scan() {
		seaCucumbers = append(seaCucumbers, []string{})
		for _, c := range scanner.Text() {
			seaCucumbers[len(seaCucumbers)-1] = append(seaCucumbers[len(seaCucumbers)-1], string(c))
		}
	}
	fd.Close()

	fmt.Println("Initial Positions")
	for _, line := range seaCucumbers {
		fmt.Println(line)
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
		fmt.Println("After step", stepCount)
		for _, line := range seaCucumbers {
			fmt.Println(line)
		}
	}

	fmt.Println("The Sea Cucumbers stopped after", stepCount, "steps")
}
