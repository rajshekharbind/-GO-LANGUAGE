// package main
// import (
// 	"fmt"
// 	"sync"
// )

// func printNumbers(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 1; i <= 5; i++ {
// 		fmt.Printf("Number: %d\n", i)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go printNumbers(&wg)
// 	wg.Wait()
// 	fmt.Println("All goroutines finished executing")
// 	// Note: In real applications, use sync.WaitGroup or other synchronization methods to manage goroutines properly.

	
// }

package main

// Import required packages
import (
	"fmt"   // Used for formatted output
	"sync"  // Provides synchronization primitives like WaitGroup
	"time"  // Used to simulate time-consuming tasks
)


// -----------------------------
// Function: printNumbers
// Purpose: Prints numbers from 1 to 5 using a goroutine
// -----------------------------
func printNumbers(wg *sync.WaitGroup) {

	// Ensure Done() is called when function exits
	// This decreases the WaitGroup counter by 1
	defer wg.Done()

	fmt.Println("printNumbers goroutine started")

	for i := 1; i <= 5; i++ {
		fmt.Printf("Number: %d\n", i)

		// Simulate some processing delay
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("printNumbers goroutine finished")
}


// -----------------------------
// Function: printLetters
// Purpose: Prints letters from A to E using a goroutine
// -----------------------------
func printLetters(wg *sync.WaitGroup) {

	// Notify WaitGroup when goroutine completes
	defer wg.Done()

	fmt.Println("printLetters goroutine started")

	for ch := 'A'; ch <= 'E'; ch++ {
		fmt.Printf("Letter: %c\n", ch)
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("printLetters goroutine finished")
}


// -----------------------------
// main function
// Entry point of the program
// -----------------------------
func main() {

	fmt.Println("Main function started")

	// Declare a WaitGroup
	// WaitGroup is used to wait for multiple goroutines to finish
	var wg sync.WaitGroup

	// Add the number of goroutines to wait for
	// IMPORTANT: Add() must be called BEFORE starting goroutines
	wg.Add(2)

	// Start goroutines
	go printNumbers(&wg)
	go printLetters(&wg)

	// Wait until all goroutines call Done()
	wg.Wait()

	// This line executes only after all goroutines finish
	fmt.Println("All goroutines finished executing")

	fmt.Println("Main function finished")
}
