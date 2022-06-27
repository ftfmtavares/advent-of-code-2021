package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	fmt.Println("Advent of Code 2021 - Puzzle Solver")
	fmt.Println("")

	day := flag.Int("day", 0, "Specify the puzzle day (1-25)")
	part := flag.Int("part", 0, "Specify which part of the puzzle (1-2)")
	filename := flag.String("input-file", "", "Specify the filename containing the puzzle input")
	flag.Parse()
	valid := true
	if *day == 0 || *part == 0 || *filename == "" {
		flag.PrintDefaults()
		valid = false
	} else if *day < 1 || *day > 25 {
		fmt.Println("Day number must be between 1 and 25")
		fmt.Println("")
		flag.PrintDefaults()
		valid = false
	} else if *part < 1 || *part > 2 {
		fmt.Println("Part must be between 1 and 2")
		fmt.Println("")
		flag.PrintDefaults()
		valid = false
	} else if !fs.ValidPath(*filename) {
		fmt.Println("Invalid Input-File name")
		fmt.Println("")
		flag.PrintDefaults()
		valid = false
	} else {
		_, err := os.Stat(*filename)
		if err != nil {
			fmt.Println(err)
			fmt.Println("")
			flag.PrintDefaults()
			valid = false
		}
	}
	if valid {
		if *day == 25 && *part == 2 {
			fmt.Println("There is no part 2 on puzzle day 25")
			fmt.Println("")
		} else {
			text := "Solving Puzzle " + fmt.Sprint(*day) + ", Part " + fmt.Sprint(*part) + " using \"" + *filename + "\" as the puzzle input"
			fmt.Println(text)
			fmt.Println("")
		}

		switch *day {
		case 1:
			if *part == 1 {
				advent011(*filename)
			} else {
				advent012(*filename)
			}
		case 2:
			if *part == 1 {
				advent021(*filename)
			} else {
				advent022(*filename)
			}
		case 3:
			if *part == 1 {
				advent031(*filename)
			} else {
				advent032(*filename)
			}
		case 4:
			if *part == 1 {
				advent041(*filename)
			} else {
				advent042(*filename)
			}
		case 5:
			if *part == 1 {
				advent051(*filename)
			} else {
				advent052(*filename)
			}
		case 6:
			if *part == 1 {
				advent061(*filename)
			} else {
				advent062(*filename)
			}
		case 7:
			if *part == 1 {
				advent071(*filename)
			} else {
				advent072(*filename)
			}
		case 8:
			if *part == 1 {
				advent081(*filename)
			} else {
				advent082(*filename)
			}
		case 9:
			if *part == 1 {
				advent091(*filename)
			} else {
				advent092(*filename)
			}
		case 10:
			if *part == 1 {
				advent101(*filename)
			} else {
				advent102(*filename)
			}
		case 11:
			if *part == 1 {
				advent111(*filename)
			} else {
				advent112(*filename)
			}
		case 12:
			if *part == 1 {
				advent121(*filename)
			} else {
				advent122(*filename)
			}
		case 13:
			if *part == 1 {
				advent131(*filename)
			} else {
				advent132(*filename)
			}
		case 14:
			if *part == 1 {
				advent141(*filename)
			} else {
				advent142(*filename)
			}
		case 15:
			if *part == 1 {
				advent151(*filename)
			} else {
				advent152(*filename)
			}
		case 16:
			if *part == 1 {
				advent161(*filename)
			} else {
				advent162(*filename)
			}
		case 17:
			if *part == 1 {
				advent171(*filename)
			} else {
				advent172(*filename)
			}
		case 18:
			if *part == 1 {
				advent181(*filename)
			} else {
				advent182(*filename)
			}
		case 19:
			if *part == 1 {
				advent191(*filename)
			} else {
				advent192(*filename)
			}
		case 20:
			if *part == 1 {
				advent201(*filename)
			} else {
				advent202(*filename)
			}
		case 21:
			if *part == 1 {
				advent211(*filename)
			} else {
				advent212(*filename)
			}
		case 22:
			if *part == 1 {
				advent221(*filename)
			} else {
				advent222(*filename)
			}
		case 23:
			if *part == 1 {
				advent231(*filename)
			} else {
				advent232(*filename)
			}
		case 24:
			if *part == 1 {
				advent241(*filename)
			} else {
				advent242(*filename)
			}
		case 25:
			if *part == 1 {
				advent251(*filename)
			}
		}
	}
}
