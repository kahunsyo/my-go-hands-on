package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const DREAM string = "dream"
	const DREAMER string = "dreamer"
	const ERASE string = "erase"
	const ERASER string = "eraser"

	var s string

	fmt.Scan(&s)

	check_order := [...]string{DREAMER, ERASER, DREAM, ERASE}

	for len_s := utf8.RuneCountInString(s); len_s > 0; len_s = utf8.RuneCountInString(s) {
		//fmt.Println(s)
		for _, check_word := range check_order {
			left_s, is_deleted := del_word_at_tail(check_word, s)
			if is_deleted == true {
				s = left_s
				goto next_del
			}
		}

		fmt.Println("NO")
		return

	next_del:
	}
	fmt.Println("YES")

}

func split_at_n_th_char(n int, s string) (string, string) {
	return s[:n], s[n:]
}

func del_word_at_tail(word string, s string) (string, bool) {
	len_word := utf8.RuneCountInString(word)
	len_s := utf8.RuneCountInString(s)

	if len_word > len_s {
		return s, false
	}

	head, tail := split_at_n_th_char(len_s-len_word, s)
	if tail == word {
		return head, true
	} else {
		return s, false
	}
}
