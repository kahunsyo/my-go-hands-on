package main

import (
	"fmt"
)

func main() {
	var a, b, c, x int

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&x)

	var n_500 int = a
	var n_100 int = b
	var n_50 int = c
	var target int = x

	result_n_conbinations := how_many_combinations_exsist(n_500, n_100, n_50, target)
	fmt.Println(result_n_conbinations)
}

func how_many_combinations_exsist(n_500 int, n_100 int, n_50 int, target int) int {
	n_comb := how_many_combinations_exsist_500(n_500, n_100, n_50, target)
	return n_comb
}

func how_many_combinations_exsist_500(n_500 int, n_100 int, n_50 int, target int) int {
	var n_comb int
	//fmt.Println(n_500, n_100, n_50, target)
	if n_500 < 0 || n_100 < 0 || n_50 < 0 {
		return 0
	}

	sum := n_500*500 + n_100*100 + n_50*50

	if target == sum {
		return 1
	} else if target > sum {
		return 0
	} else {
		n_comb += how_many_combinations_exsist_500(n_500-1, n_100, n_50, target)
		n_comb += how_many_combinations_exsist_100(n_500, n_100-1, n_50, target)
		n_comb += how_many_combinations_exsist_50(n_500, n_100, n_50-1, target)

		return n_comb
	}

}

func how_many_combinations_exsist_100(n_500 int, n_100 int, n_50 int, target int) int {
	var n_comb int
	//fmt.Println(n_500, n_100, n_50, target)
	if n_500 < 0 || n_100 < 0 || n_50 < 0 {
		return 0
	}

	sum := n_500*500 + n_100*100 + n_50*50

	if target == sum {
		return 1
	} else if target > sum {
		return 0
	} else {
		//n_comb += how_many_cmbinations_exsist(n_500-1, n_100, n_50, target)
		n_comb += how_many_combinations_exsist_100(n_500, n_100-1, n_50, target)
		n_comb += how_many_combinations_exsist_50(n_500, n_100, n_50-1, target)

		return n_comb
	}

}
func how_many_combinations_exsist_50(n_500 int, n_100 int, n_50 int, target int) int {
	var n_comb int
	//fmt.Println(n_500, n_100, n_50, target)
	if n_500 < 0 || n_100 < 0 || n_50 < 0 {
		return 0
	}

	sum := n_500*500 + n_100*100 + n_50*50

	if target == sum {
		return 1
	} else if target > sum {
		return 0
	} else {
		//n_comb += how_many_combinations_exsist(n_500-1, n_100, n_50, target)
		//n_comb += how_many_combinations_exsist(n_500, n_100-1, n_50, target)
		n_comb += how_many_combinations_exsist_50(n_500, n_100, n_50-1, target)

		return n_comb
	}

}
