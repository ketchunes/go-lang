package main

import "fmt"

func main() {
	number := 10
	pointer := &number
	kek(pointer)
	fmt.Println(pointer)
}
func foo(n *int) {
	fmt.Println(n)
	fmt.Println(*n)
}
func kek(a int) {
	a = 5
}
