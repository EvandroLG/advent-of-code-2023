package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input, err := os.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	result := calculateCalibration(string(input))
	fmt.Printf("Result: %d", result)
}

func calculateCalibration(input string) int {
	fragments := strings.Split(input, "\n")
	sum := 0

	for _, value := range fragments {
		calibration, _ := strconv.Atoi(findCalibrationValue(value))
		sum += calibration
	}

	return sum
}

func findCalibrationValue(value string) string {
	first := ""
	last := ""

	for _, c := range value {
		if unicode.IsDigit(c) {
			if first == "" {
				first = string(c)
			}

			last = string(c)
		}
	}

	return first + last
}
