package main

import (
	"fmt"
	"study/greeting"
	"study/user"
)

func main() {
	i := greeting.GiveMeInt()
	fmt.Println("i:", i)
	u := user.User{}
	fmt.Println(u)

}
