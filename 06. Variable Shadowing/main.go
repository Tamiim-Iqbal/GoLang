package main
import "fmt"

var a = 10 
func main(){
	age := 30
	
	if age > 18 {
		a := 47
		fmt.Println("Inside if block:", a) // prints 47
	}
	fmt.Println("Outside if block:", a) // prints 10
}