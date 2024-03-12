package main

import (
	"fmt"
)

var a int = 10

func main() {
	fmt.Println("Enter a choice ")
	fmt.Println("1.Addition")
	fmt.Println("2.Subtarction")
	fmt.Println("3.Multiplication")
	fmt.Println("4.Division")
	var c int
	fmt.Scanln(&c)
	var a int
	var b int
	fmt.Println("Enter a value ")
	fmt.Scanln(&a)
	fmt.Println("Enter b value ")
	fmt.Scanln(&b)
	switch c {
	case 1:
		fmt.Println("Result: ", addition(a, b))
	case 2:
		fmt.Println("Result: ", subtarction(a, b))
	case 3:
		fmt.Println("Result: ", multiplication(a, b))
	case 4:
		fmt.Println("Result: ", divison(a, b))
	default:
		fmt.Println("Enter valid choice")
	}
}
func addition(a, b int) int {
	return a + b
}

func subtarction(a, b int) int {
	return a - b
}

func divison(a, b int) int {
	return a / b
}

func multiplication(a, b int) int {
	return a * b
}
