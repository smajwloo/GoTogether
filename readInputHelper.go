package main

import (
	"strconv"
	"strings"
)

func getGameID(line string) (int, error) {
	parts := strings.Split(line, " ")[1]
	idStr := strings.Split(parts, ":")[0]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func getRounds(line string) []string {
	roundStr := strings.Split(line, ":")[1]
	rounds := strings.Split(roundStr, ";")

	for i := 0; i < len(rounds); i++ {
		rounds[i] = strings.TrimSpace(rounds[i])
	}
	return rounds
}
