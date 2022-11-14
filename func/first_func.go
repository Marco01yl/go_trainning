package main

import (
	"fmt"
)

var blank_table = func() {
	fmt.Printf("%v\n|%-10v|%-10v|\n%v\n", "=======================", "°C", "°f", "=======================")
}

func write_table(max float64, min float64) {
	for max >= min {
		farehen := ((min * 9.0) / 5.0) + 32.0
		fmt.Printf("|%-10.2f|%-10.2f|\n", min, farehen)
		min += 5.0
	}
}

func main() {
	blank_table()
	write := write_table
	write(100.0, -40.0)

}
