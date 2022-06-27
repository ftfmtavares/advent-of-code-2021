package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func absolute(a int) int {
	if a < 0 {
		a = -1 * a
	}
	return a
}

func readTextFile(filename string) []string {
	fd, err := os.Open(filename)
	check(err)
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	textContent := []string{}
	for scanner.Scan() {
		textContent = append(textContent, scanner.Text())
	}

	return textContent
}
