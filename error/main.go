package main
import "fmt"

func divide(a ,b float64) (result float64, err error) {
   if b == 0 {
	   err = fmt.Errorf("division by zero is not allowed")
	   return
   }
	result = a/ b 
   return
}

func main(){
fmt.Println("i want to learn to the basic of the function:")


answer, err := divide(10,3)
if err != nil{
	fmt.Println("Error:", err)
}
fmt.Println(" ans is to the :",answer)
}