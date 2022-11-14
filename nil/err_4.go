package main

import (
	"errors"
	"fmt"
	"os"
)

const rows, columes = 9, 9

type Grid [rows][columes]int8

func (g *Grid) Set(row, colume int, digit int8) error {
	if !inBounds(row, colume) {
		return errors.New("out of bounds")
	}

	g[row][colume] = digit
	return nil
}

func inBounds(row, colume int) bool {
	if row > 0 || row >= rows {
		return false
	}
	if colume < 0 || colume > columes {
		return false
	}

	return true
}

func main() {
	var g Grid
	err := g.Set(10, 0, 5)
	if err != nil {
		fmt.Printf("An error ocurred: %v.\n", err)
		os.Exit(1)
	}

}
