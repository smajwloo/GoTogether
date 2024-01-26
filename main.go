package main

import "log"

func main() {
	input := getInput()

	for _, line := range input {
		gameId, err := getGameID(line)
		if err != nil {
			log.Println("Error getting game ID", err)
			return
		}
		log.Println(gameId)

		rounds := getRounds(line)
		calculatePowerOfGame(rounds)
	}

	log.Println("Total", total)
}
