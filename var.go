package main

import (
	"fmt"
)

var y = 84

func main() {

	test()

	x := 42
	fmt.Println(x)

	fmt.Println(y)

	321()

	fmt.Println("Test", x)
}

func test() {
	fmt.Println("jo jo, test", y)
}

func 321() {

	fmt.Println(y)
}
