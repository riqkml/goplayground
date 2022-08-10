package main

import "fmt"

type student struct {
	name  string
	grade int
}

func main() {
	var s1 student

	s1.name = "riqki"

	var s2 = student{"kamal", 10}

	var s3 = student{name: "amrela"}

	fmt.Println(s1.name,s2.name,s3.name)
}