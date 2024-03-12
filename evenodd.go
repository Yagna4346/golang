package main

import "fmt"

func main() {
	fmt.Println("Even Numbers")
	for i := range 20 {
		if i%2 == 0 {
			fmt.Println(i)
		}

	}
	fmt.Println("Odd Number")
	for i := range 20 {
		if i%2 == 1 {
			fmt.Println(i)
		}
vbn
	}
}
