package main

import "fmt"

func Max(number ...int) int {
	max := number[0]
	for text := 0; text < len(number); text++ {
		if number[text] > max {
			max = number[text]
		}
	}
	return max
}

func main() {

	max := Max(4, 8, 9, 3, 21, 5, 10, 56, 84, 71, 32)
	fmt.Println(max)
}
