package main

import (
	"fmt"
	"strings"
)

func advent141(input []string) []string {
	output := []string{}
	const steps = 10

	polymer := ""
	insertionPairs := [][]string{}
	for _, inputLine := range input {
		if strings.Contains(inputLine, " -> ") {
			insertionPairs = append(insertionPairs, []string{})
			insertionPairs[len(insertionPairs)-1] = append(insertionPairs[len(insertionPairs)-1], strings.Split(inputLine, " -> ")...)
		} else if inputLine != "" {
			polymer = inputLine
		}
	}

	output = append(output, fmt.Sprintln(polymer))
	for i := 0; i < steps; i++ {
		newPolymer := ""
		for j := 0; j < len(polymer)-1; j++ {
			newPolymer = newPolymer + string(polymer[j])
			for _, pair := range insertionPairs {
				if polymer[j:j+2] == pair[0] {
					newPolymer = newPolymer + pair[1]
					break
				}
			}
		}
		newPolymer = newPolymer + string(polymer[len(polymer)-1])
		polymer = newPolymer
		output = append(output, fmt.Sprintln(polymer))
	}

	mostCommon := ""
	mostCommonCount := 0
	leastCommon := ""
	leastCommonCount := 0
	for i := 0; i < len(polymer)-1; i++ {
		indexCount := strings.Count(polymer, string(polymer[i]))
		if mostCommon == "" || indexCount > mostCommonCount {
			mostCommon = string(polymer[i])
			mostCommonCount = indexCount
		}
		if leastCommon == "" || indexCount < leastCommonCount {
			leastCommon = string(polymer[i])
			leastCommonCount = indexCount
		}
	}
	output = append(output, fmt.Sprintln("Most Common Element is", mostCommon, "which appears", mostCommonCount, "times"))
	output = append(output, fmt.Sprintln("Least Common Element is", leastCommon, "which appears", leastCommonCount, "times"))
	output = append(output, fmt.Sprintln("Difference is", mostCommonCount-leastCommonCount))

	return output
}

func advent142(input []string) []string {
	output := []string{}
	const steps = 40
	pairsCount := map[string]int{}
	elementsCount := map[string]int{}
	insertionPairs := map[string]string{}

	polymer := ""
	for _, inputLine := range input {
		if strings.Contains(inputLine, " -> ") {
			newPair := strings.Split(inputLine, " -> ")
			insertionPairs[newPair[0]] = newPair[1]
		} else if inputLine != "" {
			polymer = inputLine
		}
	}

	output = append(output, fmt.Sprintln(polymer))

	for i := 0; i < len(polymer)-1; i++ {
		pairsCount[polymer[i:i+2]] = pairsCount[polymer[i:i+2]] + 1
	}

	for i := 0; i < steps; i++ {
		output = append(output, fmt.Sprintln("Step", i+1))
		newPairsCount := map[string]int{}
		for pair, count := range pairsCount {
			newPairsCount[string(pair[0])+insertionPairs[pair]] = newPairsCount[string(pair[0])+insertionPairs[pair]] + count
			newPairsCount[insertionPairs[pair]+string(pair[1])] = newPairsCount[insertionPairs[pair]+string(pair[1])] + count
		}
		pairsCount = newPairsCount
		output = append(output, fmt.Sprintln(pairsCount))
	}

	for pair, count := range pairsCount {
		elementsCount[string(pair[0])] = elementsCount[string(pair[0])] + count
	}
	elementsCount[string(polymer[len(polymer)-1])] = elementsCount[string(polymer[len(polymer)-1])] + 1

	mostCommon := ""
	mostCommonCount := 0
	leastCommon := ""
	leastCommonCount := 0
	for e, ec := range elementsCount {
		if mostCommon == "" || ec > mostCommonCount {
			mostCommon = e
			mostCommonCount = ec
		}
		if leastCommon == "" || ec < leastCommonCount {
			leastCommon = e
			leastCommonCount = ec
		}
	}
	output = append(output, fmt.Sprintln("Most Common Element is", mostCommon, "which appears", mostCommonCount, "times"))
	output = append(output, fmt.Sprintln("Least Common Element is", leastCommon, "which appears", leastCommonCount, "times"))
	output = append(output, fmt.Sprintln("Difference is", mostCommonCount-leastCommonCount))

	return output
}
