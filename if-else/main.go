package main
import "fmt"
func main(){
	fmt.Println("i want to learn basic of if and else statement:")
	var age int = 18
	if age < 18 {
		fmt.Println("you are not eligible for voting")
	} else {
		fmt.Println("you are eligible for voting")
	}
	x:= 23
	if x%2 == 0 {
		fmt.Println("x is even number")
	} else {
		fmt.Println("x is odd number")
	}
	y:= 46
	if y%2 == 0 && x%2 != 0 {
		fmt.Println("both x is odd and y are even numbers")
	} else {
		fmt.Println("either x or y is not an even number")
	}
}