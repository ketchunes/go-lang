package main

import "fmt"

type User struct {
	Name    string
	Age     int
	Number  string
	IsClose bool
	Raiting float64
}
type User1 struct {
	Name    string
	Raiting float64
}

func Greeting(u User1) {
	fmt.Println("Всем привет")

	fmt.Println("Меня зовут:", u.Name)
}
func main() {
	// user := User{
	// 	Name:    "Катя",
	// 	Age:     21,
	// 	Number:  "+7 854 214 32 33",
	// 	IsClose: true,
	// 	Raiting: 54.32,
	// }
	// fmt.Println(user)
	// fmt.Println("Имя:", user.Name)
	// user.Name = "Игорь"
	// fmt.Println("имя после:", user.Name)
}
