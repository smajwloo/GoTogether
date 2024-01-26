package main

import "log"

func main() {
	input := getInput()
	getSymbols(input)

	println("Symbols found:", len(symbols))
	log.Println("Symbols:", symbols)

	for _, symbol := range symbols {
		findAdjacent(input, symbol)
		break
	}
}
