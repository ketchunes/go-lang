package main

import "fmt"

func main() {
	criminal := map[string]bool{
		"Вася":    true,
		"Игорь":   false,
		"Валерий": false,
		"Настя":   true,
		"Ваня":    true,
	}
	c, lol := criminal["ы"]
	if !lol {
		fmt.Println("Человека нет в базе")
	}
	fmt.Println(c, lol)
}
