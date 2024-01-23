package main

import "log"

func main() {
	input := getInput()
	total := 0

	for _, line := range input {
		gameId, err := getGameID(line)
		if err != nil {
			log.Println("Error getting game ID", err)
			return
		}
		log.Println(gameId)

		rounds := getRounds(line)
		isGamePossible := checkGamePossibility(rounds)

		if isGamePossible {
			total += gameId
		} else {
			log.Println("Game", gameId, "is not possible")
		}
	}

	log.Println("Total", total)
}
