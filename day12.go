package main

import (
	"fmt"
	"strings"
)

func advent121(input []string) []string {
	output := []string{}
	var numberOfPaths int = 0

	var exploreCave func(connections [][]string, currentPath []string) [][]string
	exploreCave = func(connections [][]string, currentPath []string) [][]string {
		res := [][]string{}
		for i := 0; i < len(connections); i++ {
			if connections[i][0] == currentPath[len(currentPath)-1] {
				if connections[i][1] == "end" {
					res = append(res, append(currentPath, "end"))
				} else {
					validWay := true
					if connections[i][1][0] >= 'a' && connections[i][1][0] <= 'z' {
						for j := 0; j < len(currentPath); j++ {
							if connections[i][1] == currentPath[j] {
								validWay = false
								break
							}
						}
					}
					if validWay {
						newPaths := exploreCave(connections, append(currentPath, connections[i][1]))
						for j := 0; j < len(newPaths); j++ {
							res = append(res, []string{})
							for k := 0; k < len(newPaths[j]); k++ {
								res[len(res)-1] = append(res[len(res)-1], newPaths[j][k])
							}
						}
					}
				}
			}
		}
		return res
	}

	connections := [][]string{}
	for _, inputLine := range input {
		newConnection := strings.Split(inputLine, "-")
		if newConnection[0] != "start" && newConnection[1] != "start" && newConnection[0] != "end" && newConnection[1] != "end" {
			connections = append(connections, newConnection, []string{newConnection[1], newConnection[0]})
		} else if newConnection[1] == "start" || newConnection[0] == "end" {
			connections = append(connections, []string{newConnection[1], newConnection[0]})
		} else {
			connections = append(connections, newConnection)
		}
	}

	paths := exploreCave(connections, []string{"start"})
	for j := 0; j < len(paths); j++ {
		output = append(output, fmt.Sprintln("Path:", paths[j]))
	}

	numberOfPaths = len(paths)
	output = append(output, fmt.Sprintln("Number of path is", numberOfPaths))

	return output
}

func advent122(input []string) []string {
	output := []string{}
	var numberOfPaths int = 0

	var exploreCave func(connections [][]string, currentPath []string) [][]string
	exploreCave = func(connections [][]string, currentPath []string) [][]string {
		res := [][]string{}
		alreadyDuplicate := false
		for i := 0; i < len(currentPath); i++ {
			for j := 0; j < len(currentPath); j++ {
				if i != j && currentPath[i] == currentPath[j] && currentPath[i][0] >= 'a' && currentPath[i][0] <= 'z' {
					alreadyDuplicate = true
					break
				}
			}
		}
		for i := 0; i < len(connections); i++ {
			if connections[i][0] == currentPath[len(currentPath)-1] {
				if connections[i][1] == "end" {
					res = append(res, append(currentPath, "end"))
				} else {
					validWay := true
					if connections[i][1][0] >= 'a' && connections[i][1][0] <= 'z' {
						for j := 0; j < len(currentPath); j++ {
							if connections[i][1] == currentPath[j] && alreadyDuplicate {
								validWay = false
								break
							}
						}
					}
					if validWay {
						newPaths := exploreCave(connections, append(currentPath, connections[i][1]))
						for j := 0; j < len(newPaths); j++ {
							res = append(res, []string{})
							for k := 0; k < len(newPaths[j]); k++ {
								res[len(res)-1] = append(res[len(res)-1], newPaths[j][k])
							}
						}
					}
				}
			}
		}
		return res
	}

	connections := [][]string{}
	for _, inputLine := range input {
		newConnection := strings.Split(inputLine, "-")
		if newConnection[0] != "start" && newConnection[1] != "start" && newConnection[0] != "end" && newConnection[1] != "end" {
			connections = append(connections, newConnection, []string{newConnection[1], newConnection[0]})
		} else if newConnection[1] == "start" || newConnection[0] == "end" {
			connections = append(connections, []string{newConnection[1], newConnection[0]})
		} else {
			connections = append(connections, newConnection)
		}
	}

	paths := exploreCave(connections, []string{"start"})
	for j := 0; j < len(paths); j++ {
		output = append(output, fmt.Sprintln("Path:", paths[j]))
	}

	numberOfPaths = len(paths)
	output = append(output, fmt.Sprintln("Number of path is", numberOfPaths))

	return output
}
