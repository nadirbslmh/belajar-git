package main

import "fmt"

func main() {
	fmt.Println("welcome to calculator app")

	fmt.Println(add(3, 4))
}

func add(a, b int) int {
	return a + b
}

func substract(a, b int) int {
	return a - b
}
