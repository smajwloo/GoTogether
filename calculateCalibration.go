package main

import (
	"log"
	"strconv"
	"strings"
)

var writtenNumbers = map[string]string{
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
var total int

func calculateCalibration(calibrationsArray []byte) {
	calibrationsString := string(calibrationsArray)
	calibrations := strings.Split(calibrationsString, "\n")

	for _, calibration := range calibrations {
		log.Println(calibration)
		newCalibration := replaceWrittenNumbers(calibration)
		log.Println(newCalibration)
		numbers := extractNumbers(newCalibration)
		log.Println(numbers)
		calculateTotal(numbers)
	}
}

func replaceWrittenNumbers(calibration string) string {
	replacedCalibration := calibration

	for key, value := range writtenNumbers {
		if !strings.Contains(replacedCalibration, key) {
			continue
		}

		value = string(key[0]) + value + string(key[len(key)-1])
		replacedCalibration = strings.ReplaceAll(replacedCalibration, key, value)
	}
	return replacedCalibration
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
