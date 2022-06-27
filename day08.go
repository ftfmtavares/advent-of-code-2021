package main

import (
	"fmt"
	"strings"
)

func advent081(input []string) []string {
	output := []string{}
	var knownDigits int = 0

	for _, inputLine := range input {
		lineParts := strings.Split(inputLine, "|")
		output := strings.Split(lineParts[1], " ")
		for _, s := range output {
			if len(s) == 2 || len(s) == 3 || len(s) == 4 || len(s) == 7 {
				knownDigits++
			}
		}
	}

	output = append(output, fmt.Sprintln("Number of easy digits is", knownDigits))

	return output
}

func advent082(input []string) []string {
	output := []string{}
	var outputSum int = 0

	for _, inputLine := range input {
		lineParts := strings.Split(inputLine, "|")
		patterns := strings.Split(strings.Trim(lineParts[0], " "), " ")
		digits2 := ""
		digits3 := ""
		digits4 := ""
		digits5 := []string{}
		digits6 := []string{}
		digits7 := ""
		for _, s := range patterns {
			if len(s) == 2 {
				digits2 = s
			} else if len(s) == 3 {
				digits3 = s
			} else if len(s) == 4 {
				digits4 = s
			} else if len(s) == 5 {
				digits5 = append(digits5, s)
			} else if len(s) == 6 {
				digits6 = append(digits6, s)
			} else if len(s) == 7 {
				digits7 = s
			}
		}
		output = append(output, fmt.Sprintln(digits2, digits3, digits4, digits5, digits6, digits7))
		links := [7]string{"", "", "", "", "", "", ""}

		var digits string

		digits = digits3
		for i := 0; i < len(digits2); i++ {
			digits = strings.ReplaceAll(digits, string(digits2[i]), "")
		}
		links[0] = digits

		for j := 0; j < len(digits5); j++ {
			digits := digits5[j]
			digits = strings.ReplaceAll(digits, links[0], "")
			for i := 0; i < len(digits4); i++ {
				digits = strings.ReplaceAll(digits, string(digits4[i]), "")
			}
			if len(digits) == 1 {
				links[6] = digits
				break
			}
		}

		for j := 0; j < len(digits6); j++ {
			digits := digits6[j]
			digits = strings.ReplaceAll(digits, links[0], "")
			digits = strings.ReplaceAll(digits, links[6], "")
			for i := 0; i < len(digits4); i++ {
				digits = strings.ReplaceAll(digits, string(digits4[i]), "")
			}
			if len(digits) == 1 {
				links[4] = digits
				break
			}
		}

		for j := 0; j < len(digits5); j++ {
			digits := digits5[j]
			digits = strings.ReplaceAll(digits, links[0], "")
			digits = strings.ReplaceAll(digits, links[6], "")
			digits = strings.ReplaceAll(digits, links[4], "")
			for i := 0; i < len(digits2); i++ {
				digits = strings.ReplaceAll(digits, string(digits2[i]), "")
			}
			if len(digits) == 1 {
				links[3] = digits
				break
			}
		}

		digits = digits7
		digits = strings.ReplaceAll(digits, links[0], "")
		digits = strings.ReplaceAll(digits, links[6], "")
		digits = strings.ReplaceAll(digits, links[4], "")
		digits = strings.ReplaceAll(digits, links[3], "")
		for i := 0; i < len(digits2); i++ {
			digits = strings.ReplaceAll(digits, string(digits2[i]), "")
		}
		links[1] = digits

		for j := 0; j < len(digits6); j++ {
			digits := digits6[j]
			digits = strings.ReplaceAll(digits, links[0], "")
			digits = strings.ReplaceAll(digits, links[6], "")
			digits = strings.ReplaceAll(digits, links[4], "")
			digits = strings.ReplaceAll(digits, links[3], "")
			digits = strings.ReplaceAll(digits, links[1], "")
			if len(digits) == 1 {
				links[5] = digits
				break
			}
		}

		digits = digits2
		digits = strings.ReplaceAll(digits, links[0], "")
		digits = strings.ReplaceAll(digits, links[6], "")
		digits = strings.ReplaceAll(digits, links[4], "")
		digits = strings.ReplaceAll(digits, links[3], "")
		digits = strings.ReplaceAll(digits, links[1], "")
		digits = strings.ReplaceAll(digits, links[5], "")
		links[2] = digits

		output = append(output, fmt.Sprintln(links))

		encodedDigits := strings.Split(strings.Trim(lineParts[1], " "), " ")
		decodedDigits := []int{}

		for _, s := range encodedDigits {
			if strings.Contains(s, links[0]) && strings.Contains(s, links[1]) && strings.Contains(s, links[2]) && !strings.Contains(s, links[3]) && strings.Contains(s, links[4]) && strings.Contains(s, links[5]) && strings.Contains(s, links[6]) {
				decodedDigits = append(decodedDigits, 0)
			} else if !strings.Contains(s, links[0]) && !strings.Contains(s, links[1]) && strings.Contains(s, links[2]) && !strings.Contains(s, links[3]) && !strings.Contains(s, links[4]) && strings.Contains(s, links[5]) && !strings.Contains(s, links[6]) {
				decodedDigits = append(decodedDigits, 1)
			} else if strings.Contains(s, links[0]) && !strings.Contains(s, links[1]) && strings.Contains(s, links[2]) && strings.Contains(s, links[3]) && strings.Contains(s, links[4]) && !strings.Contains(s, links[5]) && strings.Contains(s, links[6]) {
				decodedDigits = append(decodedDigits, 2)
			} else if strings.Contains(s, links[0]) && !strings.Contains(s, links[1]) && strings.Contains(s, links[2]) && strings.Contains(s, links[3]) && !strings.Contains(s, links[4]) && strings.Contains(s, links[5]) && strings.Contains(s, links[6]) {
				decodedDigits = append(decodedDigits, 3)
			} else if !strings.Contains(s, links[0]) && strings.Contains(s, links[1]) && strings.Contains(s, links[2]) && strings.Contains(s, links[3]) && !strings.Contains(s, links[4]) && strings.Contains(s, links[5]) && !strings.Contains(s, links[6]) {
				decodedDigits = append(decodedDigits, 4)
			} else if strings.Contains(s, links[0]) && strings.Contains(s, links[1]) && !strings.Contains(s, links[2]) && strings.Contains(s, links[3]) && !strings.Contains(s, links[4]) && strings.Contains(s, links[5]) && strings.Contains(s, links[6]) {
				decodedDigits = append(decodedDigits, 5)
			} else if strings.Contains(s, links[0]) && strings.Contains(s, links[1]) && !strings.Contains(s, links[2]) && strings.Contains(s, links[3]) && strings.Contains(s, links[4]) && strings.Contains(s, links[5]) && strings.Contains(s, links[6]) {
				decodedDigits = append(decodedDigits, 6)
			} else if strings.Contains(s, links[0]) && !strings.Contains(s, links[1]) && strings.Contains(s, links[2]) && !strings.Contains(s, links[3]) && !strings.Contains(s, links[4]) && strings.Contains(s, links[5]) && !strings.Contains(s, links[6]) {
				decodedDigits = append(decodedDigits, 7)
			} else if strings.Contains(s, links[0]) && strings.Contains(s, links[1]) && strings.Contains(s, links[2]) && strings.Contains(s, links[3]) && strings.Contains(s, links[4]) && strings.Contains(s, links[5]) && strings.Contains(s, links[6]) {
				decodedDigits = append(decodedDigits, 8)
			} else if strings.Contains(s, links[0]) && strings.Contains(s, links[1]) && strings.Contains(s, links[2]) && strings.Contains(s, links[3]) && !strings.Contains(s, links[4]) && strings.Contains(s, links[5]) && strings.Contains(s, links[6]) {
				decodedDigits = append(decodedDigits, 9)
			}
		}

		procOutput := 0
		for i := 0; i < len(decodedDigits); i++ {
			procOutput = procOutput * 10
			procOutput = procOutput + decodedDigits[i]
		}

		output = append(output, fmt.Sprintln(encodedDigits, decodedDigits, procOutput))

		outputSum = outputSum + procOutput
	}

	output = append(output, fmt.Sprintln("Sum of all output digits is", outputSum))

	return output
}
