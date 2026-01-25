package main

import "fmt"

type User struct {
	Name    string
	Raiting float64
	Premium bool
}

func main() {
	weather := map[int]int{
		11: +3,
		12: +6,
		13: +9,
		14: -4,
		15: +1,
	}
	for k, v := range weather {
		fmt.Println(k, v)
	}
	weather[20] = 32
	fmt.Println(weather[20])
}
