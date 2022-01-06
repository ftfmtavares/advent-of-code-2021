package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func advent211(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner

	fd, err = os.Open(filename)
	check(err)
	scanner = bufio.NewScanner(fd)

	scanner.Scan()
	line := scanner.Text()
	line = strings.Replace(line, "Player 1 starting position: ", "", -1)
	pos1, err := strconv.Atoi(line)
	check(err)
	scanner.Scan()
	line = scanner.Text()
	line = strings.Replace(line, "Player 2 starting position: ", "", -1)
	pos2, err := strconv.Atoi(line)
	check(err)

	fd.Close()

	playerScore1 := 0
	playerScore2 := 0
	die := 1
	rollsCount := 0
	for playerScore1 < 1000 && playerScore2 < 1000 {
		for i := 0; i < 3; i++ {
			pos1 = pos1 + die
			fmt.Println("Player 1 moves", die, "steps")
			if die == 100 {
				die = 1
			} else {
				die++
			}
			rollsCount++
		}
		pos1 = (pos1-1)%10 + 1
		playerScore1 = playerScore1 + pos1
		fmt.Println("Player 1 score increases to", playerScore1, "after stoping on step", pos1)
		if playerScore1 >= 1000 {
			break
		}
		for i := 0; i < 3; i++ {
			pos2 = pos2 + die
			fmt.Println("Player 2 moves", die, "steps")
			if die == 100 {
				die = 1
			} else {
				die++
			}
			rollsCount++
		}
		pos2 = (pos2-1)%10 + 1
		playerScore2 = playerScore2 + pos2
		fmt.Println("Player 2 score increases to", playerScore2, "after stoping on step", pos2)
	}

	fmt.Println("We have a winner with scores as player 1=", playerScore1, " and player 2=", playerScore2)
	fmt.Println("Die was rolled", rollsCount, "times")
	fmt.Println("Result is", min(playerScore1, playerScore2)*rollsCount)
}

