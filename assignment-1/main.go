package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numType := ""

	for _, n := range nums {
		if n%2 == 0 {
			numType = "even"
		} else {
			numType = "odd"
		}

		fmt.Printf("%v is %v \n", n, numType)
	}
}
