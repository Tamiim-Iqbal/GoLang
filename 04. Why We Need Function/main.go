package main

import "fmt"

func printWelcomeMessage() {
	fmt.Println("Welcome to the application!")
}
func getUserName() string {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	return name
}
func getTwoNumbers() (int, int) {
	var num1, num2 int
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)
	return num1, num2
}
func add(a int, b int) int {
	return a + b
}
func display (name string, sum int) {
	fmt.Println("Hello,", name) 	
    fmt.Println("The sum is =", sum)
}
func goodbyeMessage() {
	fmt.Println("Thank you for using the application. Goodbye!")
}

func main() {
	printWelcomeMessage()
	name := getUserName()
	num1, num2 := getTwoNumbers()
	sum := add(num1, num2)
	display(name, sum)
	goodbyeMessage()
}

