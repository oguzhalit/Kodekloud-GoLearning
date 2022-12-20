package main

import "fmt"

func main() {
	var a, b bool = false, false
	fmt.Println(a && b)
	fmt.Println(a || b)
	/*
		false && false = false
		false || false = false
	*/
}
