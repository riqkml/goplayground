package main

import "fmt"

func main() {
	var fruits = []string{"apel", "mangga", "pisang", "jambu"}

	var amfruit = fruits[0:2]
	var mpfruit = fruits[1:3]

	fmt.Println("buah = \t", amfruit)
	fmt.Println("buah = \t", mpfruit)

	mpfruit[0] = "melon"

	fmt.Println("buah = \t", amfruit)
	fmt.Println("buah = \t", mpfruit)

	fruits = append(fruits, "pepaya")

	fmt.Println("buah = \t", fruits)

}