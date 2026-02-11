package main
import "fmt"


func main() {
	var arr [2]int
	fmt.Println(arr)

	arr[0] = 5
	arr[1] = 10
	fmt.Println(arr)

	arr2 := [3]string{"Go", "is", "awesome"}
	fmt.Println(arr2[0])
	fmt.Println(arr2)
}