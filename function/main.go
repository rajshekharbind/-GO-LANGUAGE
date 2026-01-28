package main
import "fmt"


func simpleFunction() {
	fmt.Println("This is a simple function")
}
func add(a,b int ) (result int ) {
   result = a+ b;
   return
}

func main(){
fmt.Println("i want to learn to the basic of the function:")
simpleFunction()
answer := add(10,50)
fmt.Println(" ans is to the :",answer)
}