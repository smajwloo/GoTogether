package main

import (
	"errors"
	"io/fs"
	"os"
	"strconv"
	"strings"
)

var total int

func getInput() []byte {
	fsys := os.DirFS("resources")

	if calibrationsArray, err := fs.ReadFile(fsys, "input.txt"); err != nil {
		_ = errors.New("error reading input.txt")
		return nil
	} else {
		return calibrationsArray
	}
}

func calculateCalibration(calibrationsArray []byte) {
	calibrationsString := string(calibrationsArray)
	calibrations := strings.Split(calibrationsString, "\n")

	for _, calibration := range calibrations {
		numbers := extractNumbers(calibration)
		calculateTotal(numbers)
	}
}

func extractNumbers(calibration string) []int {
	numbers := make([]int, 0)

	for _, char := range calibration {
		number := getNumber(string(char))
		if number > -1 {
			numbers = append(numbers, number)
		}
	}

	return numbers
}

func calculateTotal(numbers []int) {
	if len(numbers) == 0 {
		return
	}

	combinedNumbers := ""
	if len(numbers) == 1 {
		combinedNumbers = combineNumbers(numbers[0], numbers[0])
	} else {
		combinedNumbers = combineNumbers(numbers[0], numbers[len(numbers)-1])
	}

	number := getNumber(combinedNumbers)
	if number > -1 {
		total += number
	}
}

func combineNumbers(firstNumber int, secondNumber int) string {
	return strings.Join([]string{strconv.Itoa(firstNumber), strconv.Itoa(secondNumber)}, "")
}

func getNumber(char string) int {
	if number, err := strconv.Atoi(char); err != nil {
		return -1
	} else {
		return number
	}
}
