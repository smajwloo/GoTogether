package main

import (
	"errors"
	"io/fs"
	"os"
)

func readFile() []byte {
	fsys := os.DirFS("resources")

	if calibrationsArray, err := fs.ReadFile(fsys, "input.txt"); err != nil {
		_ = errors.New("error reading input.txt")
		return nil
	} else {
		return calibrationsArray
	}
}
