package main

import (
	"fmt"
	"strconv"
	"strings"
)

func advent191(input []string) []string {
	var err error
	output := []string{}
	var scanners [][][3]int
	var finalBeacons [][3]int
	numberOfScanners := 0

	for _, inputLine := range input {
		line := inputLine
		if line != "" && line[:3] == "---" {
			output = append(output, fmt.Sprintln("New Scanner", numberOfScanners))
			scanners = append(scanners, [][3]int{})
			numberOfScanners++
		} else if line != "" {
			scanners[numberOfScanners-1] = append(scanners[numberOfScanners-1], [3]int{0, 0, 0})
			coords := strings.Split(line, ",")
			scanners[numberOfScanners-1][len(scanners[numberOfScanners-1])-1][0], err = strconv.Atoi(coords[0])
			check(err)
			scanners[numberOfScanners-1][len(scanners[numberOfScanners-1])-1][1], err = strconv.Atoi(coords[1])
			check(err)
			scanners[numberOfScanners-1][len(scanners[numberOfScanners-1])-1][2], err = strconv.Atoi(coords[2])
			check(err)
			output = append(output, fmt.Sprintln(scanners[numberOfScanners-1][len(scanners[numberOfScanners-1])-1][0], scanners[numberOfScanners-1][len(scanners[numberOfScanners-1])-1][1], scanners[numberOfScanners-1][len(scanners[numberOfScanners-1])-1][2]))
		}
	}

	for _, b := range scanners[0] {
		finalBeacons = append(finalBeacons, [3]int{0, 0, 0})
		finalBeacons[len(finalBeacons)-1][0] = b[0]
		finalBeacons[len(finalBeacons)-1][1] = b[1]
		finalBeacons[len(finalBeacons)-1][2] = b[2]
	}
	output = append(output, fmt.Sprintln("Final Beacons", finalBeacons))

	scanners = scanners[1:]
	for len(scanners) > 0 {
		for i := 0; i < len(scanners); i++ {
			output = append(output, fmt.Sprintln("Trying Scanner", i+1))
			scanner := scanners[i]
			overlap := false
			for alignment := 1; alignment <= 24; alignment++ {
				output = append(output, fmt.Sprintln("Trying Alignment", alignment))
				alignedScanner := [][3]int{}
				for _, beacon := range scanner {
					newAlignedBeacon := [3]int{}
					switch alignment {
					case 1: //x y z
						newAlignedBeacon[0] = beacon[0]
						newAlignedBeacon[1] = beacon[1]
						newAlignedBeacon[2] = beacon[2]
					case 2: //-y x z
						newAlignedBeacon[0] = -beacon[1]
						newAlignedBeacon[1] = beacon[0]
						newAlignedBeacon[2] = beacon[2]
					case 3: //-x -y z
						newAlignedBeacon[0] = -beacon[0]
						newAlignedBeacon[1] = -beacon[1]
						newAlignedBeacon[2] = beacon[2]
					case 4: //y -x z
						newAlignedBeacon[0] = beacon[1]
						newAlignedBeacon[1] = -beacon[0]
						newAlignedBeacon[2] = beacon[2]
					case 5: //z y -x
						newAlignedBeacon[0] = beacon[2]
						newAlignedBeacon[1] = beacon[1]
						newAlignedBeacon[2] = -beacon[0]
					case 6: //-x y -z
						newAlignedBeacon[0] = -beacon[0]
						newAlignedBeacon[1] = beacon[1]
						newAlignedBeacon[2] = -beacon[2]
					case 7: //-z y x
						newAlignedBeacon[0] = -beacon[2]
						newAlignedBeacon[1] = beacon[1]
						newAlignedBeacon[2] = beacon[0]
					case 8: //x -z y
						newAlignedBeacon[0] = beacon[0]
						newAlignedBeacon[1] = -beacon[2]
						newAlignedBeacon[2] = beacon[1]
					case 9: //x -y -z
						newAlignedBeacon[0] = beacon[0]
						newAlignedBeacon[1] = -beacon[1]
						newAlignedBeacon[2] = -beacon[2]
					case 10: //x z -y
						newAlignedBeacon[0] = beacon[0]
						newAlignedBeacon[1] = beacon[2]
						newAlignedBeacon[2] = -beacon[1]
					case 11: //z x y
						newAlignedBeacon[0] = beacon[2]
						newAlignedBeacon[1] = beacon[0]
						newAlignedBeacon[2] = beacon[1]
					case 12: //y x -z
						newAlignedBeacon[0] = beacon[1]
						newAlignedBeacon[1] = beacon[0]
						newAlignedBeacon[2] = -beacon[2]
					case 13: //-z x -y
						newAlignedBeacon[0] = -beacon[2]
						newAlignedBeacon[1] = beacon[0]
						newAlignedBeacon[2] = -beacon[1]
					case 14: //-y -z x
						newAlignedBeacon[0] = -beacon[1]
						newAlignedBeacon[1] = -beacon[2]
						newAlignedBeacon[2] = beacon[0]
					case 15: //-y -x -z
						newAlignedBeacon[0] = -beacon[1]
						newAlignedBeacon[1] = -beacon[0]
						newAlignedBeacon[2] = -beacon[2]
					case 16: //-y z -x
						newAlignedBeacon[0] = -beacon[1]
						newAlignedBeacon[1] = beacon[2]
						newAlignedBeacon[2] = -beacon[0]
					case 17: //z -y x
						newAlignedBeacon[0] = beacon[2]
						newAlignedBeacon[1] = -beacon[1]
						newAlignedBeacon[2] = beacon[0]
					case 18: //-z -y -x
						newAlignedBeacon[0] = -beacon[2]
						newAlignedBeacon[1] = -beacon[1]
						newAlignedBeacon[2] = -beacon[0]
					case 19: //-x -z -y
						newAlignedBeacon[0] = -beacon[0]
						newAlignedBeacon[1] = -beacon[2]
						newAlignedBeacon[2] = -beacon[1]
					case 20: //-x z y
						newAlignedBeacon[0] = -beacon[0]
						newAlignedBeacon[1] = beacon[2]
						newAlignedBeacon[2] = beacon[1]
					case 21: //z -x -y
						newAlignedBeacon[0] = beacon[2]
						newAlignedBeacon[1] = -beacon[0]
						newAlignedBeacon[2] = -beacon[1]
					case 22: //-z -x y
						newAlignedBeacon[0] = -beacon[2]
						newAlignedBeacon[1] = -beacon[0]
						newAlignedBeacon[2] = beacon[1]
					case 23: //y -z -x
						newAlignedBeacon[0] = beacon[1]
						newAlignedBeacon[1] = -beacon[2]
						newAlignedBeacon[2] = -beacon[0]
					case 24: //y z x
						newAlignedBeacon[0] = beacon[1]
						newAlignedBeacon[1] = beacon[2]
						newAlignedBeacon[2] = beacon[0]
					}
					alignedScanner = append(alignedScanner, newAlignedBeacon)
				}

				for _, beacon := range finalBeacons {
					for _, b1 := range alignedScanner {
						scannerPosX := beacon[0] - b1[0]
						scannerPosY := beacon[1] - b1[1]
						scannerPosZ := beacon[2] - b1[2]

						correctedScanner := [][3]int{}
						for _, b2 := range alignedScanner {
							newCorrectedBeacon := [3]int{}
							newCorrectedBeacon[0] = b2[0] + scannerPosX
							newCorrectedBeacon[1] = b2[1] + scannerPosY
							newCorrectedBeacon[2] = b2[2] + scannerPosZ
							correctedScanner = append(correctedScanner, newCorrectedBeacon)
						}

						commonBeacons := 0
						for _, b2 := range correctedScanner {
							foundBeacon := false
							for _, beacon2 := range finalBeacons {
								if beacon2[0] == b2[0] && beacon2[1] == b2[1] && beacon2[2] == b2[2] {
									foundBeacon = true
									break
								}
							}
							if foundBeacon {
								commonBeacons++
							} else if b2[0] <= 1000 && b2[0] >= -1000 && b2[1] <= 1000 && b2[1] >= -1000 && b2[2] <= 1000 && b2[2] >= -1000 {
								break
							}
						}

						if commonBeacons >= 12 {
							output = append(output, fmt.Sprintln("Found Match"))
							overlap = true
							for _, b2 := range correctedScanner {
								foundBeacon := false
								for _, beacon2 := range finalBeacons {
									if beacon2[0] == b2[0] && beacon2[1] == b2[1] && beacon2[2] == b2[2] {
										foundBeacon = true
										break
									}
								}
								if !foundBeacon {
									finalBeacons = append(finalBeacons, [3]int{b2[0], b2[1], b2[2]})
								}
							}
							break
						}
					}
					if overlap {
						break
					}
				}
				if overlap {
					break
				}
			}
			if overlap {
				scanners = append(scanners[:i], scanners[i+1:]...)
				i--
			}
		}
	}

	for _, b := range finalBeacons {
		output = append(output, fmt.Sprintln(b))
	}
	output = append(output, fmt.Sprintln("Number of Beacons is", len(finalBeacons)))

	return output
}

