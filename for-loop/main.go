package main
import "fmt"

func main(){
	fmt.Println("i want to learn to the basic of for loop:")
	for i:=1; i<=10; i++{
		fmt.Println("print to the i : value ",i)

	}
	counter :=0
	for counter <5 {
		fmt.Println("counter value is :",counter)
		counter++
	if(counter == 3){
		break
	} else {
		continue
	}
	
}
numbers := []int{2,4,6,8,10,12,14,16,18,20}
for index, value := range numbers {
	fmt.Printf("index: %d, value: %d\n", index, value)
}
data:= string("go is a programming language")
for i, word := range data {
	fmt.Printf("index: %d, word: %c\n", i, word)
 } 
}