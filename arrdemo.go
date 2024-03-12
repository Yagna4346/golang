package main

import "fmt"

func main() {
	array1 := [5]int{}
	for i := 0; i < 5; i++ {
		fmt.Scanln(&array1[i])
	}
	for i := 0; i < 5; i++ {
		fmt.Println(array1[i])
	}
}
