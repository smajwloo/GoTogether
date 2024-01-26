package main

import "strconv"

func findAdjacent(lines []string, symbol Symbol) {
	startLineIndex, endLineIndex, startColIndex, endColIndex := setIndexes(lines, symbol)
	for i := startLineIndex; i <= endLineIndex; i++ {
		for x := startColIndex; x <= endColIndex; x++ {
			if i == symbol.Line && x == symbol.Col {
				continue
			}

			combinedNumber := ""
			number := getNumber(string(lines[i][x]))
			if number != -1 {
				index := x
				for getNumber(string(lines[i][index-1])) != -1 {
					combinedNumber += string(rune(number + getNumber(string(lines[i][x-1]))))
					index--
				}
				index = x
				for getNumber(string(lines[i][index+1])) != -1 {
					combinedNumber += string(rune(number + getNumber(string(lines[i][x-1]))))
					index++
				}
				println("combinedNumber: ", combinedNumber)
			}
		}
	}
}

func getNumber(char string) int {
	number, err := strconv.Atoi(char)
	if err != nil {
		return -1
	}
	return number
}

func setIndexes(lines []string, symbol Symbol) (int, int, int, int) {
	startLineIndex := setIndex(lines, symbol.Line, true)
	endLineIndex := setIndex(lines, symbol.Line, false)
	startColIndex := setIndex(lines, symbol.Col, true)
	endColIndex := setIndex(lines, symbol.Col, false)
	return startLineIndex, endLineIndex, startColIndex, endColIndex
}

func setIndex(lines []string, index int, isStart bool) int {
	if index == 0 || index == len(lines) {
		return index
	} else if isStart {
		return index - 1
	}
	return index + 1
}
