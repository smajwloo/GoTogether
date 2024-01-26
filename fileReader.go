package main

import (
	"errors"
	"io/fs"
	"os"
	"strings"
)

var symbols = make([]Symbol, 0)

func getInput() []string {
	fsys := os.DirFS("resources")

	if lines, err := fs.ReadFile(fsys, "input.txt"); err != nil {
		_ = errors.New("error reading input.txt")
		return nil
	} else {
		return strings.Split(string(lines), "\n")
	}
}

func getSymbols(lines []string) {
	for lineNumber, line := range lines {
		for colIndex, char := range line {
			if isSymbol(char) {
				symbols = append(symbols, Symbol{Char: string(char), Line: lineNumber, Col: colIndex})
			}
		}
	}
}

func isSymbol(char rune) bool {
	symbols := []rune{'%', '#', '$', '&', '@', '=', '-', '+', '/', '*'}
	for _, symbol := range symbols {
		if char == symbol {
			return true
		}
	}
	return false
}
