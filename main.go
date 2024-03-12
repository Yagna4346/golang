package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	fmt.Println(a + b)
	i := 1
	fmt.Printf("%t", a)
	//fmt.Printf("%t")
	switch i {
	case 1:
		fmt.Println("hello world")
	default:
		fmt.Println("exit")
	}
}
