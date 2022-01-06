package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func advent041(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner
	var bingoNumbers []int
	var bingoNumber int
	const bingoCardRowCount int = 5
	const bingoCardColCount int = 5
	var bingoCards [][bingoCardColCount][bingoCardRowCount]int
	var bingoCardsColCounters [][bingoCardRowCount]int
	var bingoCardsRowCounters [][bingoCardColCount]int
	var winnerFound bool = false
	var remainingNumbersSum int = 0

	fd, err = os.Open(filename)
	check(err)

	scanner = bufio.NewScanner(fd)

	scanner.Scan()
	f := func(c rune) bool {
		return c == ','
	}
	for _, num := range strings.FieldsFunc(scanner.Text(), f) {
		bingoNumber, err = strconv.Atoi(num)
		check(err)
		bingoNumbers = append(bingoNumbers, bingoNumber)
	}

	fmt.Println("Lucky numbers are", len(bingoNumbers), "long")

	for scanner.Scan() {
		if scanner.Text() != "" {
			bingoCard := [bingoCardColCount][bingoCardRowCount]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}
			bingoCardsColCounter := [bingoCardRowCount]int{0, 0, 0, 0, 0}
			bingoCardsColCounters = append(bingoCardsColCounters, bingoCardsColCounter)
			bingoCardsRowCounter := [bingoCardColCount]int{0, 0, 0, 0, 0}
			bingoCardsRowCounters = append(bingoCardsRowCounters, bingoCardsRowCounter)
			for i := 0; i < bingoCardRowCount; i++ {
				for j, num := range strings.Fields(scanner.Text()) {
					bingoCard[i][j], err = strconv.Atoi(num)
					check(err)
				}
				scanner.Scan()
			}
			bingoCards = append(bingoCards, bingoCard)
			fmt.Println("Bingo Card", len(bingoCards), "is", bingoCard)
		}
	}

	fd.Close()

	for i := 0; i < len(bingoNumbers); i++ {
		fmt.Println("Number", bingoNumbers[i], "is drawn")
		for j := 0; j < len(bingoCards); j++ {
			for k := 0; k < bingoCardRowCount; k++ {
				for l := 0; l < bingoCardColCount; l++ {
					if bingoCards[j][k][l] == bingoNumbers[i] {
						bingoCardsColCounters[j][l]++
						bingoCardsRowCounters[j][k]++
						bingoCards[j][k][l] = -1
						fmt.Println("Bingo Card", j, "has it! Row has now", bingoCardsRowCounters[j][k], "and Col has now", bingoCardsColCounters[j][l], "drawn numbers")
						if (bingoCardsColCounters[j][l] == 5) || (bingoCardsRowCounters[j][l] == 5) {
							fmt.Println("Bingo Card", j, "is the winner!", bingoCards[j])
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
				fmt.Println("Remaining Numbers Sum is", remainingNumbersSum, "and final result is", remainingNumbersSum*bingoNumbers[i])
				break
			}
		}
		if winnerFound {
			break
		}
	}
}

func advent042(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner
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

	fd, err = os.Open(filename)
	check(err)

	scanner = bufio.NewScanner(fd)

	scanner.Scan()
	f := func(c rune) bool {
		return c == ','
	}
	for _, num := range strings.FieldsFunc(scanner.Text(), f) {
		bingoNumber, err = strconv.Atoi(num)
		check(err)
		bingoNumbers = append(bingoNumbers, bingoNumber)
	}

	fmt.Println("Lucky numbers are", len(bingoNumbers), "long")

	for scanner.Scan() {
		if scanner.Text() != "" {
			bingoCard := [bingoCardColCount][bingoCardRowCount]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}
			for i := 0; i < bingoCardRowCount; i++ {
				for j, num := range strings.Fields(scanner.Text()) {
					bingoCard[i][j], err = strconv.Atoi(num)
					check(err)
				}
				scanner.Scan()
			}
			bingoCards = append(bingoCards, bingoCard)
			fmt.Println("Bingo Card", len(bingoCards), "is", bingoCard)
			winners = append(winners, false)
			bingoCardsColCounter := [bingoCardRowCount]int{0, 0, 0, 0, 0}
			bingoCardsColCounters = append(bingoCardsColCounters, bingoCardsColCounter)
			bingoCardsRowCounter := [bingoCardColCount]int{0, 0, 0, 0, 0}
			bingoCardsRowCounters = append(bingoCardsRowCounters, bingoCardsRowCounter)
		}
	}

	fd.Close()

	for i := 0; i < len(bingoNumbers); i++ {
		fmt.Println("Number", bingoNumbers[i], "is drawn")
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
						fmt.Println("Bingo Card", j, "has it! Row has now", bingoCardsRowCounters[j][k], "and Col has now", bingoCardsColCounters[j][l], "drawn numbers")
						if (bingoCardsColCounters[j][l] == 5) || (bingoCardsRowCounters[j][l] == 5) {
							fmt.Println("Bingo Card", j, "won!", bingoCards[j])
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
				fmt.Println("Remaining Numbers Sum is", remainingNumbersSum, "and final result is", remainingNumbersSum*bingoNumbers[i])
				break
			}
		}
		if lastWinner {
			break
		}
	}
}
