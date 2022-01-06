package main

import (
	"bufio"
	"fmt"
	"os"
)

func advent031(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner
	const numberOfBits int = 12
	var bitCounters [numberOfBits]int
	var count int = 0
	var readBinary string
	var gamma int = 0
	var epsilon int = 0

	for i := 0; i < numberOfBits; i++ {
		bitCounters[i] = 0
	}

	fd, err = os.Open(filename)
	check(err)

	scanner = bufio.NewScanner(fd)

	for scanner.Scan() {
		readBinary = scanner.Text()
		for i := 0; i < len(readBinary); i++ {
			if readBinary[i] == '1' {
				bitCounters[i]++
			}
		}
		count++
		fmt.Println("Counters are", bitCounters, "after", count, "reads")
	}

	fd.Close()

	for i := 0; i < numberOfBits; i++ {
		gamma = gamma * 2
		epsilon = epsilon * 2
		if bitCounters[i] > count/2 {
			gamma++
		} else {
			epsilon++
		}
	}
	fmt.Println("Gamma is", gamma, "and epsilon is", epsilon, "making the power consumption equal to", gamma*epsilon)
}

func advent032(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner
	var diagnosticValues []string
	const numberOfBits int = 12
	var mostCommonBit string
	var leastCommonBit string
	var mostCount int
	var leastCount int
	var mostBitCount int
	var leastBitCount int
	var matchMostCommonBit bool
	var matchLeastCommonBit bool
	var oxygenString string
	var co2String string
	var oxygen int = 0
	var co2 int = 0

	fd, err = os.Open(filename)
	check(err)

	scanner = bufio.NewScanner(fd)

	for scanner.Scan() {
		diagnosticValues = append(diagnosticValues, scanner.Text())
	}

	fd.Close()

	fmt.Println("Diagnostic Values are", len(diagnosticValues), "long")

	mostCommonBit = ""
	leastCommonBit = ""
	for i := 0; i < numberOfBits; i++ {
		mostCount = 0
		leastCount = 0
		mostBitCount = 0
		leastBitCount = 0
		for j := 0; j < len(diagnosticValues); j++ {
			matchMostCommonBit = true
			matchLeastCommonBit = true
			for k := 0; k < i; k++ {
				if diagnosticValues[j][k] != mostCommonBit[k] {
					matchMostCommonBit = false
				}
				if diagnosticValues[j][k] != leastCommonBit[k] {
					matchLeastCommonBit = false
				}
				if !matchMostCommonBit && !matchLeastCommonBit {
					break
				}
			}
			if matchMostCommonBit {
				oxygenString = diagnosticValues[j]
				mostCount++
				if diagnosticValues[j][i] == '1' {
					mostBitCount++
				}
			}
			if matchLeastCommonBit {
				co2String = diagnosticValues[j]
				leastCount++
				if diagnosticValues[j][i] == '1' {
					leastBitCount++
				}
			}
		}
		fmt.Println(mostCount, "values match oxygen", mostCommonBit, "pattern with", mostBitCount, "1s and", mostCount-mostBitCount, "0s")
		fmt.Println(leastCount, "values match co2", leastCommonBit, "pattern with", leastBitCount, "1s and", leastCount-leastBitCount, "0s")
		if mostBitCount*2 >= mostCount {
			mostCommonBit = mostCommonBit + "1"
		} else {
			mostCommonBit = mostCommonBit + "0"
		}
		if leastBitCount*2 >= leastCount {
			leastCommonBit = leastCommonBit + "0"
		} else {
			leastCommonBit = leastCommonBit + "1"
		}
		if (mostCount <= 1) && (leastCount <= 1) {
			break
		}
	}

	for i := 0; i < numberOfBits; i++ {
		oxygen = oxygen * 2
		co2 = co2 * 2
		if oxygenString[i] == '1' {
			oxygen++
		}
		if co2String[i] == '1' {
			co2++
		}
	}
	fmt.Println("Found", oxygenString, "last value for Oxygen and", co2String, "last value for CO2")
	fmt.Println("Converted values are", oxygen, "Oxygen and", co2, "CO2. Final solution is", oxygen*co2)
}
