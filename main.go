package main

import (
	"fmt"
	"golang_practice/math_func"
)

func main() {
	// packages
	fmt.Println("Packages Hands- on")
	fmt.Println("-------------------------")
	fmt.Println("hello world")
	fmt.Println("add two number")
	fmt.Println(math_func.Add(5, 4))
	fmt.Println("subtract two numbers")
	fmt.Println(math_func.Subtracts(5, 4))

	//variable declaration
	datatype()
	loop()
	understand_defer()

	//go routine
	fmt.Println("\n\ngoroutine Hands- on")
	fmt.Println("-------------------------")
	go_thread()

	//channel
	fmt.Println("\n\n channel Hands- on")
	fmt.Println("-------------------------")
	channel()

	error_msg()
}
