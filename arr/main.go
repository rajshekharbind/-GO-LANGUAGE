package main
import "fmt"



func main(){
	fmt.Println("i want to learn to the basic of the array:")
	var arr [6]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50
	arr[5] = 60
	fmt.Println("array is : ",arr)
	fmt.Println("length of array is : ",len(arr))
	var arr1 = [5]string{"apple","banana","grapes","mango","orange"}
	 // we cant use append function in array
	 
	fmt.Println("string of arr1 is :",arr1)
}
