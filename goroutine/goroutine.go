package main

import (
	"fmt"
	"time"
)

func funcA(chA chan<- string) {
	for i := 0; i < 10; i++ {
		fmt.Print("A")
		time.Sleep(10 * time.Millisecond)
	}
	chA <- "Finished"
}

func main() {
	chA := make(chan string)
	defer close(chA)
	go funcA(chA)

	for i := 0; i < 10; i++ {
		fmt.Print("M")
		time.Sleep(20 * time.Millisecond)
	}

	msg := <-chA
	fmt.Println(msg)
}
