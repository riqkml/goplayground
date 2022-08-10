package main

import "fmt"

func main() {
	var data = []string{"wick", "jason", "ethan"}
	var findJason = filter(data, func(each string) bool {
		return each == "jason"
	})
	var conditional string

	if len(findJason) > 0 {
		conditional = "FIND"
	}else{
		conditional = "NOT FOUND"
	}

	fmt.Println("Jason find ? ", conditional)
}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

