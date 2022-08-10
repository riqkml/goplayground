package main

import (
	"errors"
	"fmt"
	"strings"
)

func catch(){
	if r := recover(); r != nil {
		fmt.Println("Something went wrong ",r)
	}else{
		fmt.Println("application running well")
	}
}
func validate(input string) (bool, error) {
    if strings.TrimSpace(input) == "" {
        return false, errors.New("cannot be empty")
    }
    return true, nil
}

func main() {
	defer catch()
    var name string
    fmt.Print("Type your name: ")
    fmt.Scanln(&name)

    if valid, err := validate(name); valid {
        fmt.Println("halo", name)
    } else {
		panic(err.Error())
    }
}