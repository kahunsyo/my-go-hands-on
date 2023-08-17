package main

import "fmt"

func main() {
	var n, a, b int
	var sum int = 0

	fmt.Scan(&n, &a, &b)

	for i := 1; i <= n; i++ {
		sum_of_digits_ := sum_of_digits(i)
		if a <= sum_of_digits_ && sum_of_digits_ <= b {
			sum += i
		}
	}

	fmt.Println(sum)
}

func sum_of_digits(x int) int {
	var sum int = 0
	for x > 0 {
		sum += x % 10
		x /= 10
	}

	return sum
}
