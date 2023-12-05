package day01

import (
	"log"
	"strconv"
	"strings"
)

func stringContainsComplexDigit(input string) bool {
	for _, complexString := range complexStrings {
		if strings.Contains(input, complexString) {
			return true
		}

	}
	return false
}
func isDigit(input uint8) bool {
	return input >= '0' && input <= '9'
}
func SolvePartOne(input string) int {
	sum := 0
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		log.Printf("Searching for digits in '%s' ", line)
		leftChar := ""
		rightChar := ""
		left := 0
		right := len(line) - 1

		for left <= right {
			if leftChar == "" && isDigit(line[left]) {
				leftChar = string(line[left])
			}
			if rightChar == "" && isDigit(line[right]) {
				rightChar = string(line[right])
			}
			if leftChar != "" && rightChar != "" {
				break
			}
			if leftChar == "" {
				left++
			}
			if rightChar == "" {
				right--
			}
		}

		combinedChars := ""
		if len(leftChar) == 0 {
			combinedChars = rightChar + rightChar
		} else if len(rightChar) == 0 {
			combinedChars = leftChar + leftChar
		} else {
			combinedChars = leftChar + rightChar
		}

		result, err := strconv.Atoi(combinedChars)
		if err != nil {
			log.Printf("Unable to parse found strings to int: %v", err)
			continue
		}
		log.Printf("Found: %d ", result)
		sum += result
	}

	return sum
}

func SolvePartTwo(input string) int {
	stringNumsToInt := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	convertedInput := ""
	for _, line := range input {
		// grab sections of 3 to five digits
		// check if in map
		// convert num to int
		// convert line
		// feed into convertedInput
		// feed convertedInput into part one
	}
}
