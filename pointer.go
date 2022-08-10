package main

import "fmt"

func main() {
	var numberA int = 4

	change(&numberA,10)
	
	fmt.Println(numberA)
}

func change(original *int, value int){
	*original = value
}