package main

import (
	"fmt"
)

type location struct {
	lat  float64
	long float64
	name string
}

func plus_long(a location) location {
	a.long += 1
	return a
}
func main() {
	// var curiosity struct {
	// 	lat  float64
	// 	long float64
	// 	high int
	// }

	// var spirit location
	// spirit.lat = 12.125
	// spirit.long = 88.552
	// var curiosity location
	// curiosity.lat = -4.125
	// curiosity.long = 123.564
	// fmt.Println(curiosity.lat, curiosity.long)
	// new_curio := plus_long(curiosity)
	// fmt.Println(curiosity, spirit)
	// fmt.Println(new_curio)
	lications := []location{
		{10.15, 100.15, "BL"}, {11.16, 101.16, "ACL"}, {12.17, 102.18, "BGFR"},
	}
	fmt.Println(lications)
	fmt.Printf("%+v", lications)

}
