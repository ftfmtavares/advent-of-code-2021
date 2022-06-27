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

		input := readTextFile(*filename)
		output := []string{}

		switch *day {
		case 1:
			if *part == 1 {
				output = advent011(input)
			} else {
				output = advent012(input)
			}
		case 2:
			if *part == 1 {
				output = advent021(input)
			} else {
				output = advent022(input)
			}
		case 3:
			if *part == 1 {
				output = advent031(input)
			} else {
				output = advent032(input)
			}
		case 4:
			if *part == 1 {
				output = advent041(input)
			} else {
				output = advent042(input)
			}
		case 5:
			if *part == 1 {
				output = advent051(input)
			} else {
				output = advent052(input)
			}
		case 6:
			if *part == 1 {
				output = advent061(input)
			} else {
				output = advent062(input)
			}
		case 7:
			if *part == 1 {
				output = advent071(input)
			} else {
				output = advent072(input)
			}
		case 8:
			if *part == 1 {
				output = advent081(input)
			} else {
				output = advent082(input)
			}
		case 9:
			if *part == 1 {
				output = advent091(input)
			} else {
				output = advent092(input)
			}
		case 10:
			if *part == 1 {
				output = advent101(input)
			} else {
				output = advent102(input)
			}
		case 11:
			if *part == 1 {
				output = advent111(input)
			} else {
				output = advent112(input)
			}
		case 12:
			if *part == 1 {
				output = advent121(input)
			} else {
				output = advent122(input)
			}
		case 13:
			if *part == 1 {
				output = advent131(input)
			} else {
				output = advent132(input)
			}
		case 14:
			if *part == 1 {
				output = advent141(input)
			} else {
				output = advent142(input)
			}
		case 15:
			if *part == 1 {
				output = advent151(input)
			} else {
				output = advent152(input)
			}
		case 16:
			if *part == 1 {
				output = advent161(input)
			} else {
				output = advent162(input)
			}
		case 17:
			if *part == 1 {
				output = advent171(input)
			} else {
				output = advent172(input)
			}
		case 18:
			if *part == 1 {
				output = advent181(input)
			} else {
				output = advent182(input)
			}
		case 19:
			if *part == 1 {
				output = advent191(input)
			} else {
				output = advent192(input)
			}
		case 20:
			if *part == 1 {
				output = advent201(input)
			} else {
				output = advent202(input)
			}
		case 21:
			if *part == 1 {
				output = advent211(input)
			} else {
				output = advent212(input)
			}
		case 22:
			if *part == 1 {
				output = advent221(input)
			} else {
				output = advent222(input)
			}
		case 23:
			if *part == 1 {
				output = advent231(input)
			} else {
				output = advent232(input)
			}
		case 24:
			if *part == 1 {
				output = advent241(input)
			} else {
				output = advent242(input)
			}
		case 25:
			if *part == 1 {
				output = advent251(input)
			}
		}

		for _, line := range output {
			fmt.Println(line)
		}
	}
}
