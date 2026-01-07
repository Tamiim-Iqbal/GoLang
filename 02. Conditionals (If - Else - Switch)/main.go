package main
 import "fmt"

func main() {
	age := 20
	sex := "male"
	education := "graduated"

	if age >= 18 {
		fmt.Println("You are eligible to be married")
	} else if age < 18 {
		fmt.Println("You are not eligible to be married")
	} else if age == 18 {
		fmt.Print("You are just a kid")
	}	

	if age >= 18 && sex == "male" {
		fmt.Println("You are eligible to be married")
	}

	switch education {
	case "graduated":
		fmt.Println("You are eligible for higher studies")
	case "not graduated":
		fmt.Println("You need to graduate first")
	default:
		fmt.Println("Please provide a valid education status")
	}
}