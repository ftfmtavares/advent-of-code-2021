package main

import (
	"fmt"
)

func advent101(input []string) []string {
	output := []string{}
	var sumErrorScore int = 0

	for _, inputLine := range input {
		line := inputLine
		stack := ""
		corrupted := false
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case ')':
				{
					if stack == "" || stack[len(stack)-1] != '(' {
						sumErrorScore = sumErrorScore + 3
						corrupted = true
						output = append(output, fmt.Sprintln("Corrupted Line", line[:i+1]))
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
						output = append(output, fmt.Sprintln("Corrupted Line", line[:i+1]))
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
						output = append(output, fmt.Sprintln("Corrupted Line", line[:i+1]))
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
						output = append(output, fmt.Sprintln("Corrupted Line", line[:i+1]))
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

	output = append(output, fmt.Sprintln("Syntax Error Score is", sumErrorScore))

	return output
}

func advent102(input []string) []string {
	output := []string{}
	var scores []int

	for _, inputLine := range input {
		line := inputLine
		stack := ""
		corrupted := false
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case ')':
				{
					if stack == "" || stack[len(stack)-1] != '(' {
						corrupted = true
						output = append(output, fmt.Sprintln("Corrupted Line", line[:i+1]))
						break
					} else {
						stack = stack[:len(stack)-1]
					}
				}
			case ']':
				{
					if stack == "" || stack[len(stack)-1] != '[' {
						corrupted = true
						output = append(output, fmt.Sprintln("Corrupted Line", line[:i+1]))
						break
					} else {
						stack = stack[:len(stack)-1]
					}
				}
			case '}':
				{
					if stack == "" || stack[len(stack)-1] != '{' {
						corrupted = true
						output = append(output, fmt.Sprintln("Corrupted Line", line[:i+1]))
						break
					} else {
						stack = stack[:len(stack)-1]
					}
				}
			case '>':
				{
					if stack == "" || stack[len(stack)-1] != '<' {
						corrupted = true
						output = append(output, fmt.Sprintln("Corrupted Line", line[:i+1]))
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
			output = append(output, fmt.Sprintln("Incomplete Line", stack, "with a score of", score))
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

	output = append(output, fmt.Sprintln("Scores in order", scores))
	middleScore := scores[len(scores)/2]
	output = append(output, fmt.Sprintln("Auto Complete Middle Score is", middleScore))

	return output
}
