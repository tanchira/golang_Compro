package main

import "fmt"

func hi(txt string) {
	for text := 1; text < 5; text++ {
		fmt.Println(text, ":", txt)
	}
}

func main() {
	go hi("hello")
	var input string
	fmt.Scanln(&input)

}
