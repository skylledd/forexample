package main

import (
	"fmt"
)

func main() {
	fmt.Println("Eyooo, was geht was geht")

	n, err := fmt.Println("bla bla alter, kein plan")
	fmt.Println(n)
	fmt.Println(err)

	for i := 10; i < 30; i++ {
		if i%5 == 0 {
			fmt.Println(i)
		}
	}

	mom()

	one()

	beard()

	random()
}

func beard() {
	fmt.Println("Whadup, alles was hier drinne steht macht absolut keinen Sinn ^^")
}

func one() {
	fmt.Println("1")
}

func random() {
	n, err := fmt.Println("Random")
	fmt.Println(n)
	fmt.Println(err)
}

func mom() {
	n, err := fmt.Println("Test Test, 3 2 1")
	fmt.Println(n)
	fmt.Println(err)
}
