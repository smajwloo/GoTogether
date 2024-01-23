package main

import (
	"log"
	"strconv"
	"strings"
)

var colors = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func checkGamePossibility(rounds []string) bool {
	for _, round := range rounds {
		if !checkRoundPossibility(round) {
			return false
		}
	}
	return true
}

func checkRoundPossibility(round string) bool {
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
		}

	}
	return true
}
