package main

import (
	"fmt"
	"study/greeting"
)

func main() {
	fmt.Println("lol")
	greeting.SayHello()
	greeting.SayBad()

	i := greeting.GiveMeInt()
	fmt.Println("i:", i)
}
