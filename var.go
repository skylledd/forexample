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

	fmt.Println("Test", x)

	paket()

	mom()
}

func test() {
	fmt.Println("jo jo, test", y)
}

func paket() {
	fmt.Println("Test 24", y)
}
