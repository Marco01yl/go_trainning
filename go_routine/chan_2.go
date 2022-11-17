package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	fmt.Println("pause")
	for i := 0; i < 5; i++ {
		go sleepygopher(i, c)
	}

	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		case gopherID := <-c:
			fmt.Println("gopher", gopherID, "has finished slept")
		case <-timeout:
			fmt.Println("bye!")
			return
		}
	}
}
func sleepygopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(3500)) * time.Millisecond)
	c <- id
}
