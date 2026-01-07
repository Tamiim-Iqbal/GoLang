package main

import "fmt"

var a = 10

// Standard function or Named function
func add(a int, b int) {
	fmt.Println(a + b)
}
func main() {
	fmt.Println(a)
	add(7, 4)

	// Anonymous function
	func() {
		fmt.Println("Hello from anonymous function")
	}()
	// Another way to call anonymous function
	func(a int, b int) {
		c := a + b
		fmt.Println("Sum from anonymous function:", c)
	}(7, 4)

	// Function expression or Function as/in variable
	sum := func(a int, b int) {
		c := a + b
		fmt.Println("Sum from function expression:", c)
	}
	sum(10, 15)
}


// init function
func init() {
	fmt.Println(a)
	a = 20
}
