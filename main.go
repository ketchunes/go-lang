package main

import (
	"fmt"
)

// type User struct {
// 	Name    string
// 	Age     int
// 	Number  string
// 	IsClose bool
// 	Raiting float64
// }
// type User1 struct {
// 	Name    string
// 	Raiting float64
// }

// func (u User1) Greeting() {
// 	fmt.Println("Всем привет")
// 	fmt.Println("Меня зовут:", u.Name)
// 	fmt.Println("Мой рейтинг:", u.Raiting)
// }
// func (u User1) Goodbye() {

// 	fmt.Println("Меня зовут", u.Name)
// 	fmt.Println("Всем пока")
// 	fmt.Println("Мой рейтинг щас", u.Raiting)
// }

type Product struct {
	ImgURL      string
	Name        string
	Description string
	Price       int
}

func main() {
	// user := User1{
	// 	Name:    "Максим",
	// 	Raiting: 6.0,
	// }
	// user.Greeting()
	// user.Goodbye()

	smartphone := Product{
		ImgURL:      "https://imgsite/image.png",
		Name:        "Samsung s92000",
		Description: "Bomba telephone",
		Price:       100000000,
	}

	fmt.Println("Image of phone: ", smartphone.ImgURL)
	fmt.Println("Name of phone: ", smartphone.Name)
	fmt.Println("Description of phone: ", smartphone.Description)
	fmt.Println("Price of phone: ", smartphone.Price)

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
