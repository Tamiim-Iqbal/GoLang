package main

import "fmt"

func add(num1 int, num2 int) {
	sum := num1 + num2

	fmt.Println("Sum of c & d is:", sum)
}
func getNumbers (num1 int , num2 int) (int, int) {
	sum := num1 + num2

	mul := num1 * num2

	return sum , mul
} 

func main() {
	a := 10
	b := 20
	c := 30
	d := 40

	sum := a + b

	fmt.Println("Sum of a & b is:", sum)

	add(c, d)

	// Using multiple return values

	p, q := getNumbers(a, b)

	fmt.Println("Sum:", p)
	fmt.Println("Multiplication:", q)
}
