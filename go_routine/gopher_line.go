package main

import (
	"fmt"
	"strings"
)

func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye"} {
		downstream <- v
	}
	// downstream <- " " //哨兵值
	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	for {
		item, ok := <-upstream
		// if item == " " {
		// 	downstream <- " "
		// 	return
		// }
		if !ok {
			close(downstream)
			return
		}
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
}

func printGopher(upstream chan string) {
	for {
		v, ok := <-upstream
		if !ok {
			return
		}
		// if v == " " {
		// 	return
		// }
		fmt.Println(v)
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0, c1)
	printGopher(c1)

}
