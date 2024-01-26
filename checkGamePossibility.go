package main

import (
	"log"
	"strconv"
	"strings"
)

// TODO: Remove parts that come from part 1, these are not needed anymore
var colors = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func checkGamePossibility(rounds []string) bool {
	power := 0
	for _, round := range rounds {
		if !checkRoundPossibility(round, power) {
			return false
		}
	}
	return true
}

func checkRoundPossibility(round string, power int) bool {
	for color, maxValue := range colors {
		if strings.Contains(round, color) {
			first := strings.Split(round, " "+color)[0]
			second := strings.Split(first, " ")

			number := second[len(second)-1]
			amount, err := strconv.Atoi(number)

			if err != nil {
				log.Println("Error converting red amount to int", err)
				continue
			}
			if amount > maxValue {
				return false
			}

			currentAmounts := getLowerNumber(amount)
			power = calculatePowerOfCubes(currentAmounts, power)
			log.Println("Power", power)
		}

	}
	return true
}

func calculatePowerOfCubes(currentAmounts map[string]int, power int) int {
	for _, amount := range currentAmounts {
		if power == 0 {
			power = amount
		} else {
			power = power * amount
		}
	}
	return power
}

func getLowerNumber(currentAmount int) map[string]int {
	currentAmounts := colors
	for color, amount := range currentAmounts {
		if currentAmount < amount {
			currentAmounts[color] = currentAmount
			log.Println("Current amount", currentAmount)
		}
	}
	return currentAmounts
}
