package main

import "fmt"

var z = 87

func main() {
	x := 37

	fmt.Println("Die Zahl,", z, ",ist ja mal ultra nice, Gruß, ", x)
	fmt.Printf("%T\n", z)
	fmt.Printf("%T\n", x)
	fmt.Println(x)

}