func advent212(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner

	var playGame func(pos1, pos2, score1, score2, turn int) (int, int)
	playGame = func(pos1, pos2, score1, score2, turn int) (int, int) {
		//		fmt.Println("New universe with player 1 on", pos1, "and", score1, "points and player 2 on", pos2, "and", score2, "points")
		wins1 := 0
		wins2 := 0
		newPos1 := pos1
		newPos2 := pos2
		newScore1 := score1
		newScore2 := score2

		if turn == 1 {
			newPos1 = pos1 + 3
			newPos1 = (newPos1-1)%10 + 1
			newScore1 = score1 + newPos1
			if newScore1 >= 21 {
				wins1 = wins1 + 1
			} else {
				newWins1, newWins2 := playGame(newPos1, newPos2, newScore1, newScore2, 2)
				wins1 = wins1 + newWins1*1
				wins2 = wins2 + newWins2*1
			}

			newPos1 = pos1 + 4
			newPos1 = (newPos1-1)%10 + 1
			newScore1 = score1 + newPos1
			if newScore1 >= 21 {
				wins1 = wins1 + 3
			} else {
				newWins1, newWins2 := playGame(newPos1, newPos2, newScore1, newScore2, 2)
				wins1 = wins1 + newWins1*3
				wins2 = wins2 + newWins2*3
			}

			newPos1 = pos1 + 5
			newPos1 = (newPos1-1)%10 + 1
			newScore1 = score1 + newPos1
			if newScore1 >= 21 {
				wins1 = wins1 + 6
			} else {
				newWins1, newWins2 := playGame(newPos1, newPos2, newScore1, newScore2, 2)
				wins1 = wins1 + newWins1*6
				wins2 = wins2 + newWins2*6
			}

			newPos1 = pos1 + 6
			newPos1 = (newPos1-1)%10 + 1
			newScore1 = score1 + newPos1
			if newScore1 >= 21 {
				wins1 = wins1 + 7
			} else {
				newWins1, newWins2 := playGame(newPos1, newPos2, newScore1, newScore2, 2)
				wins1 = wins1 + newWins1*7
				wins2 = wins2 + newWins2*7
			}

			newPos1 = pos1 + 7
			newPos1 = (newPos1-1)%10 + 1
			newScore1 = score1 + newPos1
			if newScore1 >= 21 {
				wins1 = wins1 + 6
			} else {
				newWins1, newWins2 := playGame(newPos1, newPos2, newScore1, newScore2, 2)
				wins1 = wins1 + newWins1*6
				wins2 = wins2 + newWins2*6
			}

			newPos1 = pos1 + 8
			newPos1 = (newPos1-1)%10 + 1
			newScore1 = score1 + newPos1
			if newScore1 >= 21 {
				wins1 = wins1 + 3
			} else {
				newWins1, newWins2 := playGame(newPos1, newPos2, newScore1, newScore2, 2)
				wins1 = wins1 + newWins1*3
				wins2 = wins2 + newWins2*3
			}

			newPos1 = pos1 + 9
			newPos1 = (newPos1-1)%10 + 1
			newScore1 = score1 + newPos1
			if newScore1 >= 21 {
				wins1 = wins1 + 1
			} else {
				newWins1, newWins2 := playGame(newPos1, newPos2, newScore1, newScore2, 2)
				wins1 = wins1 + newWins1*1
				wins2 = wins2 + newWins2*1
			}
		} else {
			newPos2 = pos2 + 3
			newPos2 = (newPos2-1)%10 + 1
			newScore2 = score2 + newPos2
			if newScore2 >= 21 {
				wins2 = wins2 + 1
			} else {
				newWins1, newWins2 := playGame(newPos1, newPos2, newScore1, newScore2, 1)
				wins1 = wins1 + newWins1*1
				wins2 = wins2 + newWins2*1
			}

			newPos2 = pos2 + 4
			newPos2 = (newPos2-1)%10 + 1
			newScore2 = score2 + newPos2
			if newScore2 >= 21 {
				wins2 = wins2 + 3
			} else {
				newWins1, newWins2 := playGame(newPos1, newPos2, newScore1, newScore2, 1)
				wins1 = wins1 + newWins1*3
				wins2 = wins2 + newWins2*3
			}

			newPos2 = pos2 + 5
			newPos2 = (newPos2-1)%10 + 1
			newScore2 = score2 + newPos2
			if newScore2 >= 21 {
				wins2 = wins2 + 6
			} else {
				newWins1, newWins2 := playGame(newPos1, newPos2, newScore1, newScore2, 1)
				wins1 = wins1 + newWins1*6
				wins2 = wins2 + newWins2*6
			}

			newPos2 = pos2 + 6
			newPos2 = (newPos2-1)%10 + 1
			newScore2 = score2 + newPos2
			if newScore2 >= 21 {
				wins2 = wins2 + 7
			} else {
				newWins1, newWins2 := playGame(newPos1, newPos2, newScore1, newScore2, 1)
				wins1 = wins1 + newWins1*7
				wins2 = wins2 + newWins2*7
			}

			newPos2 = pos2 + 7
			newPos2 = (newPos2-1)%10 + 1
			newScore2 = score2 + newPos2
			if newScore2 >= 21 {
				wins2 = wins2 + 6
			} else {
				newWins1, newWins2 := playGame(newPos1, newPos2, newScore1, newScore2, 1)
				wins1 = wins1 + newWins1*6
				wins2 = wins2 + newWins2*6
			}

			newPos2 = pos2 + 8
			newPos2 = (newPos2-1)%10 + 1
			newScore2 = score2 + newPos2
			if newScore2 >= 21 {
				wins2 = wins2 + 3
			} else {
				newWins1, newWins2 := playGame(newPos1, newPos2, newScore1, newScore2, 1)
				wins1 = wins1 + newWins1*3
				wins2 = wins2 + newWins2*3
			}

			newPos2 = pos2 + 9
			newPos2 = (newPos2-1)%10 + 1
			newScore2 = score2 + newPos2
			if newScore2 >= 21 {
				wins2 = wins2 + 1
			} else {
				newWins1, newWins2 := playGame(newPos1, newPos2, newScore1, newScore2, 1)
				wins1 = wins1 + newWins1*1
				wins2 = wins2 + newWins2*1
			}
		}

		return wins1, wins2
	}

	fd, err = os.Open(filename)
	check(err)
	scanner = bufio.NewScanner(fd)

	scanner.Scan()
	line := scanner.Text()
	line = strings.Replace(line, "Player 1 starting position: ", "", -1)
	pos1, err := strconv.Atoi(line)
	check(err)
	scanner.Scan()
	line = scanner.Text()
	line = strings.Replace(line, "Player 2 starting position: ", "", -1)
	pos2, err := strconv.Atoi(line)
	check(err)

	fd.Close()

	wins1, wins2 := playGame(pos1, pos2, 0, 0, 1)

	fmt.Println("Player 1 wins at", wins1, "universes")
	fmt.Println("Player 2 wins at", wins2, "universes")
}
