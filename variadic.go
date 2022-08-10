package main

import (
	"fmt"
)

func main() {
	var result = calcAverage(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("AVG = %.2f", result)
	fmt.Println(msg)
}

func calcAverage(scores ...int) float64 {
	total := 0

	for _, score := range scores {
		total += score
	}
	var avg = float64(total) / float64(len(scores))

	return avg
}