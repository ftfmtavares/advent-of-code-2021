package main

import (
	"fmt"
	"strconv"
)

func advent181(input []string) []string {
	output := []string{}
	const maxNestedPairs = 4
	const splitValue = 10
	var magnitude int

	type pair struct {
		value1 int
		value2 int
		pair1  *pair
		pair2  *pair
	}

	var printPair func(p pair) string
	var readPair func(expression1, expression2 string) pair
	var addLeft func(p *pair, add int) int
	var addRight func(p *pair, add int) int
	var explodePair func(p *pair, level int) (int, int, int)
	var splitNumbers func(p *pair, level int) (int, int)
	var reducePair func(p *pair, splitNumber int, explodeNumber int) int
	var getMagnitude func(p *pair) int

	printPair = func(p pair) string {
		result := ""
		if p.value1 != -1 {
			result = result + strconv.Itoa(p.value1)
		} else {
			result = result + "[" + printPair(*p.pair1) + "]"
		}
		result = result + ","
		if p.value2 != -1 {
			result = result + strconv.Itoa(p.value2)
		} else {
			result = result + "[" + printPair(*p.pair2) + "]"
		}
		return result
	}

	getBracketContent := func(expression string) (string, string) {
		if expression[0] != '[' {
			return "", ""
		}
		bracketCount := 0
		start := 1
		middle := -1
		end := -1
		for i, c := range expression {
			if c == '[' {
				bracketCount++
			} else if c == ']' {
				bracketCount--
				if bracketCount == 0 {
					end = i
					break
				}
			} else if c == ',' && bracketCount == 1 {
				middle = i
			}
		}
		return expression[start:middle], expression[middle+1 : end]
	}

	readPair = func(expression1, expression2 string) pair {
		result := pair{
			value1: -1,
			value2: -1,
			pair1:  nil,
			pair2:  nil,
		}
		if expression1[0] == '[' {
			nestedElement1, nestedElement2 := getBracketContent(expression1)
			newpairElement1 := readPair(nestedElement1, nestedElement2)
			result.pair1 = &newpairElement1
		} else {
			result.value1 = int(expression1[0] - '0')
		}
		if expression2[0] == '[' {
			nestedElement1, nestedElement2 := getBracketContent(expression2)
			newpairElement2 := readPair(nestedElement1, nestedElement2)
			result.pair2 = &newpairElement2
		} else {
			result.value2 = int(expression2[0] - '0')
		}
		return result
	}

	addLeft = func(p *pair, add int) int {
		res := 0
		if p.value2 != -1 {
			p.value2 = p.value2 + add
			if p.value2 >= splitValue {
				res++
			}
		} else {
			res = addLeft(p.pair2, add)
		}
		return res
	}

	addRight = func(p *pair, add int) int {
		res := 0
		if p.value1 != -1 {
			p.value1 = p.value1 + add
			if p.value1 >= splitValue {
				res++
			}
		} else {
			res = addRight(p.pair1, add)
		}
		return res
	}

	explodePair = func(p *pair, level int) (int, int, int) {
		res1 := -1
		res2 := -1
		res3 := 0
		if level >= maxNestedPairs {
			if p.value1 == -1 && p.pair1.value1 != -1 && p.pair1.value2 != -1 {
				res1 = p.pair1.value1
				res2 = p.pair1.value2
				p.pair1 = nil
				p.value1 = 0
				if p.value2 != -1 {
					p.value2 = p.value2 + res2
					if p.value2 >= splitValue {
						res3++
					}
				} else {
					res3 = res3 + addRight(p.pair2, res2)
				}
				res2 = 0
				return res1, res2, res3
			}
			if p.value2 == -1 && p.pair2.value1 != -1 && p.pair2.value2 != -1 {
				res1 = p.pair2.value1
				res2 = p.pair2.value2
				p.pair2 = nil
				p.value2 = 0
				if p.value1 != -1 {
					p.value1 = p.value1 + res1
					if p.value1 >= splitValue {
						res3++
					}
				} else {
					res3 = res3 + addLeft(p.pair1, res1)
				}
				res1 = 0
				return res1, res2, res3
			}
		}

		if p.value1 == -1 {
			res1, res2, res3 = explodePair(p.pair1, level+1)
			if res2 > 0 {
				if p.value2 != -1 {
					p.value2 = p.value2 + res2
					if p.value2 >= splitValue {
						res3++
					}
				} else {
					res3 = res3 + addRight(p.pair2, res2)
				}
				res2 = 0
			}
			if res1 > -1 || res2 > -1 {
				return res1, res2, res3
			}
		}

		if p.value2 == -1 {
			res1, res2, res3 = explodePair(p.pair2, level+1)
			if res1 > 0 {
				if p.value1 != -1 {
					p.value1 = p.value1 + res1
					if p.value1 >= splitValue {
						res3++
					}
				} else {
					res3 = res3 + addLeft(p.pair1, res1)
				}
				res1 = 0
			}
			if res1 > -1 || res2 > -1 {
				return res1, res2, res3
			}
		}

		return res1, res2, res3
	}

	splitNumbers = func(p *pair, level int) (int, int) {
		res1 := -1
		res2 := 0

		if p.value1 == -1 {
			res1, res2 = splitNumbers(p.pair1, level+1)
			if res1 == 0 {
				return res1, res2
			}
		} else if p.value1 >= splitValue {
			newPair := pair{
				value1: -1,
				value2: -1,
				pair1:  nil,
				pair2:  nil,
			}
			newPair.value1 = p.value1 / 2
			newPair.value2 = p.value1 - p.value1/2
			p.value1 = -1
			p.pair1 = &newPair
			res1 = 0
			if level >= maxNestedPairs {
				res2 = 1
			}
			return res1, res2
		}

		if p.value2 == -1 {
			res1, res2 = splitNumbers(p.pair2, level+1)
			if res1 == 0 {
				return res1, res2
			}
		} else if p.value2 >= splitValue {
			newPair := pair{
				value1: -1,
				value2: -1,
				pair1:  nil,
				pair2:  nil,
			}
			newPair.value1 = p.value2 / 2
			newPair.value2 = p.value2 - p.value2/2
			p.value2 = -1
			p.pair2 = &newPair
			res1 = 0
			if level >= maxNestedPairs {
				res2 = 1
			}
			return res1, res2
		}

		return res1, res2
	}

	reducePair = func(p *pair, splitNumber int, explodeNumber int) int {
		res1 := -1
		res2 := -1
		res3 := 0
		if splitNumber == -1 && explodeNumber == -1 {
			res1, res2, _ = explodePair(p, 1)
			output = append(output, fmt.Sprintln(printPair(*p)))
			if res1 == -1 && res2 == -1 {
				return -1
			} else {
				return 0
			}
		} else if splitNumber > 0 {
			res3, explodeNumber = splitNumbers(p, 1)
			output = append(output, fmt.Sprintln(printPair(*p)))
			for explodeNumber > 0 {
				reducePair(p, -1, explodeNumber)
				explodeNumber--
			}
			return res3
		} else if explodeNumber > 0 {
			res1, res2, splitNumber = explodePair(p, 1)
			output = append(output, fmt.Sprintln(printPair(*p)))
			for splitNumber > 0 {
				reducePair(p, splitNumber, -1)
				splitNumber--
			}
			if res1 == -1 && res2 == -1 {
				return -1
			} else {
				return 0
			}
		}
		return -1
	}

	getMagnitude = func(p *pair) int {
		result := 0
		if p.value1 == -1 {
			result = result + 3*getMagnitude(p.pair1)
		} else {
			result = result + 3*p.value1
		}
		if p.value2 == -1 {
			result = result + 2*getMagnitude(p.pair2)
		} else {
			result = result + 2*p.value2
		}
		return result
	}

	stringElement1, stringElement2 := getBracketContent(input[0])
	tempResult := readPair(stringElement1, stringElement2)
	result := &tempResult
	output = append(output, fmt.Sprintln(input[0]))
	output = append(output, fmt.Sprintln(printPair(*result)))
	for ind := 1; ind < len(input); ind++ {
		stringElement1, stringElement2 := getBracketContent(input[ind])
		addingPair := readPair(stringElement1, stringElement2)
		output = append(output, fmt.Sprintln(input[ind]))
		output = append(output, fmt.Sprintln(printPair(addingPair)))

		newResultPair := pair{
			value1: -1,
			value2: -1,
			pair1:  nil,
			pair2:  nil,
		}
		newResultPair.pair1 = result
		newResultPair.pair2 = &addingPair
		result = &newResultPair
		output = append(output, fmt.Sprintln(printPair(*result)))

		res := 0
		for res == 0 {
			res = reducePair(result, -1, -1)
		}
		res = 0
		for res == 0 {
			res = reducePair(result, 1, -1)
		}
		output = append(output, fmt.Sprintln("End of Addition ", printPair(*result)))
	}

	magnitude = getMagnitude(result)
	output = append(output, fmt.Sprintln("Magnitude is ", magnitude))

	return output
}

