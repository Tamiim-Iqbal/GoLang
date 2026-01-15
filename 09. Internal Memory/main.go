package main
import "fmt"

var a = 10
const p = 11

func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func call(){
	sum := func(x int, y int) {
		z := x + y
		fmt.Println(z)
	}
	sum(5, 6)
	sum(p, a)
}

func main() {
	add (5, 4)
	add(a, 3)

	call()
	fmt.Println(p)
}

func init() {
	fmt.Println("Hello")
}