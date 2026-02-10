package main

import "fmt"

type User struct {                   // type User (Custom Type)
	Name string
	Age  int
}

func printUserDetails (usr User) {
	fmt.Println("Name: ", usr.Name)
	fmt.Println("Age: ", usr.Age)
}

// receiver function 
func (usr User) printDetails () {
	fmt.Println("Name: ", usr.Name)
	fmt.Println("Age: ", usr.Age)
}

func main () {

	var user1 User 
	user1 = User{ 
		Name : "Tamim",             // property or member variable
		Age  : 25, 
	}

	printUserDetails(user1)
	user1.printDetails()            // calling receiver function

	user2 := User{                  // another way to declare and initialize a struct variable (num := 10)
		Name : "Roki",
		Age  : 30,
	}   
	fmt.Println("*** User 2 ***")
	fmt.Println("All Data: ",user2)
	fmt.Println("Name: ", user2.Name)
	fmt.Println("Age: ", user2.Age)          

	
	/*
	*** Compilation ***
	== Code Segment ==
	1. User = type User struct {...}
	2. printUserDetails = func (usr User) {...}
	3. printDetails = func (usr User) {...}
	3. main = func () {...}
	*/
}