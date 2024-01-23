package main

import (
	"errors"
	"io/fs"
	"os"
	"strings"
)

func getInput() []string {
	fsys := os.DirFS("resources")

	if lines, err := fs.ReadFile(fsys, "input.txt"); err != nil {
		_ = errors.New("error reading input.txt")
		return nil
	} else {
		return strings.Split(string(lines), "\n")
	}
}
