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
	var first, last string

	for i, c := range value {
		if unicode.IsDigit(c) {
			if len(first) == 0 {
				first = string(c)
			}

			last = string(c)
		} else {
			digit, err := getDigit(value, i)

			if err != nil {
				continue
			}

			if len(first) == 0 {
				first = digit
			}

			last = digit
		}
	}

	return first + last
}

func getDigit(value string, startIndex int) (string, error) {
	letters := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for endIndex := startIndex; endIndex <= len(value); endIndex++ {
		sliced := value[startIndex:endIndex]

		for key := range letters {
			if sliced == key {
				return letters[key], nil
			}
		}
	}

	return "", fmt.Errorf("No digit found")
}
