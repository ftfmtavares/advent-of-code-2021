package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func advent141(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner
	const steps = 10

	fd, err = os.Open(filename)
	check(err)
	scanner = bufio.NewScanner(fd)

	polymer := ""
	insertionPairs := [][]string{}
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), " -> ") {
			insertionPairs = append(insertionPairs, []string{})
			insertionPairs[len(insertionPairs)-1] = append(insertionPairs[len(insertionPairs)-1], strings.Split(scanner.Text(), " -> ")...)
		} else if scanner.Text() != "" {
			polymer = scanner.Text()
		}
	}
	fd.Close()

	fmt.Println(polymer)
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
		fmt.Println(polymer)
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
	fmt.Println("Most Common Element is", mostCommon, "which appears", mostCommonCount, "times")
	fmt.Println("Least Common Element is", leastCommon, "which appears", leastCommonCount, "times")
	fmt.Println("Difference is", mostCommonCount-leastCommonCount)
}

func advent142(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner
	const steps = 40
	pairsCount := map[string]int{}
	elementsCount := map[string]int{}
	insertionPairs := map[string]string{}

	fd, err = os.Open(filename)
	check(err)
	scanner = bufio.NewScanner(fd)

	polymer := ""
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), " -> ") {
			newPair := strings.Split(scanner.Text(), " -> ")
			insertionPairs[newPair[0]] = newPair[1]
		} else if scanner.Text() != "" {
			polymer = scanner.Text()
		}
	}
	fd.Close()

	fmt.Println(polymer)

	for i := 0; i < len(polymer)-1; i++ {
		pairsCount[polymer[i:i+2]] = pairsCount[polymer[i:i+2]] + 1
	}

	for i := 0; i < steps; i++ {
		fmt.Println("Step", i+1)
		newPairsCount := map[string]int{}
		for pair, count := range pairsCount {
			newPairsCount[string(pair[0])+insertionPairs[pair]] = newPairsCount[string(pair[0])+insertionPairs[pair]] + count
			newPairsCount[insertionPairs[pair]+string(pair[1])] = newPairsCount[insertionPairs[pair]+string(pair[1])] + count
		}
		pairsCount = newPairsCount
		fmt.Println(pairsCount)
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
	fmt.Println("Most Common Element is", mostCommon, "which appears", mostCommonCount, "times")
	fmt.Println("Least Common Element is", leastCommon, "which appears", leastCommonCount, "times")
	fmt.Println("Difference is", mostCommonCount-leastCommonCount)
}
