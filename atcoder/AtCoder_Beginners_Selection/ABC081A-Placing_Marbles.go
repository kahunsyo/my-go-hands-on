package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	var result int = 0

	fmt.Scan(&input)
	//fmt.Println(input)
	input_n, _ := strconv.Atoi(input)
	for input_n > 0 {
		result += input_n % 10
		input_n /= 10
	}

	fmt.Println(result)
}
