package main

import (
	"errors"
	"fmt"
	"os"
)

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
	if_()
	switch_()
	for_()
	goto_()
	fmt.Println(func_(1, 2))
	fmt.Println(func_2(1, 2))
	func_3(1, 2, 3, 4)

	//---Struct---
	var p1 Person
	p1.SetPerson("yamada", 22)
	name, age := p1.GetPerson()
	fmt.Println(name, age)

	p2 := Person{"tarou", 99}
	fmt.Println(p2)

	//---Interface---
	p3 := Person{"hanako", 99}
	b1 := Book{"book titile", 199}

	PrintOut(p3)
	PrintOut(b1)

	interface_type_(1)
	interface_type_("a")

	dict_ := dict{
		"name": "yamada",
		"age":  22,
		"address": dict{
			"zip": "0000",
			"tel": "000-0000-0000",
		},
	}
	fmt.Println(dict_)

	pointer_()
	new_()
	defer_()
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

func convert_type_() {
	var a1 uint16 = 1234
	var a2 uint32 = uint32(a1)
	fmt.Println(a2)
}

func array_() {
	a := [3]string{}
	a[0] = "Red"
	a[1] = "Green"
	a[2] = "Blue"

	fmt.Println(a)
	fmt.Println(a[0])

	b := [3]string{"R", "G", "B"}
	fmt.Println(b)

	c := [...]int{1, 2, 3, 4}
	fmt.Println(c)
}

func slice_() {
	a := []string{}
	a = append(a, "Red")
	a = append(a, "Green")
	a = append(a, "Blue")
	fmt.Println(a)

	fmt.Println(len(a), cap(a))

	b := make([]int, 3, 1024)
	fmt.Println(b)
	fmt.Println(len(b), cap(b))
}

func map_() {
	a := map[string]int{
		"x": 100,
		"y": 200,
	}
	fmt.Println(a)
}

func if_() {
	var x int = 3
	var y int = 2

	if x > y {
		fmt.Println("Ture")
	} else if x == y {
		fmt.Println("equal")
	} else {
		fmt.Println("False")
	}

}

func switch_() {
	var mode string = "Running"

	switch mode {
	case "running":
		fmt.Println("Running")
	case "stop":
		fmt.Println("Stop")
	default:
		fmt.Println("unknown")
	}

	var x int = 1
	var y int = 2

	switch {
	case x > y:
		fmt.Println("Big")
	case x < y:
		fmt.Println("Small")
	default:
		fmt.Println("Equal")
	}

	switch mode {
	case "Running":
		fmt.Println("Running")
		fallthrough
	case "Stop":
		fmt.Println("Stop")
		fallthrough
	default:
		fmt.Println("Complete")
	}
}

func for_() {
	var x int = 0
	var y int = 10
	for x < y {
		x++
		fmt.Println(x)
	}

	colors := [...]string{"red", "green", "blue"}
	for i, color := range colors {
		fmt.Printf("%d: %s\n", i, color)
	}
}

func goto_() {

	GetFileName := func() (string, error) {
		return "sample.txt", nil
	}

	ReadFile := func(filename string) (string, error) {
		return "Hello World", errors.New("Can't read file")
	}

	funcA := func() (string, error) {
		var err error
		filename := ""
		data := ""

		filename, err = GetFileName()
		if err != nil {
			fmt.Println(err)
			goto Done
		}

		data, err = ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			goto Done
		}

		fmt.Println(data)

	Done:
		return data, err
	}

	funcA()

}

func func_(x int, y int) int {
	return x + y
}

func func_2(x int, y int) (int, int) {
	return x, y
}

func func_3(x ...int) {
	for i, num := range x {
		fmt.Println(i, num)
	}
}

// ---Struct---
type Person struct {
	name string
	age  int
}

func (p *Person) SetPerson(name string, age int) {
	p.name = name
	p.age = age
}

func (p *Person) GetPerson() (string, int) {
	return p.name, p.age
}

func (p Person) ToString() string {
	return p.name
}

type Book struct {
	title string
	yen   int
}

func (b Book) ToString() string {
	return b.title
}

type Printable interface {
	ToString() string
}

func PrintOut(p Printable) {
	fmt.Println(p.ToString())
}

func interface_type_(a interface{}) {
	fmt.Println(a)
}

type dict map[string]interface{}

func pointer_() {
	var a int
	var p *int

	p = &a
	*p = 10

	fmt.Println(a)
	fmt.Println(p, *p)

	zero := func(x int, p *int) {
		x = 0
		*p = 0
	}

	x := 1
	y := 1
	fmt.Println(x, y)
	zero(x, &y)
	fmt.Println(x, y)
}

func new_() {
	bookList := []*Book{}

	for i := 0; i < 10; i++ {
		book := new(Book)
		book.title = "none"
		bookList = append(bookList, book)
	}

	for _, book := range bookList {
		fmt.Println(book.title)
	}
}

func defer_() {
	fp, err := os.Open("Sample.txt")
	if err != nil {
		return
	}

	defer fp.Close()

	//Do something...

}
