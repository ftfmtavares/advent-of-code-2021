package main

import (
	"fmt"
	"strconv"
	"strings"
)

func advent041(input []string) []string {
	var err error
	output := []string{}
	var bingoNumbers []int
	var bingoNumber int
	const bingoCardRowCount int = 5
	const bingoCardColCount int = 5
	var bingoCards [][bingoCardColCount][bingoCardRowCount]int
	var bingoCardsColCounters [][bingoCardRowCount]int
	var bingoCardsRowCounters [][bingoCardColCount]int
	var winnerFound bool = false
	var remainingNumbersSum int = 0

	f := func(c rune) bool {
		return c == ','
	}
	for _, num := range strings.FieldsFunc(input[0], f) {
		bingoNumber, err = strconv.Atoi(num)
		check(err)
		bingoNumbers = append(bingoNumbers, bingoNumber)
	}

	output = append(output, fmt.Sprintln("Lucky numbers are", len(bingoNumbers), "long"))

	for ind := 1; ind < len(input); ind++ {
		if input[ind] != "" {
			bingoCard := [bingoCardColCount][bingoCardRowCount]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}
			bingoCardsColCounter := [bingoCardRowCount]int{0, 0, 0, 0, 0}
			bingoCardsColCounters = append(bingoCardsColCounters, bingoCardsColCounter)
			bingoCardsRowCounter := [bingoCardColCount]int{0, 0, 0, 0, 0}
			bingoCardsRowCounters = append(bingoCardsRowCounters, bingoCardsRowCounter)
			for i := 0; i < bingoCardRowCount; i++ {
				for j, num := range strings.Fields(input[ind]) {
					bingoCard[i][j], err = strconv.Atoi(num)
					check(err)
				}
				ind++
			}
			bingoCards = append(bingoCards, bingoCard)
			output = append(output, fmt.Sprintln("Bingo Card", len(bingoCards), "is", bingoCard))
		}
	}

	for i := 0; i < len(bingoNumbers); i++ {
		output = append(output, fmt.Sprintln("Number", bingoNumbers[i], "is drawn"))
		for j := 0; j < len(bingoCards); j++ {
			for k := 0; k < bingoCardRowCount; k++ {
				for l := 0; l < bingoCardColCount; l++ {
					if bingoCards[j][k][l] == bingoNumbers[i] {
						bingoCardsColCounters[j][l]++
						bingoCardsRowCounters[j][k]++
						bingoCards[j][k][l] = -1
						output = append(output, fmt.Sprintln("Bingo Card", j, "has it! Row has now", bingoCardsRowCounters[j][k], "and Col has now", bingoCardsColCounters[j][l], "drawn numbers"))
						if (bingoCardsColCounters[j][l] == 5) || (bingoCardsRowCounters[j][l] == 5) {
							output = append(output, fmt.Sprintln("Bingo Card", j, "is the winner!", bingoCards[j]))
							winnerFound = true
							break
						}
					}
				}
				if winnerFound {
					break
				}
			}
			if winnerFound {
				for k := 0; k < bingoCardRowCount; k++ {
					for l := 0; l < bingoCardColCount; l++ {
						if bingoCards[j][k][l] != -1 {
							remainingNumbersSum = remainingNumbersSum + bingoCards[j][k][l]
						}
					}
				}
				output = append(output, fmt.Sprintln("Remaining Numbers Sum is", remainingNumbersSum, "and final result is", remainingNumbersSum*bingoNumbers[i]))
				break
			}
		}
		if winnerFound {
			break
		}
	}

	return output
}

func advent042(input []string) []string {
	var err error
	output := []string{}
	var bingoNumbers []int
	var bingoNumber int
	const bingoCardRowCount int = 5
	const bingoCardColCount int = 5
	var bingoCards [][bingoCardColCount][bingoCardRowCount]int
	var bingoCardsColCounters [][bingoCardRowCount]int
	var bingoCardsRowCounters [][bingoCardColCount]int
	var winnersCount int = 0
	var winners []bool
	var remainingNumbersSum int = 0
	var lastWinner bool = false

	f := func(c rune) bool {
		return c == ','
	}
	for _, num := range strings.FieldsFunc(input[0], f) {
		bingoNumber, err = strconv.Atoi(num)
		check(err)
		bingoNumbers = append(bingoNumbers, bingoNumber)
	}

	output = append(output, fmt.Sprintln("Lucky numbers are", len(bingoNumbers), "long"))

	for ind := 1; ind < len(input); ind++ {
		if input[ind] != "" {
			bingoCard := [bingoCardColCount][bingoCardRowCount]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}
			for i := 0; i < bingoCardRowCount; i++ {
				for j, num := range strings.Fields(input[ind]) {
					bingoCard[i][j], err = strconv.Atoi(num)
					check(err)
				}
				ind++
			}
			bingoCards = append(bingoCards, bingoCard)
			output = append(output, fmt.Sprintln("Bingo Card", len(bingoCards), "is", bingoCard))
			winners = append(winners, false)
			bingoCardsColCounter := [bingoCardRowCount]int{0, 0, 0, 0, 0}
			bingoCardsColCounters = append(bingoCardsColCounters, bingoCardsColCounter)
			bingoCardsRowCounter := [bingoCardColCount]int{0, 0, 0, 0, 0}
			bingoCardsRowCounters = append(bingoCardsRowCounters, bingoCardsRowCounter)
		}
	}

	for i := 0; i < len(bingoNumbers); i++ {
		output = append(output, fmt.Sprintln("Number", bingoNumbers[i], "is drawn"))
		for j := 0; j < len(bingoCards); j++ {
			if winners[j] {
				continue
			}
			for k := 0; k < bingoCardRowCount; k++ {
				for l := 0; l < bingoCardColCount; l++ {
					if bingoCards[j][k][l] == bingoNumbers[i] {
						bingoCardsColCounters[j][l]++
						bingoCardsRowCounters[j][k]++
						bingoCards[j][k][l] = -1
						output = append(output, fmt.Sprintln("Bingo Card", j, "has it! Row has now", bingoCardsRowCounters[j][k], "and Col has now", bingoCardsColCounters[j][l], "drawn numbers"))
						if (bingoCardsColCounters[j][l] == 5) || (bingoCardsRowCounters[j][l] == 5) {
							output = append(output, fmt.Sprintln("Bingo Card", j, "won!", bingoCards[j]))
							winners[j] = true
							winnersCount++
							if winnersCount == len(bingoCards) {
								lastWinner = true
							}
							break
						}
					}
					if lastWinner || winners[j] {
						break
					}
				}
				if lastWinner || winners[j] {
					break
				}
			}
			if lastWinner {
				for k := 0; k < bingoCardRowCount; k++ {
					for l := 0; l < bingoCardColCount; l++ {
						if bingoCards[j][k][l] != -1 {
							remainingNumbersSum = remainingNumbersSum + bingoCards[j][k][l]
						}
					}
				}
				output = append(output, fmt.Sprintln("Remaining Numbers Sum is", remainingNumbersSum, "and final result is", remainingNumbersSum*bingoNumbers[i]))
				break
			}
		}
		if lastWinner {
			break
		}
	}

	return output
}
