package main

import "fmt"

func main() {
	dawn1 := []int{1, 2, 3}
	dawn2 := dawn1
	fmt.Println(dawn1, dawn2)
	dawn1[1] = 5
	fmt.Println(dawn1, dawn2)
	fmt.Println("-----")
	dawn3 := []int{1, 2, 3}

}
