package main

import "fmt"

func main() {
	number := 0
	for countdown := 1; countdown < 10; countdown-- {
		number -= countdown
	}
	fmt.Println(number)
}