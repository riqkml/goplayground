package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Print("Input some number ")
	fmt.Scanln(&input)

	var number int 
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "is number")
	}else{
		fmt.Println("error")
		fmt.Println(err.Error())
	}
}

