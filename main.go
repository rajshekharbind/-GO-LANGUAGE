package main

import (
	"fmt"
	"mylearning/myutil"
)
func main(){
	fmt.Println("Hello , go language")
	myutil.PrintMessage("hello go language")
	var a int = 10;
	fmt.Println(a);
	var b string = "Goland"
	fmt.Println(b)
	var c bool = true
	fmt.Println(c)
	const pi = 3.14
	fmt.Println(pi)
	var d float32 = 6.89
	fmt.Println(d)

	age := 20
	height := 5.8
	name := "raj"
	fmt.Println("age :", age, "height :", height, "name :", name )
	fmt.Printf("age is raj %d\n", age)
	fmt.Printf("height is raj %.2f\n", height)
	fmt.Printf("name is raj %s\n", name)
	
}