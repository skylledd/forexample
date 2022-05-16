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

}

func test() {
	fmt.Println("jo jo, test", y)
}
