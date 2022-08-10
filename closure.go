package main

import "fmt"

func main() {

	var getMinMax = func(n []int) (min int, max int) {

		for idx, element := range n {
			switch {
			case idx == 0:
				min, max = element, element
			case element > max:
				max = element
			case element < min:
				min = element
			}
		}
		return 
	}
	var sample = []int{2, 3, 4, 3, 4, 2, 3}
	var min,max  = getMinMax(sample)
	
    fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n",sample , min, max)
}