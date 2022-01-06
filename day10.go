package main

import (
	"bufio"
	"fmt"
	"os"
)

func advent101(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner
	var sumErrorScore int = 0

	fd, err = os.Open(filename)
	check(err)

	scanner = bufio.NewScanner(fd)

	for scanner.Scan() {
		line := scanner.Text()
		stack := ""
		corrupted := false
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case ')':
				{
					if stack == "" || stack[len(stack)-1] != '(' {
						sumErrorScore = sumErrorScore + 3
						corrupted = true
						fmt.Println("Corrupted Line", line[:i+1])
						break
					} else {
						stack = stack[:len(stack)-1]
					}
				}
			case ']':
				{
					if stack == "" || stack[len(stack)-1] != '[' {
						sumErrorScore = sumErrorScore + 57
						corrupted = true
						fmt.Println("Corrupted Line", line[:i+1])
						break
					} else {
						stack = stack[:len(stack)-1]
					}
				}
			case '}':
				{
					if stack == "" || stack[len(stack)-1] != '{' {
						sumErrorScore = sumErrorScore + 1197
						corrupted = true
						fmt.Println("Corrupted Line", line[:i+1])
						break
					} else {
						stack = stack[:len(stack)-1]
					}
				}
			case '>':
				{
					if stack == "" || stack[len(stack)-1] != '<' {
						sumErrorScore = sumErrorScore + 25137
						corrupted = true
						fmt.Println("Corrupted Line", line[:i+1])
						break
					} else {
						stack = stack[:len(stack)-1]
					}
				}
			case '(':
				{
					stack = stack + "("
				}
			case '[':
				{
					stack = stack + "["
				}
			case '{':
				{
					stack = stack + "{"
				}
			case '<':
				{
					stack = stack + "<"
				}
			}
			if corrupted {
				break
			}
		}
	}

	fd.Close()

	fmt.Println("Syntax Error Score is", sumErrorScore)
}

func advent102(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner
	var scores []int

	fd, err = os.Open(filename)
	check(err)

	scanner = bufio.NewScanner(fd)

	for scanner.Scan() {
		line := scanner.Text()
		stack := ""
		corrupted := false
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case ')':
				{
					if stack == "" || stack[len(stack)-1] != '(' {
						corrupted = true
						fmt.Println("Corrupted Line", line[:i+1])
						break
					} else {
						stack = stack[:len(stack)-1]
					}
				}
			case ']':
				{
					if stack == "" || stack[len(stack)-1] != '[' {
						corrupted = true
						fmt.Println("Corrupted Line", line[:i+1])
						break
					} else {
						stack = stack[:len(stack)-1]
					}
				}
			case '}':
				{
					if stack == "" || stack[len(stack)-1] != '{' {
						corrupted = true
						fmt.Println("Corrupted Line", line[:i+1])
						break
					} else {
						stack = stack[:len(stack)-1]
					}
				}
			case '>':
				{
					if stack == "" || stack[len(stack)-1] != '<' {
						corrupted = true
						fmt.Println("Corrupted Line", line[:i+1])
						break
					} else {
						stack = stack[:len(stack)-1]
					}
				}
			case '(':
				{
					stack = stack + "("
				}
			case '[':
				{
					stack = stack + "["
				}
			case '{':
				{
					stack = stack + "{"
				}
			case '<':
				{
					stack = stack + "<"
				}
			}
			if corrupted {
				break
			}
		}
		if !corrupted {
			score := 0
			for i := len(stack) - 1; i >= 0; i-- {
				score = score * 5
				switch stack[i] {
				case '(':
					{
						score = score + 1
					}
				case '[':
					{
						score = score + 2
					}
				case '{':
					{
						score = score + 3
					}
				case '<':
					{
						score = score + 4
					}
				}
			}
			fmt.Println("Incomplete Line", stack, "with a score of", score)
			if len(scores) == 0 || score > scores[len(scores)-1] {
				scores = append(scores, score)
			} else {
				for i := 0; i < len(scores); i++ {
					if score < scores[i] {
						scores = append(scores[:i+1], scores[i:]...)
						scores[i] = score
						break
					}
				}
			}
		}
	}

	fd.Close()

	fmt.Println("Scores in order", scores)
	middleScore := scores[len(scores)/2]
	fmt.Println("Auto Complete Middle Score is", middleScore)
}
