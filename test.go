package main

import "fmt"

func main() {
	fmt.Println("Eyooo, was geht was geht")

	fmt.Println("bla bla alter, kein plan")

	for i := 5; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	one()

	beard()
}

func beard() {
	fmt.Println("Whadup, alles was hier drinne steht macht absolut keinen Sinn ^^")
}

func one() {
	fmt.Println("1")
}
