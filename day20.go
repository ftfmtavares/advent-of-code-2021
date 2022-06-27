package main

import (
	"fmt"
)

func advent201(input []string) []string {
	output := []string{}
	const iterations = 2

	getBinary := func(c byte) int {
		if c == '#' {
			return 1
		} else {
			return 0
		}
	}

	enhanceString := input[0]

	image := []string{}
	for ind := 1; ind < len(input); ind++ {
		if input[ind] != "" {
			image = append(image, input[ind])
		}
	}

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
		output = append(output, fmt.Sprintln("Iteration", i+1))
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
			output = append(output, fmt.Sprintln(newLine))
		}

		binaryNum := 0
		for j := 0; j < 9; j++ {
			binaryNum = binaryNum * 2
			binaryNum = binaryNum + getBinary(borderPixels[0])
		}
		borderPixels = string(enhanceString[binaryNum])

		copy(image, newImage)
	}

	litPixels := 0
	for _, s := range image {
		for _, p := range s {
			if p == '#' {
				litPixels++
			}
		}
	}
	output = append(output, fmt.Sprintln("Lit Pixels are", litPixels))

	return output
}

func advent202(input []string) []string {
	output := []string{}
	const iterations = 50

	getBinary := func(c byte) int {
		if c == '#' {
			return 1
		} else {
			return 0
		}
	}

	enhanceString := input[0]

	image := []string{}
	for ind := 1; ind < len(input); ind++ {
		if input[ind] != "" {
			image = append(image, input[ind])
		}
	}

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
		output = append(output, fmt.Sprintln("Iteration", i+1))
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
			output = append(output, fmt.Sprintln(newLine))
		}

		binaryNum := 0
		for j := 0; j < 9; j++ {
			binaryNum = binaryNum * 2
			binaryNum = binaryNum + getBinary(borderPixels[0])
		}
		borderPixels = string(enhanceString[binaryNum])

		copy(image, newImage)
	}

	litPixels := 0
	for _, s := range image {
		for _, p := range s {
			if p == '#' {
				litPixels++
			}
		}
	}
	output = append(output, fmt.Sprintln("Lit Pixels are", litPixels))

	return output
}
