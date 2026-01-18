package main

import "fmt"

type User struct {                   // type User (Custom Type)
	Name string
	Age  int
}

func main () {
	// var a int                    // int type 
	// a = 10
	// fmt.Println(a)

	var user1 User 
	user1 = User{ 
		Name : "Tamim",           // property or member variable
		Age  : 25, 
	}

	user2 := User{                  // another way to declare and initialize a struct variable (num := 10)
		Name : "Roki",
		Age  : 30,
	}   

	fmt.Println("*** User 1 ***")
	fmt.Println("All Data: ",user1)
	fmt.Println("Name: ", user1.Name)
	fmt.Println("Age: ", user1.Age)

	fmt.Println("*** User 2 ***")
	fmt.Println("All Data: ",user2)
	fmt.Println("Name: ", user2.Name)
	fmt.Println("Age: ", user2.Age)

	/*
	*** Compilation ***
	== Code Segment ==
	1. User = type User struct {...}
	2. main = func () {...}


	*/
}