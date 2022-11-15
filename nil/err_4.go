package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Grid [rows][columes]int8

const rows, columes = 9, 9

func validDigit(digit int8) bool {
	if digit > 9 {
		return false
	}
	return true
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

var (
	ErrBounds = errors.New("out of bounds!")
	ErrDigit  = errors.New("invalid digit")
)

type SudokuError []error

func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}

func (g *Grid) Set(row, colume int, digit int8) error {
	var errs SudokuError
	if !inBounds(row, colume) {
		errs = append(errs, ErrBounds) //return ErrBounds//errors.New("out of bounds")
	}
	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}
	if len(errs) > 0 {
		return errs
	}

	g[row][colume] = digit
	return nil
}
func main() {
	var g Grid
	err := g.Set(12, 0, 15)

	if err != nil {
		if errs, ok := err.(SudokuError); ok {
			fmt.Printf("%d Error(s) ocured:\n", len(errs))
			for _, e := range errs {
				fmt.Printf("- %v\n", e)
			}
		}
		os.Exit(1)
	}
	// if err != nil {
	// 	switch err {
	// 	case ErrBounds, ErrDigit:
	// 		fmt.Println("Les errurs de parameters ")
	// 	default:
	// 		fmt.Println(err)
	// 	}
	// 	os.Exit(1)

	// fmt.Printf("An error ocurred: %v.\n", err)
	// os.Exit(1)

}
