package main

import "fmt"

func main() {

	nestedloop:
	for i := 0; i < 4; i++ {
		for c := 0; c < 4; c++ {
			if(i == 1){
				break nestedloop
			}
			fmt.Println("Matrix [",i,"][",c,"]")
		}
	}
}