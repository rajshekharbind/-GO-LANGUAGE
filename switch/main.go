package main
import "fmt"

func main(){
	fmt.Println("i want to learn to the basic of the switch statement:")
	var day int = 4
	switch day {
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("tuesday")
	case 3:
		fmt.Println("wednesday")
	
	case 4:
		fmt.Println("thursday")
	case 5:
		fmt.Println("friday")
	case 6:
		fmt.Println("saturday")
	case 7:
		fmt.Println("sunday")
	default:
		fmt.Println("invalid day")
	}
var month int = 3
switch month {
case 1: 
	fmt.Println("january")
case 2:
	fmt.Println("february")
case 3:
	fmt.Println("march")
default:
	fmt.Println("invalid month")	
}
}