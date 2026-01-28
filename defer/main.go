package main
import "fmt"
func add(a int, b int) int {
	return a + b
}
func main() {
	fmt.Println("i want to learn consepts of defer in golang:")
	defer fmt.Println("This is the first deferred statement.")
	defer fmt.Println("This is the second deferred statement.")
	fmt.Println("This is the main function execution.")
	var result int = add(5, 10)
	fmt.Printf("Result of addition: %d\n", result)

	defer fmt.Println("This is the third deferred statement.")
	fmt.Println("Exiting main function.")

	
}