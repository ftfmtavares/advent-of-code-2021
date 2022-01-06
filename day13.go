package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func advent131(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner

	fd, err = os.Open(filename)
	check(err)
	scanner = bufio.NewScanner(fd)

	dots := [][]rune{}
	folds := []string{}
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "fold along ") {
			folds = append(folds, strings.Replace(scanner.Text(), "fold along ", "", -1))
		} else if scanner.Text() != "" {
			newDot := strings.Split(scanner.Text(), ",")
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
	fd.Close()

	for i := 0; i < len(dots); i++ {
		fmt.Println(string(dots[i]))
	}

	fmt.Println(folds[0])
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
		dots = append(dots[:pos])
	} else {
		for j := 0; j < len(dots); j++ {
			for k := pos + 1; k < len(dots[j]); k++ {
				if dots[j][k] == '#' {
					dots[j][2*pos-k] = '#'
				}
			}
			dots[j] = append(dots[j][:pos])
		}
	}
	for j := 0; j < len(dots); j++ {
		fmt.Println(string(dots[j]))
	}

	visibleDots := 0
	for i := 0; i < len(dots); i++ {
		for j := 0; j < len(dots[i]); j++ {
			if dots[i][j] == '#' {
				visibleDots++
			}
		}
	}
	fmt.Println("Visible dots are", visibleDots)
}

func advent132(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner

	fd, err = os.Open(filename)
	check(err)
	scanner = bufio.NewScanner(fd)

	dots := [][]rune{}
	folds := []string{}
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "fold along ") {
			folds = append(folds, strings.Replace(scanner.Text(), "fold along ", "", -1))
		} else if scanner.Text() != "" {
			newDot := strings.Split(scanner.Text(), ",")
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
	fd.Close()

	for i := 0; i < len(dots); i++ {
		fmt.Println(string(dots[i]))
	}

	for i := 0; i < len(folds); i++ {
		fmt.Println(folds[i])
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
			dots = append(dots[:pos])
		} else {
			for j := 0; j < len(dots); j++ {
				for k := pos + 1; k < len(dots[j]); k++ {
					if dots[j][k] == '#' {
						dots[j][2*pos-k] = '#'
					}
				}
				dots[j] = append(dots[j][:pos])
			}
		}
		for j := 0; j < len(dots); j++ {
			fmt.Println(string(dots[j]))
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
	fmt.Println("Visible dots are", visibleDots)
}
