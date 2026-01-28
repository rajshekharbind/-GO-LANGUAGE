package main
import "fmt"


func modifyValueByReference(val *int) {
	*val = *val + 20
	fmt.Println("Value inside modifyValueByReference:", *val)
}


func main(){
	fmt.Println("i want to learn to the basic of the pointer concept:")
	var a int = 10
	var p *int = &a  // p is a pointer to an integer, storing the address of a
	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", &a)
	fmt.Println("Value of p (address of a):", p)
	fmt.Println("Value pointed to by p:", *p) // dereferencing pointer p to get the value of a

	var pointer *int  // pointer initialized to nil
	if pointer == nil {
		fmt.Println("pointer is nil")
	}
	// modifying value using pointer
	b:= 60
	modifyValueByReference(&b)
	fmt.Println("Value of b after modifyValueByReference:", b)

}