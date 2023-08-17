package main

import (
	"fmt"
)

func main() {
	var n int
	var myset MySet
	fmt.Scan(&n)

	var tmp int

	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		myset.Add(tmp)
	}

	fmt.Println(myset.len())
}

type MySet struct {
	Set []int
}

func (set *MySet) IsContain(x int) bool {
	//TODO: Repalce to more effective searching
	for _, v := range set.Set {
		if x == v {
			return true
		}
	}

	return false
}

func (set *MySet) Add(x int) {
	if set.IsContain(x) == true {
		return
	} else {
		set.Set = append(set.Set, x)
	}
}

func (set *MySet) len() int {
	return len(set.Set)
}
