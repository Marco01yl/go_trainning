package main

// stardate returns a fiction measure of time for given date.
// // func stardate(t time.Time) float64 {
func stardate(t stardater) float64 {
	doy := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + doy + h
}

type soy int

func (s soy) YearDay() int {
	return int(s % 668)
}
func (s soy) Hour() int {
	return 0
}

type stardater interface {
	YearDay() int
	Hour() int
}

// func main() {
// 	day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)
// 	fmt.Printf("%.1f Curiosity has landed \n", stardate(day))

// 	s := soy(1422)
// 	fmt.Printf("%.1f Happy birthday!\n", stardate(s))
// }
