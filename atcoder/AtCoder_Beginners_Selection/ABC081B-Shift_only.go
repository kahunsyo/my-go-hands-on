package main

import (
	"fmt"
)

func main() {
	var N int
	var a int
	var result int = -1

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&a)
		tmp := how_many_2_is_contain(a)
		//fmt.Println(tmp)
		if result > tmp || result == -1 {
			result = tmp
		}
	}

	fmt.Println(result)
}

func how_many_2_is_contain(num int) int {
	var result int = 0

	for num > 0 {
		if num%2 == 0 {
			result += 1
			num /= 2
		} else {
			break
		}
	}

	return result
}
