package main

import "fmt"

func main() {
	var messages = make(chan string)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("Hello %s",who)
		messages <- data
	}

	go sayHelloTo("Riqki")
	go sayHelloTo("Kamal")
	go sayHelloTo("Amrela")

	var message1 = <- messages
	fmt.Println(message1)
	
	var message2 = <- messages
	fmt.Println(message2)
	
	var message3 = <- messages
	fmt.Println(message3)
}