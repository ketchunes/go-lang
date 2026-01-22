package main

import "fmt"

type User struct {
	Name    string
	Raiting float64
	Premium bool
}

func main() {
	userArray := [4]User{
		User{
			Name:    "Данил",
			Raiting: 5.5,
			Premium: true,
		},
		User{
			Name:    "Рома",
			Raiting: 3.1,
			Premium: false,
		},
		User{
			Name:    "Федя",
			Raiting: 1.9,
			Premium: true,
		},
		User{
			Name:    "Катя",
			Raiting: 9.1,
			Premium: true,
		},
	}

	fmt.Println("До:")
	for i := 0; i < 4; i++ {
		fmt.Println(userArray[i])
	}
	fmt.Println("")

	for i := 0; i < 4; i++ {
		if userArray[i].Premium {
			userArray[i].Raiting += 1.0
		}
	}
	fmt.Println("После:")
	for i := 0; i < 4; i++ {
		fmt.Println(userArray[i])
	}

}
