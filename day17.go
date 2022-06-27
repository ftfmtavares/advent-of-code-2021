package main

import (
	"fmt"
	"strconv"
	"strings"
)

func advent171(input []string) []string {
	var err error
	output := []string{}
	var startX int
	var endX int
	var startY int
	var endY int

	procInput := input[0]

	procInput = strings.Replace(procInput, "target area: ", "", -1)
	targetParts := strings.Split(procInput, ", ")
	targetParts[0] = strings.Replace(targetParts[0], "x=", "", -1)
	xRange := strings.Split(targetParts[0], "..")
	startX, err = strconv.Atoi(xRange[0])
	check(err)
	endX, err = strconv.Atoi(xRange[1])
	check(err)
	targetParts[1] = strings.Replace(targetParts[1], "y=", "", -1)
	yRange := strings.Split(targetParts[1], "..")
	startY, err = strconv.Atoi(yRange[0])
	check(err)
	endY, err = strconv.Atoi(yRange[1])
	check(err)

	output = append(output, fmt.Sprintln("Target Range is X:", startX, "-", endX, " Y:", startY, "-", endY))

	highPoint := 0
	speedY := -startY - 1
	for {
		if speedY > 0 {
			highPoint = highPoint + speedY
			speedY--
			output = append(output, fmt.Sprintln("Rising at", highPoint, "with speed at", speedY))
		} else {
			break
		}
	}

	output = append(output, fmt.Sprintln("Highest Point is at", highPoint))

	return output
}

func advent172(input []string) []string {
	var err error
	output := []string{}
	var startX int
	var endX int
	var startY int
	var endY int
	shots := []string{}

	procInput := input[0]

	procInput = strings.Replace(procInput, "target area: ", "", -1)
	targetParts := strings.Split(procInput, ", ")
	targetParts[0] = strings.Replace(targetParts[0], "x=", "", -1)
	xRange := strings.Split(targetParts[0], "..")
	startX, err = strconv.Atoi(xRange[0])
	check(err)
	endX, err = strconv.Atoi(xRange[1])
	check(err)
	targetParts[1] = strings.Replace(targetParts[1], "y=", "", -1)
	yRange := strings.Split(targetParts[1], "..")
	startY, err = strconv.Atoi(yRange[0])
	check(err)
	endY, err = strconv.Atoi(yRange[1])
	check(err)

	output = append(output, fmt.Sprintln("Target Range is X:", startX, "-", endX, " Y:", startY, "-", endY))

	for speedY := startY; speedY <= absolute(startY); speedY++ {
		currentSpeedY := speedY
		posY := 0
		step := 0
		for posY >= startY {
			posY = posY + currentSpeedY
			currentSpeedY--
			step++
			if posY >= startY && posY <= endY {
				for speedX := 1; speedX <= endX; speedX++ {
					currentSpeedX := speedX
					posX := 0
					for i := 0; i < step; i++ {
						posX = posX + currentSpeedX
						currentSpeedX--
						if currentSpeedX == 0 {
							break
						}
					}
					if posX >= startX && posX <= endX {
						exists := false
						shot := strconv.Itoa(speedX) + "," + strconv.Itoa(speedY)
						for _, s := range shots {
							if s == shot {
								exists = true
							}
						}
						if !exists {
							shots = append(shots, shot)
							output = append(output, fmt.Sprintln("Valid Shot ", shot))
						}
					}
				}
			}
		}
	}

	output = append(output, fmt.Sprintln("Possible shots are", len(shots)))

	return output
}
