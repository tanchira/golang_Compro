package main

import "fmt"

type I interface{}

func main() {
	var any I
	any = "Hello"
	yes, ok := any(string)
	fmt.Println(yes.ok)
}
