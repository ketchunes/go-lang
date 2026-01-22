package main

import "fmt"

func main() {
	arr := [5]int{5, 66, 7, 100, 1}
	for i := 0; i < 5; i++ {
		if arr[i]%2 == 0 {
			arr[i] *= 2
		}
	}
	for i := 0; i < 5; i++ {
		fmt.Println(i, ":", arr[i])
	}
}
