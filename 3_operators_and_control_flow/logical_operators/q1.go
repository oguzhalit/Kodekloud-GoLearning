package main

import "fmt"

func main() {
	var a, b bool = true, false
	fmt.Println(a && b)
	fmt.Println(a || b)
	/*
	  true && false = false
	  true || false = true
	*/
}
