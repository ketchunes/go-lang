package main

import "fmt"

func main() {
	fmt.Println("Вызываю square:")
	square(4)
	square(12)
	square(10)
	fmt.Println("Закончил выполнение")
	zakaz("danil", "spagetti")
}
func square(x int) {
	fmt.Println("Мы присвоили в функцию переменную x:", x)
	fmt.Println("x в квадрате:", x*x)
}
func zakaz(name, lol string) {
	fmt.Println("Накрываем стол")
	fmt.Println("Привет:", name)
	fmt.Println("Ваш заказ:", lol)
	fmt.Println("Принести:", lol)
}
