package main

import (
	"fmt"
)

func main() {
	var numbers []int

	for i := 0; i <= 10; i++ {
		numbers = append(numbers, i)
	}

	for n := range numbers {
		if n%2 == 0 {
			fmt.Printf("%v is even.\n", n)
		} else {
			fmt.Printf("%d is odd.\n", n)
		}
	}
}
