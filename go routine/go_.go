package main

import (
	"fmt"
	"time"
)

func main() {
	go sleepygopher()
	time.Sleep(4 * time.Second)
	fmt.Println("wait ande meet snore")
}

func sleepygopher() {
	time.Sleep(3 * time.Second)
	fmt.Println("...snore...")
}
