package main
import (
	"fmt"
	"example/mathlib"
)


var ( 
	a = 20
 	b = 30
)

func add(x int, y int){
	z := x + y
	fmt.Println("Sum is:", z)
}

func main () {
	var p = 30
	var q = 40
	
	add(p, q)

	add(a, q)

	// add(a, z) // This will cause a compile-time error because 'z' is not defined in this scope

	// Local Scope 

	x := 18

	if x >= 18 {
		// if block scope
		var g = 2
		fmt.Println("I'm matured boy")
		fmt.Println("I have ", g, "girlfriends")
	}

	// Package Scope
	var m = 4
	var n = 5
	//fmt.Println("Multiplication is:", multi(m, n))

	// showing custom package
	fmt.Println("Custom Package: ")

	mathlib.Sub(m , n)

}