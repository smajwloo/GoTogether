package main

import (
	"log"
	"strconv"
	"strings"
)

var total int

func calculatePowerOfGame(rounds []string) {
	var colors = map[string]int{
		"red":   -1,
		"green": -1,
		"blue":  -1,
	}
	var power = 0
	for _, round := range rounds {
		calculateMinimalValue(round, colors)
	}
	power = calculatePowerOfCubes(colors, power)
	log.Println("Power", power)
	total += power
}

func calculateMinimalValue(round string, colors map[string]int) {
	for color, lowestValue := range colors {
		if strings.Contains(round, color) {
			first := strings.Split(round, " "+color)[0]
			second := strings.Split(first, " ")

			number := second[len(second)-1]
			amount, err := strconv.Atoi(number)

			if err != nil {
				log.Println("Error converting red amount to int", err)
				continue
			}

			if amount > lowestValue {
				colors[color] = amount
			}
		}
	}
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
