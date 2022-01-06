package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func advent241(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner
	type key struct {
		z     int
		index int
	}
	knownRes := map[key]bool{}

	var process func(z int, operations [][]string, phase int) int
	process = func(z int, operations [][]string, phase int) int {
		originalZ := z
		for w := 9; w >= 1; w-- {
			x := 0
			y := 0
			z = originalZ
			for _, oper := range operations[phase] {
				par := strings.Split(oper, " ")
				var a *int
				b := 0
				switch par[1] {
				case "w":
					a = &w
				case "x":
					a = &x
				case "y":
					a = &y
				case "z":
					a = &z
				}
				switch par[2] {
				case "w":
					b = w
				case "x":
					b = x
				case "y":
					b = y
				case "z":
					b = z
				default:
					b, err = strconv.Atoi(par[2])
					check(err)
				}
				switch par[0] {
				case "add":
					*a = *a + b
				case "mul":
					*a = *a * b
				case "div":
					*a = *a / b
				case "mod":
					*a = *a % b
				case "eql":
					if *a == b {
						*a = 1
					} else {
						*a = 0
					}
				}
			}

			if phase == len(operations)-1 {
				if z == 0 {
					return w
				}
			} else {
				if !knownRes[key{z: z, index: phase}] {
					res := process(z, operations, phase+1)
					if res != 0 {
						return res*10 + w
					}
					knownRes[key{z: z, index: phase}] = true
				}
			}
		}
		return 0
	}

	fd, err = os.Open(filename)
	check(err)
	scanner = bufio.NewScanner(fd)

	operations := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line[:3] == "inp" {
			operations = append(operations, []string{})
		} else {
			operations[len(operations)-1] = append(operations[len(operations)-1], line)
		}
	}
	fd.Close()

	res := process(0, operations, 0)

	resString := fmt.Sprint(res)
	invertedRes := ""
	for i := len(resString) - 1; i >= 0; i-- {
		invertedRes = invertedRes + string(resString[i])
	}
	fmt.Println("Solution is", invertedRes)
}

func advent242(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner
	type key struct {
		z     int
		index int
	}
	knownRes := map[key]bool{}

	var process func(z int, operations [][]string, phase int) int
	process = func(z int, operations [][]string, phase int) int {
		originalZ := z
		for w := 1; w <= 9; w++ {
			x := 0
			y := 0
			z = originalZ
			for _, oper := range operations[phase] {
				par := strings.Split(oper, " ")
				var a *int
				b := 0
				switch par[1] {
				case "w":
					a = &w
				case "x":
					a = &x
				case "y":
					a = &y
				case "z":
					a = &z
				}
				switch par[2] {
				case "w":
					b = w
				case "x":
					b = x
				case "y":
					b = y
				case "z":
					b = z
				default:
					b, err = strconv.Atoi(par[2])
					check(err)
				}
				switch par[0] {
				case "add":
					*a = *a + b
				case "mul":
					*a = *a * b
				case "div":
					*a = *a / b
				case "mod":
					*a = *a % b
				case "eql":
					if *a == b {
						*a = 1
					} else {
						*a = 0
					}
				}
			}

			if phase == len(operations)-1 {
				if z == 0 {
					return w
				}
			} else {
				if !knownRes[key{z: z, index: phase}] {
					res := process(z, operations, phase+1)
					if res != 0 {
						return res*10 + w
					}
					knownRes[key{z: z, index: phase}] = true
				}
			}
		}
		return 0
	}

	fd, err = os.Open(filename)
	check(err)
	scanner = bufio.NewScanner(fd)

	operations := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line[:3] == "inp" {
			operations = append(operations, []string{})
		} else {
			operations[len(operations)-1] = append(operations[len(operations)-1], line)
		}
	}
	fd.Close()

	res := process(0, operations, 0)

	resString := fmt.Sprint(res)
	invertedRes := ""
	for i := len(resString) - 1; i >= 0; i-- {
		invertedRes = invertedRes + string(resString[i])
	}
	fmt.Println("Solution is", invertedRes)
}
