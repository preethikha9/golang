package main

import (
	"fmt"
	"golang_practice/math_func"
)

func datatype() {
	//various forms of variable declaration
	fmt.Println("\n\nExecuting datatype block")
	fmt.Println("-------------------------")
	var num1 int
	var num2 int = 2
	num3 := 4
	num1 = 1

	//constant declation
	const num4 int = 1
	fmt.Println("add two number")
	fmt.Println(math_func.Add(num1, num2))
	fmt.Println("subtract two numbers")
	fmt.Println(math_func.Subtracts(num4, num3))

	//multiple varaible declation
	var num11, num22 = 2, 3
	num111, num222 := 2, 3
	fmt.Println(num11, num22)
	fmt.Println(num111, num222)

	var name string = "Preethikha"
	fmt.Println("Hi", name)

}

func loop() {
	fmt.Println("\n\nloop")
	fmt.Println("-------------------------")
	for i := 1; i < 5; i++ {
		fmt.Println(i)
		if i == 3 {
			break
		}
	}
	fmt.Println("\nRange")
	words := []string{"word1", "word2"}
	for _, val := range words {
		fmt.Println(val)
	}

}

func understand_defer() {
	defer fmt.Println("defer executes when the function comms to end")
	fmt.Println("Main")
	value := test()
	fmt.Println(value)

}

func test() (size int) {
	defer func() { size = 20 }()
	size = 30
	return
}
