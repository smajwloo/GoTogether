package main

import "log"

func main() {
	input := readFile()
	calculateCalibration(input)

	log.Println(total)
}
