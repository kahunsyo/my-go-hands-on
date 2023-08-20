package main

import "fmt"

func main() {
	var n, y int
	var result_n_10000 int = 0
	var result_n_5000 int = 0
	var result_n_1000 int = 0

	fmt.Scan(&n, &y)

	result_n_10000, result_n_5000, result_n_1000 = search_comb_of_bills(n, y)

	fmt.Println(result_n_10000, result_n_5000, result_n_1000)
}

func search_comb_of_bills(n int, y int) (int, int, int) {
	var n_10000 int = 0
	var n_5000 int = 0
	var n_1000 int = 0
	var is_exist bool = false

	for n_5000 = 0; n_5000 <= n; n_5000++ {
		n_10000, is_exist = cal_n_10000(n, y, n_5000)
		if is_exist == false {
			continue
		}

		n_1000, is_exist = cal_n_1000(n, y, n_5000)
		if is_exist == false {
			continue
		}

		return n_10000, n_5000, n_1000
	}

	//Not found
	return -1, -1, -1
}

func cal_n_10000(n int, y int, n_5000 int) (int, bool) {
	if (y-1000*n-4000*n_5000)%9000 != 0 {
		return -1, false
	}

	if (y-1000*n-4000*n_5000)/9000 < 0 {
		return -1, false
	}

	return (y - 1000*n - 4000*n_5000) / 9000, true
}

func cal_n_1000(n int, y int, n_5000 int) (int, bool) {
	if (-y+10000*n-5000*n_5000)%9000 != 0 {
		return -1, false
	}

	if (-y+10000*n-5000*n_5000)/9000 < 0 {
		return -1, false
	}

	return (-y + 10000*n - 5000*n_5000) / 9000, true
}
