package main

import "fmt"

func main() {

	var cars = [...]string{"volvo", "KIA", "ford", "mercedez"}

	fmt.Println("Total cars = \t",len(cars))

	for i := 0; i < len(cars); i++ {
		fmt.Println("cars no.",i," = \t",cars[i])
	}
	
	for a, v := range cars {
		fmt.Println("cars idx.",a," = \t",v)
	}
}