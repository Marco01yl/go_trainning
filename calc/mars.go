package main

import "fmt"

func main() {
	// fmt.Print("My")
	// fmt.Printf("weight is %v lbs.", 122.0*0.3783)
	// fmt.Printf("It would be %v years old\n", 25*365/687)
	// fmt.Printf("My weight on the surface of %v is %v lbs.\n ", "earth", 122.0)

	// const a = 10000
	// var b = 100

	// fmt.Printf("%-15v $%4v\n", "SpaceX", a/b)

	// b = 90
	// fmt.Printf("%-15v $%4v\n", "Virgin Galactic", a/b)

	const lightspeed = 299792
	var distance = 56000000

	fmt.Printf("it is %20v", distance/lightspeed, "seconds\n")

	distance = 401000000
	fmt.Println(distance/lightspeed, "seconds")

}
