package main

import (
	"fmt"
	"strings"
)

//	var t interface {
//		talk() string
//	}
type talker interface {
	talk() string
}
type martain struct{}

func (m martain) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}
func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type starship struct {
	laser
}

// func main() {
// 	s := starship{laser(3)} //{laser:3}
// 	fmt.Println(s.talk())
// 	fmt.Println(s.laser.talk())
// 	shout(s)

// 	// shout(martain{})
// 	// shout(laser(4))
// 	// t = martain{}
// 	// fmt.Println(t.talk())
// 	// t = laser(3)
// 	// fmt.Println(t.talk())

// }
