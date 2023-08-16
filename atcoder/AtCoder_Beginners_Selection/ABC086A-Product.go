package main

import "fmt"

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	prod := a * b

	if prod%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
