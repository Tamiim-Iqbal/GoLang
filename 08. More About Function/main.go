package main
import "fmt"

// Higher order function that takes another function as a parameter
func processOperation(a int, b int, op func(p int, q int)) {
	op(a, b)
}

func add(a int, b int) {                // parameters => a, b
	c := a + b
	fmt.Println("Sum:", c)
}

func main() {
	processOperation(2, 5, add)         // arguments => 3, 5
} 