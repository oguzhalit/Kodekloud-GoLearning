package main

import "fmt"

func main() {
	var a bool = true
	result := 10 > 50
	fmt.Println(!(a || result))
	/*
		false
	*/
}
