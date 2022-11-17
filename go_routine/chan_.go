package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	c := make(chan int)
// 	fmt.Println("pause")
// 	for i := 0; i < 5; i++ {
// 		go sleepygopher(i, c)
// 	}

// 	for i := 0; i < 5; i++ {
// 		gopherID := <-c
// 		fmt.Println("gopher", gopherID, "has finished sleeping")
// 	}
// }
// func sleepygopher(id int, c chan int) {
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("...", id, "...snore...")
// 	c <- id
// }
