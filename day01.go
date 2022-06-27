package main

import (
	"fmt"
	"strconv"
)

func advent011(input []string) []string {
	var err error
	output := []string{}
	var prev int
	var curr int
	var count int = 0

	prev, err = strconv.Atoi(input[0])
	check(err)
	output = append(output, fmt.Sprintf("First value to be compared: %d", prev))

	for i := 1; i < len(input); i++ {
		curr, err = strconv.Atoi(input[i])
		check(err)
		if curr > prev {
			count++
		}
		output = append(output, fmt.Sprintf("Comparison of %d > %d is %t and Count=%d", curr, prev, curr > prev, count))
		prev = curr
	}

	return output
}

func advent012(input []string) []string {
	var err error
	output := []string{}
	const slidingWindow int = 3
	var prev int
	var curr int
	var windowArray [slidingWindow]int
	var count int = 0
	var readValue int

	for i := 0; i < slidingWindow; i++ {
		readValue, err = strconv.Atoi(input[i])
		check(err)
		windowArray[i] = readValue
	}

	for ind := slidingWindow; ind < len(input); ind++ {
		readValue, err = strconv.Atoi(input[ind])
		check(err)
		prev = 0
		curr = 0
		prev = prev + windowArray[0]
		for i := 1; i < slidingWindow; i++ {
			prev = prev + windowArray[i]
			curr = curr + windowArray[i]
			windowArray[i-1] = windowArray[i]
		}
		curr = curr + readValue
		windowArray[slidingWindow-1] = readValue
		if curr > prev {
			count++
		}
		output = append(output, fmt.Sprintf("Comparison of %d > %d is %t and Count=%d\n", curr, prev, curr > prev, count))
	}

	return output
}
