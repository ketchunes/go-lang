package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите команду")
	// ok := scanner.Scan()
	// if !ok {
	// 	fmt.Println("Ошибка пользовательского ввода")
	// 	return
	// }
	if ok := scanner.Scan(); !ok {
		fmt.Println("Ошибка пользовательского ввода")
		return
	}
	text := scanner.Text()
	fields := strings.Fields(text)
	fmt.Println("text:", text)
	fmt.Println("Слова:", fields)
}
