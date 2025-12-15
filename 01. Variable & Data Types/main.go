package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	a := 5
	fmt.Println("Value of a:", a)

	/* Data Types in Go:
		int
		float32
		bool
		string
	*/
	
	var x int = 10
	fmt.Println("Value of x:", x)
	
	var y = 20.5
	fmt.Println("Value of y:", y)

	var name string = "GoLang"
	fmt.Println("Value of name:", name)

	b := true
	fmt.Println("Value of b:", b)

	b = false 
	fmt.Println("Updated value of b:", b)

	// b = "golang"     -> cannot use "golang" (untyped string constant) as bool value in assignment

	const pi = 3.14
	fmt.Println("Value of pi:", pi)

	// pi = 3.14159    -> cannot assign to pi (constant value)

}