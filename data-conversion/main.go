package main
import (
	"fmt"
	"strconv"
)
func main(){
	fmt.Println("Data Conversion in Go:")
	// Converting integer to string
	var num int =90
	strconvedStr := strconv.Itoa(num)
	fmt.Println("Integer to String conversion:", strconvedStr)
    fmt.Printf("type of the data %T\n",num)

	// Converting string to integer
	strnum := "120"
	strconvedInt, err := strconv.Atoi(strnum)
	if err !=nil {
		fmt.Println("Error converting string to integer:", err)
	}
	fmt.Println("String to Integer conversion:", strconvedInt)
	fmt.Printf("type of the data %T\n",strnum)


	number_string := "654321"
	number_int, _ := strconv.Atoi(number_string)
	fmt.Println("Converted Integer:", number_int)
	fmt.Printf("Type of number_int: %T\n", number_int)

	float_string := "3.14159"
	float_num, _ := strconv.ParseFloat(float_string, 64)
	fmt.Println("Converted Float:", float_num)
	fmt.Printf("Type of float_num: %T\n", float_num)

	bool_string := "true"
	bool_val, _ := strconv.ParseBool(bool_string)
	fmt.Println("Converted Boolean:", bool_val)
	fmt.Printf("Type of bool_val: %T\n", bool_val)

	 
}