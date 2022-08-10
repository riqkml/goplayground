package main

import (
	"fmt"
	"os"
)

func main() {
	orderSomeFood("pizza")
	orderSomeFood("burger")
	os.Exit(500)
}

func orderSomeFood(menu string) {
	defer fmt.Println("terimakasih telah order!")
	if menu == "pizza" {
        fmt.Print("Pilihan tepat!", " ")
        fmt.Print("Pizza ditempat kami paling enak!", "\n")
		return 
	}

	
    fmt.Println("Pesanan anda:", menu)
}