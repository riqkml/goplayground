package main

import "fmt"

func main() {
	var number = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var howMany, findProcess = findMax(number, 4)

	fmt.Println("found \t:", howMany) // 9
	fmt.Println("value \t:", findProcess())
}

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int

	for _, number := range numbers {
		if number < max {
			res = append(res, number)
		}
	}

	return len(numbers), func() []int {
		return res
	}
}
