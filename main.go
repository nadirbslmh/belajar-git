package main

import "fmt"

func main() {
	fmt.Println("welcome to calculator app")

	fmt.Println(add(3, 4))
	fmt.Println(multiply(2, 3))
}

func add(a, b int) int {
	return a + b
}

func substract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func div(a, b int) int {
	if b == 0 {
		return 0
	}

	return a / b
}
