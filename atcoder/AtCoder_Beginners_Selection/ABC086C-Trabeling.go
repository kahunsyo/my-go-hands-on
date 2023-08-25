package main

import (
	"fmt"
)

func main() {
	var n int
	var input_t, input_x, input_y int
	var now_location, to_location Location
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&input_t, &input_x, &input_y)
		to_location.time = input_t
		to_location.x = input_x
		to_location.y = input_y

		if is_able_to_move(now_location, to_location) == true {
			now_location = to_location
		} else {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}

type Location struct {
	time int
	x    int
	y    int
}

func is_able_to_move(now_location Location, to_location Location) bool {
	var delta_time int = abs(to_location.time - now_location.time)
	var delta_x int = abs(to_location.x - now_location.x)
	var delta_y int = abs(to_location.y - now_location.y)
	left_time := delta_time - delta_x - delta_y

	if left_time < 0 {
		return false
	} else {
		//If time is left, you can return to the same place by going back and forth.
		//Goint back and forth consist of 2 steps of moving.
		if left_time%2 == 0 {
			return true
		}
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
