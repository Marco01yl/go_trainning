package main

import (
	// "math/big"
	// "strings"
	// "time"
	// "math/rand"
	// "unicode/utf8"
	// "math"
	"fmt"
	// "strings.Replace"
)

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

	// const lightspeed = 299792
	// var distance = 56000000

	// fmt.Printf("it is %20v %v ", distance/lightspeed, "seconds\n")

	// distance = 401000000
	// fmt.Println(distance/lightspeed, "seconds")

	// var weight = 122.0
	// weight = weight * 0.3783
	// weight *= 2
	// fmt.Printf("%v", weight)

	// var num = rand.Intn(10) + 1
	// fmt.Println(num)

	// num = rand.Intn(10) + 1
	// fmt.Println(num)

	// fmt.Println("111111111111")
	// var c = "walk outside"
	// var exit = strings.Contains(c, "outside")
	// fmt.Println("result is:", exit)

	// var room = "lake"

	// switch {
	// case room == "cave":
	// 	fmt.Printf("dimly")
	// case room == "lake":
	// 	fmt.Printf("solid\n")
	// 	fallthrough
	// case room == "underwater":
	// 	fmt.Printf("freezing!")
	// }

	// var count = 5
	// for {
	// 	if count == 0 {
	// 		fmt.Println("off!")
	// 		break
	// 	}
	// 	fmt.Println(count)
	// 	time.Sleep(time.Second) //等待1s，2s就*2
	// 	count--

	// }
	// count = 5
	// for count > 0 {
	// 	fmt.Println(count)
	// 	time.Sleep(time.Second) //等待1s，2s就*2
	// 	count--
	// }
	// fmt.Println("off!")

	// const target = 16//猜数
	// var guess = rand.Intn(21)
	// var times = 1
	// for guess != target {
	// 	if guess > target {
	// 		fmt.Printf("%v is too big!\n", guess)
	// 	}
	// 	if guess < target {
	// 		fmt.Printf("%v is too small!\n", guess)
	// 	}
	// 	guess = rand.Intn(21)
	// 	times++
	// }
	// fmt.Printf("%v is right, we used %v times!", guess, times)

	// f1  := 10.0 / 3
	// fmt.Printf("itis %4.2f", f1)

	// year := 2018
	// fmt.Printf("Type %T for %v\n", year, year)
	// a := "text"
	// fmt.Printf("Type %T for %[1]v\n", a)
	// b := 42
	// fmt.Printf("Type %T for %[1]v\n", b)
	// c := 3.14
	// fmt.Printf("Type %T for %[1]v\n", c)
	// d := true
	// fmt.Printf("Type %T for %[1]v\n", d)

	// future := time.Unix(12622780800, 0)
	// fmt.Println(future)

	// distance := new(big.Int)
	// distance.SetString("240000000000000000000000", 10)
	// fmt.Println(distance)

	// var shalom string = "shlalom"
	// fmt.Printf("%c\n%c\n%c\n%c\n%c\n%c\n%c\n", shalom[0], shalom[1], shalom[2], shalom[3], shalom[4], shalom[5], shalom[6])

	// question := "ques tion。"
	// q, size := utf8.DecodeRuneInString(question) //按顺序返回两个值
	// fmt.Printf("rune:%c %v .", q, size)

	// v := 42
	// if v >= 0 && v <= math.MaxUint8 {
	// 	v8 := uint8(v)
	// 	fmt.Println("converted:", v8)
	// }

	face_ := "CSOITEUIWUIZNSROCNKFD"
	len := len(face_)
	key_ := "GOLANG"

	ans := ""

	for i := 0; i < len; i++ {
		k := i
		for k > 5 {
			k -= 6

		}

		cell := int(face_[i])
		step := int(key_[k]) - 65
		ans_ := cell + step
		if ans_ > 90 {
			ans_ -= 26
		}
		ans_cell := string(ans_)
		ans = ans + ans_cell
	}
	fmt.Println(ans)

}
