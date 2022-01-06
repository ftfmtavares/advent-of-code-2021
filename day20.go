package main

import (
	"bufio"
	"fmt"
	"os"
)

func advent201(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner
	const iterations = 2

	getBinary := func(c byte) int {
		if c == '#' {
			return 1
		} else {
			return 0
		}
	}

	fd, err = os.Open(filename)
	check(err)
	scanner = bufio.NewScanner(fd)

	scanner.Scan()
	enhanceString := scanner.Text()

	image := []string{}
	for scanner.Scan() {
		if scanner.Text() != "" {
			image = append(image, scanner.Text())
		}
	}
	fd.Close()

	for i := range image {
		for j := 0; j < iterations; j++ {
			image[i] = "." + image[i] + "."
		}
	}
	newline := ""
	for i := 0; i < len(image[0]); i++ {
		newline = newline + "."
	}
	borderLines := []string{}
	for i := 0; i < iterations; i++ {
		borderLines = append(borderLines, newline)
	}
	image = append(borderLines, image...)
	image = append(image, borderLines...)

	borderPixels := "."
	for i := 0; i < iterations; i++ {
		fmt.Println("Iteration", i+1)
		newImage := []string{}
		for y := 0; y < len(image); y++ {
			newLine := ""
			for x := 0; x < len(image[y]); x++ {
				binaryNum := 0
				if y == 0 || x == 0 {
					binaryNum = binaryNum + getBinary(borderPixels[0])
				} else {
					binaryNum = binaryNum + getBinary(image[y-1][x-1])
				}
				binaryNum = 2 * binaryNum
				if y == 0 {
					binaryNum = binaryNum + getBinary(borderPixels[0])
				} else {
					binaryNum = binaryNum + getBinary(image[y-1][x])
				}
				binaryNum = 2 * binaryNum
				if y == 0 || x == len(image[y])-1 {
					binaryNum = binaryNum + getBinary(borderPixels[0])
				} else {
					binaryNum = binaryNum + getBinary(image[y-1][x+1])
				}
				binaryNum = 2 * binaryNum
				if x == 0 {
					binaryNum = binaryNum + getBinary(borderPixels[0])
				} else {
					binaryNum = binaryNum + getBinary(image[y][x-1])
				}
				binaryNum = 2 * binaryNum
				binaryNum = binaryNum + getBinary(image[y][x])
				binaryNum = 2 * binaryNum
				if x == len(image[y])-1 {
					binaryNum = binaryNum + getBinary(borderPixels[0])
				} else {
					binaryNum = binaryNum + getBinary(image[y][x+1])
				}
				binaryNum = 2 * binaryNum
				if y == len(image)-1 || x == 0 {
					binaryNum = binaryNum + getBinary(borderPixels[0])
				} else {
					binaryNum = binaryNum + getBinary(image[y+1][x-1])
				}
				binaryNum = 2 * binaryNum
				if y == len(image)-1 {
					binaryNum = binaryNum + getBinary(borderPixels[0])
				} else {
					binaryNum = binaryNum + getBinary(image[y+1][x])
				}
				binaryNum = 2 * binaryNum
				if y == len(image)-1 || x == len(image[y])-1 {
					binaryNum = binaryNum + getBinary(borderPixels[0])
				} else {
					binaryNum = binaryNum + getBinary(image[y+1][x+1])
				}

				newLine = newLine + string(enhanceString[binaryNum])
			}
			newImage = append(newImage, newLine)
			fmt.Println(newLine)
		}

		binaryNum := 0
		for j := 0; j < 9; j++ {
			binaryNum = binaryNum * 2
			binaryNum = binaryNum + getBinary(borderPixels[0])
		}
		borderPixels = string(enhanceString[binaryNum])

		for j := 0; j < len(newImage); j++ {
			image[j] = newImage[j]
		}
	}

	litPixels := 0
	for _, s := range image {
		for _, p := range s {
			if p == '#' {
				litPixels++
			}
		}
	}
	fmt.Println("Lit Pixels are", litPixels)
}

