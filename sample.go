package main

import "fmt"

func main() {
	fmt.Println("hello, world")

	num := 123
	str := "ABC"

	fmt.Println(num, str)

	var_()
	const_()
	comment_()
	semicoron_()
	custom_type_()
	convert_type_()
	array_()
	slice_()
}

func var_() {
	var x int = 123
	var y int
	z := 22
	fmt.Println(x, y, z)
}

func const_() {
	const foo int = 123
	const bar = 111
	fmt.Println(foo, bar)
}

func comment_() {
	// this is a comment
	/* this
	is
	a
	comment
	*/
}

func semicoron_() {
	x := 1
	x++
	fmt.Println(x)
}

func custom_type_() {
	type UtcTime string
	type JstTime string
	var t1 UtcTime = "00:00"
	var t2 JstTime = "09:00"

	//t1 = t2 //error
	t1 = UtcTime(t2)
	fmt.Println(t1)
}

func convert_type_(){
	var a1 uint16 = 1234
	var a2 uint32 = uint32(a1)
	fmt.Println(a2)
}

func array_(){
	a := [3]string{}
	a[0] = "Red"
	a[1] = "Green"
	a[2] = "Blue"

	fmt.Println(a)
	fmt.Println(a[0])

	b := [3]string{"R","G","B"}
	fmt.Println(b)

	c := [...]int{1,2,3,4}
	fmt.Println(c)
}

func slice_(){
	a := []string{}
	a = append(a, "Red")
	a = append(a, "Green")
	a = append(a, "Blue")
	fmt.Println(a)

	fmt.Println(len(a), cap(a))

	b := make([]int, 3 ,1024)
	fmt.Println(b)
	fmt.Println(len(b), cap(b))
}

func map_(){
	a := map[string]int{
	"x": 100,
	"y": 200,
}

}