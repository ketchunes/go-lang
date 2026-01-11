package main

import "fmt"

func main() {
	var lol int8 = 15
	if lol >= 10 && lol < 15 {
		fmt.Println("ты красавчик")
	} else if lol >= 15 {
		fmt.Println("Ты мега красавчик")
	} else {
		fmt.Println("Попробуй еще")
	}
}
