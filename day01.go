package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func advent011(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner
	var prev int
	var curr int
	var count int = 0

	fd, err = os.Open(filename)
	check(err)

	scanner = bufio.NewScanner(fd)

	scanner.Scan()
	prev, err = strconv.Atoi(scanner.Text())
	check(err)
	fmt.Printf("First value to be compared: %d\n", prev)

	for scanner.Scan() {
		curr, err = strconv.Atoi(scanner.Text())
		check(err)
		if curr > prev {
			count++
		}
		fmt.Printf("Comparison of %d > %d is %t and Count=%d\n", curr, prev, curr > prev, count)
		prev = curr
	}

	fd.Close()
}

func advent012(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner
	const slidingWindow int = 3
	var prev int
	var curr int
	var windowArray [slidingWindow]int
	var count int = 0
	var readValue int

	fd, err = os.Open(filename)
	check(err)

	scanner = bufio.NewScanner(fd)

	for i := 0; i < slidingWindow; i++ {
		scanner.Scan()
		readValue, err = strconv.Atoi(scanner.Text())
		check(err)
		windowArray[i] = readValue
	}

	for scanner.Scan() {
		readValue, err = strconv.Atoi(scanner.Text())
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
		fmt.Printf("Comparison of %d > %d is %t and Count=%d\n", curr, prev, curr > prev, count)
	}

	fd.Close()
}
