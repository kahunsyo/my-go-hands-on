package main

import (
	"fmt"
	"sort"
)

func main() {
	var sum_Alice int = 0
	var sum_Bob int = 0

	var n int
	fmt.Scan(&n)

	var a []int
	a = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	for i, v := range a {
		if i%2 == 0 {
			sum_Alice += v
		} else {
			sum_Bob += v
		}
	}

	fmt.Println(sum_Alice - sum_Bob)
}
