package main

import "fmt"

func main() {
	fmt.Println("welcome to calculator app")

	var a, b int

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)

	fmt.Println(add(a, b))
	fmt.Println(substract(a, b))
	fmt.Println(multiply(a, b))
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

func mod(a, b int) int {
	return a % b
}
