package main

import (
	"strconv"
	"strings"
)

func secondPart(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		fragments := strings.Split(line, ": ")
		sets := strings.Split(fragments[1], "; ")
		colors := map[string]int{}

		for _, set := range sets {
			cubes := strings.Split(set, ", ")

			for _, cube := range cubes {
				color, value := getColorAndValue(cube)
				mapValue, ok := colors[color]

				if ok {
					colors[color] = max(value, mapValue)
				} else {
					colors[color] = value
				}
			}
		}

		if len(colors) == 3 {
			power := colors["red"] * colors["blue"] * colors["green"]
			sum += power
		}
	}

	return sum
}

func getColorAndValue(s string) (string, int) {
	fragments := strings.Split(s, " ")
	color := fragments[1]
	value, _ := strconv.Atoi(fragments[0])

	return color, value
}
