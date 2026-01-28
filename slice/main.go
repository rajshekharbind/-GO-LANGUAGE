package main
import "fmt"
func main(){
	fmt.Println("i want to learn to the basic of the slice:")
	var slice1 =[]int{2,7,6,5,4}
	slice1 = append(slice1,10,56,78,90)
	fmt.Println("appending arary : ",slice1)
	fmt.Println("length of slice1 is : ",len(slice1))
	fmt.Println("capacity of slice1 is : ",cap(slice1))
	var slice2 = make([]int,4,8)
	slice2[0] = 12
	slice2[1] = 15
	slice2[2] = 18
	slice2[3] = 20
	fmt.Println("slice2 is : ",slice2)
	fmt.Println("length of slice2 is : ",len(slice2))
	fmt.Println("capacity of slice2 is : ",cap(slice2))


}