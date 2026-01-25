package main

import "fmt"

type User struct {
	Name    string
	Raiting float64
	Premium bool
}

func main() {
	intSlice := make([]int, 0)
	intSlice = append(intSlice, 10, 15, 20, 150)
	fmt.Println(intSlice)
}
