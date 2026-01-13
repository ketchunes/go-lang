package main

import "fmt"

func main() {
	fmt.Println("Функция main")
	defer func() {
		fmt.Println("Анонимная функция main")
	}()
	hello()
}
func hello() {
	fmt.Println("Функция hello")
	defer func() {
		fmt.Println("defer 1")
	}()
	defer func() {
		fmt.Println("defer 2")
	}()
	defer func() {
		fmt.Println("defer 3")
	}()
	fmt.Println("lol")
}
