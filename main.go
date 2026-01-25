package main

import "fmt"

type User struct {
	Name    string
	Raiting float64
	Premium bool
}

func main() {
	userArray := []User{
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
	fmt.Println("len", len(userArray))
	fmt.Println("cap", cap(userArray))
	userArray = append(userArray,
		User{
			Name:    "Ваня",
			Raiting: 2.2,
			Premium: true,
		},
	)
	fmt.Println("len", len(userArray))
	fmt.Println("cap", cap(userArray))
	// fmt.Println("До:")
	for index, user := range userArray {
		fmt.Println(index, user)
	}
	// // for i := 0; i < len(userArray); i++ {
	// // 	fmt.Println(userArray[i])
	// // }
	// fmt.Println("")

	// for index, user := range userArray {
	// 	if user.Premium {
	// 		userArray[index].Raiting += 1.0
	// 	}
	// }
	// // for i := 0; i < len(userArray); i++ {
	// // 	if userArray[i].Premium {
	// // 		userArray[i].Raiting += 1.0
	// // 	}
	// // }
	// fmt.Println("После:")
	// for _, user := range userArray {
	// 	fmt.Println(user)
	// }

}