func advent192(input []string) []string {
	var err error
	output := []string{}
	var scanners [][][3]int
	var finalBeacons [][3]int
	numberOfScanners := 0
	var scannersPos [][3]int

	for _, inputLine := range input {
		line := inputLine
		if line != "" && line[:3] == "---" {
			output = append(output, fmt.Sprintln("New Scanner", numberOfScanners))
			scanners = append(scanners, [][3]int{})
			numberOfScanners++
		} else if line != "" {
			scanners[numberOfScanners-1] = append(scanners[numberOfScanners-1], [3]int{0, 0, 0})
			coords := strings.Split(line, ",")
			scanners[numberOfScanners-1][len(scanners[numberOfScanners-1])-1][0], err = strconv.Atoi(coords[0])
			check(err)
			scanners[numberOfScanners-1][len(scanners[numberOfScanners-1])-1][1], err = strconv.Atoi(coords[1])
			check(err)
			scanners[numberOfScanners-1][len(scanners[numberOfScanners-1])-1][2], err = strconv.Atoi(coords[2])
			check(err)
			output = append(output, fmt.Sprintln(scanners[numberOfScanners-1][len(scanners[numberOfScanners-1])-1][0], scanners[numberOfScanners-1][len(scanners[numberOfScanners-1])-1][1], scanners[numberOfScanners-1][len(scanners[numberOfScanners-1])-1][2]))
		}
	}

	for _, b := range scanners[0] {
		finalBeacons = append(finalBeacons, [3]int{0, 0, 0})
		finalBeacons[len(finalBeacons)-1][0] = b[0]
		finalBeacons[len(finalBeacons)-1][1] = b[1]
		finalBeacons[len(finalBeacons)-1][2] = b[2]
	}
	output = append(output, fmt.Sprintln("Final Beacons", finalBeacons))
	scannersPos = append(scannersPos, [3]int{0, 0, 0})

	scanners = scanners[1:]
	for len(scanners) > 0 {
		for i := 0; i < len(scanners); i++ {
			output = append(output, fmt.Sprintln("Trying Scanner", i+1))
			scanner := scanners[i]
			overlap := false
			for alignment := 1; alignment <= 24; alignment++ {
				output = append(output, fmt.Sprintln("Trying Alignment", alignment))
				alignedScanner := [][3]int{}
				for _, beacon := range scanner {
					newAlignedBeacon := [3]int{}
					switch alignment {
					case 1: //x y z
						newAlignedBeacon[0] = beacon[0]
						newAlignedBeacon[1] = beacon[1]
						newAlignedBeacon[2] = beacon[2]
					case 2: //-y x z
						newAlignedBeacon[0] = -beacon[1]
						newAlignedBeacon[1] = beacon[0]
						newAlignedBeacon[2] = beacon[2]
					case 3: //-x -y z
						newAlignedBeacon[0] = -beacon[0]
						newAlignedBeacon[1] = -beacon[1]
						newAlignedBeacon[2] = beacon[2]
					case 4: //y -x z
						newAlignedBeacon[0] = beacon[1]
						newAlignedBeacon[1] = -beacon[0]
						newAlignedBeacon[2] = beacon[2]
					case 5: //z y -x
						newAlignedBeacon[0] = beacon[2]
						newAlignedBeacon[1] = beacon[1]
						newAlignedBeacon[2] = -beacon[0]
					case 6: //-x y -z
						newAlignedBeacon[0] = -beacon[0]
						newAlignedBeacon[1] = beacon[1]
						newAlignedBeacon[2] = -beacon[2]
					case 7: //-z y x
						newAlignedBeacon[0] = -beacon[2]
						newAlignedBeacon[1] = beacon[1]
						newAlignedBeacon[2] = beacon[0]
					case 8: //x -z y
						newAlignedBeacon[0] = beacon[0]
						newAlignedBeacon[1] = -beacon[2]
						newAlignedBeacon[2] = beacon[1]
					case 9: //x -y -z
						newAlignedBeacon[0] = beacon[0]
						newAlignedBeacon[1] = -beacon[1]
						newAlignedBeacon[2] = -beacon[2]
					case 10: //x z -y
						newAlignedBeacon[0] = beacon[0]
						newAlignedBeacon[1] = beacon[2]
						newAlignedBeacon[2] = -beacon[1]
					case 11: //z x y
						newAlignedBeacon[0] = beacon[2]
						newAlignedBeacon[1] = beacon[0]
						newAlignedBeacon[2] = beacon[1]
					case 12: //y x -z
						newAlignedBeacon[0] = beacon[1]
						newAlignedBeacon[1] = beacon[0]
						newAlignedBeacon[2] = -beacon[2]
					case 13: //-z x -y
						newAlignedBeacon[0] = -beacon[2]
						newAlignedBeacon[1] = beacon[0]
						newAlignedBeacon[2] = -beacon[1]
					case 14: //-y -z x
						newAlignedBeacon[0] = -beacon[1]
						newAlignedBeacon[1] = -beacon[2]
						newAlignedBeacon[2] = beacon[0]
					case 15: //-y -x -z
						newAlignedBeacon[0] = -beacon[1]
						newAlignedBeacon[1] = -beacon[0]
						newAlignedBeacon[2] = -beacon[2]
					case 16: //-y z -x
						newAlignedBeacon[0] = -beacon[1]
						newAlignedBeacon[1] = beacon[2]
						newAlignedBeacon[2] = -beacon[0]
					case 17: //z -y x
						newAlignedBeacon[0] = beacon[2]
						newAlignedBeacon[1] = -beacon[1]
						newAlignedBeacon[2] = beacon[0]
					case 18: //-z -y -x
						newAlignedBeacon[0] = -beacon[2]
						newAlignedBeacon[1] = -beacon[1]
						newAlignedBeacon[2] = -beacon[0]
					case 19: //-x -z -y
						newAlignedBeacon[0] = -beacon[0]
						newAlignedBeacon[1] = -beacon[2]
						newAlignedBeacon[2] = -beacon[1]
					case 20: //-x z y
						newAlignedBeacon[0] = -beacon[0]
						newAlignedBeacon[1] = beacon[2]
						newAlignedBeacon[2] = beacon[1]
					case 21: //z -x -y
						newAlignedBeacon[0] = beacon[2]
						newAlignedBeacon[1] = -beacon[0]
						newAlignedBeacon[2] = -beacon[1]
					case 22: //-z -x y
						newAlignedBeacon[0] = -beacon[2]
						newAlignedBeacon[1] = -beacon[0]
						newAlignedBeacon[2] = beacon[1]
					case 23: //y -z -x
						newAlignedBeacon[0] = beacon[1]
						newAlignedBeacon[1] = -beacon[2]
						newAlignedBeacon[2] = -beacon[0]
					case 24: //y z x
						newAlignedBeacon[0] = beacon[1]
						newAlignedBeacon[1] = beacon[2]
						newAlignedBeacon[2] = beacon[0]
					}
					alignedScanner = append(alignedScanner, newAlignedBeacon)
				}

				for _, beacon := range finalBeacons {
					for _, b1 := range alignedScanner {
						scannerPosX := beacon[0] - b1[0]
						scannerPosY := beacon[1] - b1[1]
						scannerPosZ := beacon[2] - b1[2]

						correctedScanner := [][3]int{}
						for _, b2 := range alignedScanner {
							newCorrectedBeacon := [3]int{}
							newCorrectedBeacon[0] = b2[0] + scannerPosX
							newCorrectedBeacon[1] = b2[1] + scannerPosY
							newCorrectedBeacon[2] = b2[2] + scannerPosZ
							correctedScanner = append(correctedScanner, newCorrectedBeacon)
						}

						commonBeacons := 0
						for _, b2 := range correctedScanner {
							foundBeacon := false
							for _, beacon2 := range finalBeacons {
								if beacon2[0] == b2[0] && beacon2[1] == b2[1] && beacon2[2] == b2[2] {
									foundBeacon = true
									break
								}
							}
							if foundBeacon {
								commonBeacons++
							} else if b2[0] <= 1000 && b2[0] >= -1000 && b2[1] <= 1000 && b2[1] >= -1000 && b2[2] <= 1000 && b2[2] >= -1000 {
								break
							}
						}

						if commonBeacons >= 12 {
							output = append(output, fmt.Sprintln("Found Match"))
							overlap = true
							for _, b2 := range correctedScanner {
								foundBeacon := false
								for _, beacon2 := range finalBeacons {
									if beacon2[0] == b2[0] && beacon2[1] == b2[1] && beacon2[2] == b2[2] {
										foundBeacon = true
										break
									}
								}
								if !foundBeacon {
									finalBeacons = append(finalBeacons, [3]int{b2[0], b2[1], b2[2]})
								}
							}
							scannersPos = append(scannersPos, [3]int{scannerPosX, scannerPosY, scannerPosZ})
							break
						}
					}
					if overlap {
						break
					}
				}
				if overlap {
					break
				}
			}
			if overlap {
				scanners = append(scanners[:i], scanners[i+1:]...)
				i--
			}
		}
	}

	for _, b := range finalBeacons {
		output = append(output, fmt.Sprintln(b))
	}

	maxDistance := 0
	for _, b1 := range scannersPos {
		for _, b2 := range scannersPos {
			maxDistance = max(absolute(b2[0]-b1[0])+absolute(b2[1]-b1[1])+absolute(b2[2]-b1[2]), maxDistance)
		}
	}
	output = append(output, fmt.Sprintln("Most distant scanners are", maxDistance, "apart"))

	return output
}
