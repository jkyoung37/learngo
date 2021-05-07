package main

import "fmt"

func superApp(numbers ...int) int {

	total := 0
	for _, number := range numbers {
		total = total + number
	}

	return total
}

func main() {
	result := superApp(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