func advent202(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner
	const iterations = 50

	getBinary := func(c byte) int {
		if c == '#' {
			return 1
		} else {
			return 0
		}
	}

	fd, err = os.Open(filename)
	check(err)
	scanner = bufio.NewScanner(fd)

	scanner.Scan()
	enhanceString := scanner.Text()

	image := []string{}
	for scanner.Scan() {
		if scanner.Text() != "" {
			image = append(image, scanner.Text())
		}
	}
	fd.Close()

	for i := range image {
		for j := 0; j < iterations; j++ {
			image[i] = "." + image[i] + "."
		}
	}
	newline := ""
	for i := 0; i < len(image[0]); i++ {
		newline = newline + "."
	}
	borderLines := []string{}
	for i := 0; i < iterations; i++ {
		borderLines = append(borderLines, newline)
	}
	image = append(borderLines, image...)
	image = append(image, borderLines...)

	borderPixels := "."
	for i := 0; i < iterations; i++ {
		fmt.Println("Iteration", i+1)
		newImage := []string{}
		for y := 0; y < len(image); y++ {
			newLine := ""
			for x := 0; x < len(image[y]); x++ {
				binaryNum := 0
				if y == 0 || x == 0 {
					binaryNum = binaryNum + getBinary(borderPixels[0])
				} else {
					binaryNum = binaryNum + getBinary(image[y-1][x-1])
				}
				binaryNum = 2 * binaryNum
				if y == 0 {
					binaryNum = binaryNum + getBinary(borderPixels[0])
				} else {
					binaryNum = binaryNum + getBinary(image[y-1][x])
				}
				binaryNum = 2 * binaryNum
				if y == 0 || x == len(image[y])-1 {
					binaryNum = binaryNum + getBinary(borderPixels[0])
				} else {
					binaryNum = binaryNum + getBinary(image[y-1][x+1])
				}
				binaryNum = 2 * binaryNum
				if x == 0 {
					binaryNum = binaryNum + getBinary(borderPixels[0])
				} else {
					binaryNum = binaryNum + getBinary(image[y][x-1])
				}
				binaryNum = 2 * binaryNum
				binaryNum = binaryNum + getBinary(image[y][x])
				binaryNum = 2 * binaryNum
				if x == len(image[y])-1 {
					binaryNum = binaryNum + getBinary(borderPixels[0])
				} else {
					binaryNum = binaryNum + getBinary(image[y][x+1])
				}
				binaryNum = 2 * binaryNum
				if y == len(image)-1 || x == 0 {
					binaryNum = binaryNum + getBinary(borderPixels[0])
				} else {
					binaryNum = binaryNum + getBinary(image[y+1][x-1])
				}
				binaryNum = 2 * binaryNum
				if y == len(image)-1 {
					binaryNum = binaryNum + getBinary(borderPixels[0])
				} else {
					binaryNum = binaryNum + getBinary(image[y+1][x])
				}
				binaryNum = 2 * binaryNum
				if y == len(image)-1 || x == len(image[y])-1 {
					binaryNum = binaryNum + getBinary(borderPixels[0])
				} else {
					binaryNum = binaryNum + getBinary(image[y+1][x+1])
				}

				newLine = newLine + string(enhanceString[binaryNum])
			}
			newImage = append(newImage, newLine)
			fmt.Println(newLine)
		}

		binaryNum := 0
		for j := 0; j < 9; j++ {
			binaryNum = binaryNum * 2
			binaryNum = binaryNum + getBinary(borderPixels[0])
		}
		borderPixels = string(enhanceString[binaryNum])

		for j := 0; j < len(newImage); j++ {
			image[j] = newImage[j]
		}
	}

	litPixels := 0
	for _, s := range image {
		for _, p := range s {
			if p == '#' {
				litPixels++
			}
		}
	}
	fmt.Println("Lit Pixels are", litPixels)
}
