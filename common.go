package main

import "fmt"

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
