package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines, shouldReturn := readAndParseFile()
	if shouldReturn {
		return
	}

	currentCalories := 0
	maxCalories := make([]int, 3)
	for _, input := range lines {
		cleanedInput := strings.TrimSpace(input)

		if cleanedInput == "" {
			SelectMaxValues(maxCalories, currentCalories)

			currentCalories = 0
			continue
		}

		calories, err := strconv.Atoi(cleanedInput)
		if err != nil {
			fmt.Println(fmt.Sprintf("found invalid input line: %v", cleanedInput))
		}

		currentCalories += calories
	}

	SelectMaxValues(maxCalories, currentCalories)

	fmt.Println(fmt.Sprintf("Top 3 Calories: %v", Sum(maxCalories)))
}

func readAndParseFile() ([]string, bool) {
	fmt.Println("Please enter the unix-style path to the test file:")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(fmt.Sprintf("Received error while reading input: %s", err.Error()))
		return nil, true
	}

	trimmedPath := strings.TrimSpace(text)
	if trimmedPath == "" {
		// Default to sample file
		trimmedPath = "./sample_files/input.txt"
	}

	bytes, err := os.ReadFile(trimmedPath)
	if err != nil {
		fmt.Println(fmt.Sprintf("Received error while reading file: %s", err.Error()))
		return nil, true
	}

	lines := strings.Split(string(bytes), "\n")
	return lines, false
}

// This will test each value in the current max array and if the candidate is larger than that one,
// it will replace it and use that value as the new candidate.  This will maintain the
// array in descending order.
func SelectMaxValues(currentMax []int, candidate int) {
	for offset, value := range currentMax {
		if value < candidate {
			currentMax[offset] = candidate
			// We reassign to candidate here to check subsequent elements in the array to verify
			// that none of them are less than the value we replaced.
			candidate = value
		}
	}
}

func Sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
