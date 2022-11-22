package main

import (
	"image"
	"log"
	"time"
)

// func worker() {
// 	pos := image.Point{X:10, Y: 10}
// 	direction := image.Point{X: 1, Y: 0}
// 	next := time.After(time.Second)
// 	for {
// 		select {
// 		case <-next:
// 			pos = pos.Add(direction)
// 			fmt.Println("current position is", pos)
// 			next = time.After(time.Second)
// 		}
// 	}
// }

type command int

const (
	right = command(0)
	left  = command(1)
)

type RoverDriver struct {
	commandCh chan command
}

func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandCh: make(chan command),
	}
	go r.drive()
	return r
}

func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-r.commandCh:
			switch c {
			case right:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			case left:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			}
			log.Printf("new direction %v", direction)
		case <-nextMove:
			pos = pos.Add(direction)
			log.Printf("moved to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

func (r *RoverDriver) Left() {
	r.commandCh <- left
}
func (r *RoverDriver) Right() {
	r.commandCh <- right
}

func main() {
	r := NewRoverDriver()
	time.Sleep(2 * time.Second)
	r.Left()
	time.Sleep(2 * time.Second)
	r.Right()
	time.Sleep(2 * time.Second)

}