func advent182(input []string) []string {
	output := []string{}
	const maxNestedPairs = 4
	const splitValue = 10
	var maxMagnitude int

	type pair struct {
		value1 int
		value2 int
		pair1  *pair
		pair2  *pair
	}

	var printPair func(p pair) string
	var readPair func(expression1, expression2 string) pair
	var addLeft func(p *pair, add int) int
	var addRight func(p *pair, add int) int
	var explodePair func(p *pair, level int) (int, int, int)
	var splitNumbers func(p *pair, level int) (int, int)
	var reducePair func(p *pair, splitNumber int, explodeNumber int) int
	var getMagnitude func(p *pair) int

	printPair = func(p pair) string {
		result := ""
		if p.value1 != -1 {
			result = result + strconv.Itoa(p.value1)
		} else {
			result = result + "[" + printPair(*p.pair1) + "]"
		}
		result = result + ","
		if p.value2 != -1 {
			result = result + strconv.Itoa(p.value2)
		} else {
			result = result + "[" + printPair(*p.pair2) + "]"
		}
		return result
	}

	getBracketContent := func(expression string) (string, string) {
		if expression[0] != '[' {
			return "", ""
		}
		bracketCount := 0
		start := 1
		middle := -1
		end := -1
		for i, c := range expression {
			if c == '[' {
				bracketCount++
			} else if c == ']' {
				bracketCount--
				if bracketCount == 0 {
					end = i
					break
				}
			} else if c == ',' && bracketCount == 1 {
				middle = i
			}
		}
		return expression[start:middle], expression[middle+1 : end]
	}

	readPair = func(expression1, expression2 string) pair {
		result := pair{
			value1: -1,
			value2: -1,
			pair1:  nil,
			pair2:  nil,
		}
		if expression1[0] == '[' {
			nestedElement1, nestedElement2 := getBracketContent(expression1)
			newpairElement1 := readPair(nestedElement1, nestedElement2)
			result.pair1 = &newpairElement1
		} else {
			result.value1 = int(expression1[0] - '0')
		}
		if expression2[0] == '[' {
			nestedElement1, nestedElement2 := getBracketContent(expression2)
			newpairElement2 := readPair(nestedElement1, nestedElement2)
			result.pair2 = &newpairElement2
		} else {
			result.value2 = int(expression2[0] - '0')
		}
		return result
	}

	addLeft = func(p *pair, add int) int {
		res := 0
		if p.value2 != -1 {
			p.value2 = p.value2 + add
			if p.value2 >= splitValue {
				res++
			}
		} else {
			res = addLeft(p.pair2, add)
		}
		return res
	}

	addRight = func(p *pair, add int) int {
		res := 0
		if p.value1 != -1 {
			p.value1 = p.value1 + add
			if p.value1 >= splitValue {
				res++
			}
		} else {
			res = addRight(p.pair1, add)
		}
		return res
	}

	explodePair = func(p *pair, level int) (int, int, int) {
		res1 := -1
		res2 := -1
		res3 := 0
		if level >= maxNestedPairs {
			if p.value1 == -1 && p.pair1.value1 != -1 && p.pair1.value2 != -1 {
				res1 = p.pair1.value1
				res2 = p.pair1.value2
				p.pair1 = nil
				p.value1 = 0
				if p.value2 != -1 {
					p.value2 = p.value2 + res2
					if p.value2 >= splitValue {
						res3++
					}
				} else {
					res3 = res3 + addRight(p.pair2, res2)
				}
				res2 = 0
				return res1, res2, res3
			}
			if p.value2 == -1 && p.pair2.value1 != -1 && p.pair2.value2 != -1 {
				res1 = p.pair2.value1
				res2 = p.pair2.value2
				p.pair2 = nil
				p.value2 = 0
				if p.value1 != -1 {
					p.value1 = p.value1 + res1
					if p.value1 >= splitValue {
						res3++
					}
				} else {
					res3 = res3 + addLeft(p.pair1, res1)
				}
				res1 = 0
				return res1, res2, res3
			}
		}

		if p.value1 == -1 {
			res1, res2, res3 = explodePair(p.pair1, level+1)
			if res2 > 0 {
				if p.value2 != -1 {
					p.value2 = p.value2 + res2
					if p.value2 >= splitValue {
						res3++
					}
				} else {
					res3 = res3 + addRight(p.pair2, res2)
				}
				res2 = 0
			}
			if res1 > -1 || res2 > -1 {
				return res1, res2, res3
			}
		}

		if p.value2 == -1 {
			res1, res2, res3 = explodePair(p.pair2, level+1)
			if res1 > 0 {
				if p.value1 != -1 {
					p.value1 = p.value1 + res1
					if p.value1 >= splitValue {
						res3++
					}
				} else {
					res3 = res3 + addLeft(p.pair1, res1)
				}
				res1 = 0
			}
			if res1 > -1 || res2 > -1 {
				return res1, res2, res3
			}
		}

		return res1, res2, res3
	}

	splitNumbers = func(p *pair, level int) (int, int) {
		res1 := -1
		res2 := 0

		if p.value1 == -1 {
			res1, res2 = splitNumbers(p.pair1, level+1)
			if res1 == 0 {
				return res1, res2
			}
		} else if p.value1 >= splitValue {
			newPair := pair{
				value1: -1,
				value2: -1,
				pair1:  nil,
				pair2:  nil,
			}
			newPair.value1 = p.value1 / 2
			newPair.value2 = p.value1 - p.value1/2
			p.value1 = -1
			p.pair1 = &newPair
			res1 = 0
			if level >= maxNestedPairs {
				res2 = 1
			}
			return res1, res2
		}

		if p.value2 == -1 {
			res1, res2 = splitNumbers(p.pair2, level+1)
			if res1 == 0 {
				return res1, res2
			}
		} else if p.value2 >= splitValue {
			newPair := pair{
				value1: -1,
				value2: -1,
				pair1:  nil,
				pair2:  nil,
			}
			newPair.value1 = p.value2 / 2
			newPair.value2 = p.value2 - p.value2/2
			p.value2 = -1
			p.pair2 = &newPair
			res1 = 0
			if level >= maxNestedPairs {
				res2 = 1
			}
			return res1, res2
		}

		return res1, res2
	}

	reducePair = func(p *pair, splitNumber int, explodeNumber int) int {
		res1 := -1
		res2 := -1
		res3 := 0
		if splitNumber == -1 && explodeNumber == -1 {
			res1, res2, _ = explodePair(p, 1)
			output = append(output, fmt.Sprintln(printPair(*p)))
			if res1 == -1 && res2 == -1 {
				return -1
			} else {
				return 0
			}
		} else if splitNumber > 0 {
			res3, explodeNumber = splitNumbers(p, 1)
			output = append(output, fmt.Sprintln(printPair(*p)))
			for explodeNumber > 0 {
				reducePair(p, -1, explodeNumber)
				explodeNumber--
			}
			return res3
		} else if explodeNumber > 0 {
			res1, res2, splitNumber = explodePair(p, 1)
			output = append(output, fmt.Sprintln(printPair(*p)))
			for splitNumber > 0 {
				reducePair(p, splitNumber, -1)
				splitNumber--
			}
			if res1 == -1 && res2 == -1 {
				return -1
			} else {
				return 0
			}
		}
		return -1
	}

	getMagnitude = func(p *pair) int {
		result := 0
		if p.value1 == -1 {
			result = result + 3*getMagnitude(p.pair1)
		} else {
			result = result + 3*p.value1
		}
		if p.value2 == -1 {
			result = result + 2*getMagnitude(p.pair2)
		} else {
			result = result + 2*p.value2
		}
		return result
	}

	numbers := []string{}
	numbers = append(numbers, input...)

	maxMagnitude = 0
	for i, num1 := range numbers {
		for j, num2 := range numbers {
			if i != j {
				stringElement1, stringElement2 := getBracketContent(num1)
				tempResult := readPair(stringElement1, stringElement2)
				result := &tempResult
				output = append(output, fmt.Sprintln(input[len(input)-1]))
				output = append(output, fmt.Sprintln(printPair(*result)))

				stringElement1, stringElement2 = getBracketContent(num2)
				addingPair := readPair(stringElement1, stringElement2)
				output = append(output, fmt.Sprintln(input[len(input)-1]))
				output = append(output, fmt.Sprintln(printPair(addingPair)))

				newResultPair := pair{
					value1: -1,
					value2: -1,
					pair1:  nil,
					pair2:  nil,
				}
				newResultPair.pair1 = result
				newResultPair.pair2 = &addingPair
				result = &newResultPair
				output = append(output, fmt.Sprintln(printPair(*result)))

				res := 0
				for res == 0 {
					res = reducePair(result, -1, -1)
				}
				res = 0
				for res == 0 {
					res = reducePair(result, 1, -1)
				}
				output = append(output, fmt.Sprintln("End of Addition ", printPair(*result)))

				maxMagnitude = max(getMagnitude(result), maxMagnitude)
			}
		}
	}
	output = append(output, fmt.Sprintln("Maximum Magnitude is ", maxMagnitude))

	return output
}
