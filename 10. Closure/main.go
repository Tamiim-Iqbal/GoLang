package main
import "fmt"

const a = 10
var p = 100

func outer() func(){          // return a function
	money := 100
	age := 30

	fmt.Println("Age = ", age)

	show := func(){
		money = money + a + p
		fmt.Println("Money = ", money)
	}
	return show
}

func call() {
	incr1 := outer()  // show function returned by outer
	incr1()			  // call show 
	incr1()           // call show again

	incr2 := outer()  
	incr2()           
	incr2()           
}

func main() {
	call()
}

func init(){
	fmt.Println("=== Bank ===")
}



/* 

***** Code Segment : Compile Time : Build (Binary) ***** 
a = 10
outer = func() {...}
outerAnonymous = func() {...}          // just for illustration
call = func() {...}
main = func() {...}

***** Data Segment *****
p = 100

***** Stack Segment *****
init() 


**************************
Output:


**************************
*/