package main

type stats struct {
	level, endurance, health int
}

func levelup(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

type character struct {
	name  string
	stats stats
}

// func main() {
// 	player1 := character{name: "Matthias"}
// 	levelup(&player1.stats)
// 	fmt.Printf("%+v\n", player1)
// 	fmt.Printf("%+v\n", player1.stats)
// }
