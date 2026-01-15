package main

import "fmt"

func main() {
	number := 10
	pointer := &number
	foo(pointer)
}
func foo(n *int) {
	fmt.Println(n)
}
