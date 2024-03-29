package main

import (
	"strconv"
	"strings"
)

func firstPart(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		fragments := strings.Split(line, ": ")
		sets := strings.Split(fragments[1], "; ")

		var isGameValid = true

		for _, set := range sets {
			cubes := strings.Split(set, ", ")

			for _, cube := range cubes {
				if isGamePossible(cube) == false {
					isGameValid = false
					continue
				}
			}
		}

		if isGameValid == true {
			var id = getId(fragments[0])
			sum += id
		}

	}

	return sum
}

func isGamePossible(cube string) bool {
	colors := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	fragments := strings.Split(cube, " ")
	color := fragments[1]
	total, err := strconv.Atoi(fragments[0])

	return err == nil && total <= colors[color]
}

func getId(value string) int {
	id, _ := strconv.Atoi(strings.Split(value, " ")[1])
	return id
}
