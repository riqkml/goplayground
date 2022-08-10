package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("start")
	<-time.After(4 * time.Second)
	fmt.Println("finish")
}